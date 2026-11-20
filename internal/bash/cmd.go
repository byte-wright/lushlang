package bash

import (
	"strconv"
	"strings"

	"github.com/byte-wright/lush/internal/bcode"
)

// activeCmd is a cmd whose body is being translated. Leaving the body early has to stop
// its procs.
type activeCmd struct {
	cmd   *bcode.Cmd
	files string
	// loops is the number of loops around the cmd, a break at that depth leaves it
	loops int
}

// The procs are jobs with exec, so a signal reaches the command and not a subshell. exec
// also skips the script's functions and builtins. A FIFO only opens when both sides open
// it, so the opens of a pipe rendezvous in the background jobs and the reader.
func (b *bash) cmd(st *bcode.Cmd) {
	main := st.Procs[len(st.Procs)-1]
	files := ""

	if needsFiles(st) {
		files = main.Code.Name + "_f"
		b.useFunc("cmdInit")
		b.print("lsh__cmdInit")
		b.print("L_cmd_seq=$(( L_cmd_seq + 1 ))")
		b.print("local " + files + "=\"$L_cmd_dir/$L_cmd_seq\"")
	}

	fifos := []string{}
	for i, p := range st.Procs {
		if p.Stdin != nil && p.Stdin.Kind == bcode.InputString {
			b.print("printf '%s' " + b.atom(p.Stdin.Value) + " > " + file(files, i, "i"))
		}

		if pipeOf(p) != "" {
			fifos = append(fifos, file(files, i, "p"))
		}
	}

	if len(fifos) > 0 {
		b.print("mkfifo " + strings.Join(fifos, " ") + " || exit 2")
	}

	index := map[*bcode.Proc]int{}
	for i, p := range st.Procs {
		index[p] = i
		// the command of the statement runs in the foreground when nothing reads it, so
		// it gets Ctrl-C like any command and a program like ssh can handle it
		b.start(p, i, index, files, i < len(st.Procs)-1 || st.Pipe != nil)
	}

	if st.Pipe != nil {
		b.runBody(st, main, files)
	}

	b.waitSources(st)

	for i, p := range st.Procs {
		b.collect(p.Stdout, file(files, i, "o"))
		b.collect(p.Stderr, file(files, i, "e"))
	}

	if files != "" {
		b.print(`rm -f "$` + files + `".*`)
	}
}

func needsFiles(st *bcode.Cmd) bool {
	for _, p := range st.Procs {
		if p.Stdin != nil && p.Stdin.Kind == bcode.InputString {
			return true
		}

		for _, out := range []*bcode.ProcOutput{p.Stdout, p.Stderr} {
			if out != nil && (out.Kind == bcode.OutCollect || out.Kind == bcode.OutPipe) {
				return true
			}
		}
	}

	return false
}

// file is the path of a temp file of the statement, of proc i
func file(files string, i int, kind string) string {
	return `"$` + files + "." + strconv.Itoa(i) + kind + `"`
}

func pid(p *bcode.Proc) string {
	return p.Code.Name + "_p"
}

// pipeOf returns the fd of the proc's pipe, "" without one
func pipeOf(p *bcode.Proc) string {
	switch {
	case p.Stdout != nil && p.Stdout.Kind == bcode.OutPipe:
		return ">"
	case p.Stderr != nil && p.Stderr.Kind == bcode.OutPipe:
		return "2>"
	}

	return ""
}

func (b *bash) start(p *bcode.Proc, i int, index map[*bcode.Proc]int, files string, background bool) {
	job := "( "

	// bash 3.2 keeps ignoring SIGINT in a background job, there the EXIT trap stops it
	if background {
		job += "trap - INT QUIT; "
	}

	if p.Dir != nil {
		job += "CDPATH= cd -- " + b.atom(p.Dir) + " || exit; "
	}

	// one export expands all values before it sets one, an env var named like a
	// variable of the script can't change the next value
	if len(p.Env) > 0 {
		env := []string{}
		for _, e := range p.Env {
			env = append(env, e.Name+"="+b.atom(e.Value))
		}
		job += "export " + strings.Join(env, " ") + "; "
	}

	words := []string{"exec", "--", b.atom(p.Name)}
	for _, arg := range p.Args {
		switch {
		case arg.Source != nil:
			words = append(words, file(files, index[arg.Source], "p"))
		case arg.Spread:
			words = append(words, `"${`+arg.Value.(*bcode.VarValue).Name+`[@]}"`)
		default:
			words = append(words, b.atom(arg.Value))
		}
	}

	job += strings.Join(words, " ") + " )" + b.redirections(p, i, index, files, background)

	if !background {
		b.print(job)
		b.print("local " + p.Code.Name + "=$?")
		return
	}

	// a job that gets a signal right after its fork runs the EXIT trap of the script,
	// this tells it that it's not the script
	b.print("L_cmd_fork=1")
	b.print(job + " &")
	b.print("local " + pid(p) + "=$!")
	b.print("L_cmd_fork=")
}

// redirections opens the pipe first: when another redirection fails before, its reader
// would wait forever. stderr comes before stdout, so the error of a file that can't be
// opened goes where stderr goes.
func (b *bash) redirections(p *bcode.Proc, i int, index map[*bcode.Proc]int, files string, background bool) string {
	redirs := []string{}

	if op := pipeOf(p); op != "" {
		redirs = append(redirs, op+" "+file(files, i, "p"))
	}

	if p.Stderr != nil && p.Stderr.Kind != bcode.OutPipe && p.Stderr.Kind != bcode.OutStdout {
		redirs = append(redirs, b.output("2>", p.Stderr, file(files, i, "e")))
	}

	if p.Stdout != nil && p.Stdout.Kind != bcode.OutPipe {
		redirs = append(redirs, b.output(">", p.Stdout, file(files, i, "o")))
	}

	if p.Stderr != nil && p.Stderr.Kind == bcode.OutStdout {
		redirs = append(redirs, "2>&1")
	}

	switch {
	case p.Stdin == nil && background:
		// without it, a background job reads /dev/null
		redirs = append(redirs, "<&0")
	case p.Stdin == nil:
	case p.Stdin.Kind == bcode.InputString:
		redirs = append(redirs, "< "+file(files, i, "i"))
	case p.Stdin.Kind == bcode.InputFile:
		redirs = append(redirs, "< "+b.atom(p.Stdin.Value))
	case p.Stdin.Kind == bcode.InputSource:
		redirs = append(redirs, "< "+file(files, index[p.Stdin.Source], "p"))
	}

	if len(redirs) == 0 {
		return ""
	}

	return " " + strings.Join(redirs, " ")
}

func (b *bash) output(op string, out *bcode.ProcOutput, collected string) string {
	switch out.Kind {
	case bcode.OutCollect:
		return op + " " + collected
	case bcode.OutFile:
		return op + " " + b.atom(out.Path)
	case bcode.OutAppend:
		return op + "> " + b.atom(out.Path)
	}

	return op + " /dev/null"
}

// runBody reads the pipe of the last proc. When the body ends before the pipe's end, the
// proc is stopped. bash 3.2 has no exec {fd}<, the statements nest like a stack, also
// across function calls, so a counter finds a free fd.
func (b *bash) runBody(st *bcode.Cmd, main *bcode.Proc, files string) {
	fd := st.Pipe.Name

	b.print("L_cmd_fd=$(( L_cmd_fd + 1 ))")
	b.print("local " + fd + "=$L_cmd_fd " + fd + "_eof=false")
	b.print(`eval "exec $` + fd + `<\"\$` + files + "." + strconv.Itoa(len(st.Procs)-1) + `p\""`)

	b.active = append(b.active, &activeCmd{cmd: st, files: files, loops: b.loops})
	b.block(st.Body)
	b.active = b.active[:len(b.active)-1]

	// SIGPIPE would be enough for most commands, but Go programs like docker ignore it
	b.print(`[ "$` + fd + `_eof" = true ] || kill -TERM "$` + pid(main) + `" 2>/dev/null`)
	b.print(`eval "exec $` + fd + `<&-"`)
	b.print("L_cmd_fd=$(( L_cmd_fd - 1 ))")
	// the redirection hides the message of bash 3.2 about a job that was terminated
	b.print(`wait "$` + pid(main) + `" 2>/dev/null`)
	b.print("local " + main.Code.Name + "=$?")
	b.print("local " + main.Stopped.Name + "=false")
	b.print(`[ "$` + fd + `_eof" = true ] || ` + main.Stopped.Name + "=true")
}

// waitSources stops the sources after the command they feed has ended, it doesn't read
// them anymore. Without it, docker logs -f would run until its next line. 141 is
// SIGPIPE, a source whose reader went away.
func (b *bash) waitSources(st *bcode.Cmd) {
	sources := st.Procs[:len(st.Procs)-1]

	for _, p := range sources {
		b.print(`kill -TERM "$` + pid(p) + `" 2>/dev/null`)
	}

	for _, p := range sources {
		b.print(`wait "$` + pid(p) + `" 2>/dev/null`)
		b.print("local " + p.Code.Name + "=$?")
		b.print("local " + p.Stopped.Name + "=false")
		b.print(`case "$` + p.Code.Name + `" in 141|143) ` + p.Stopped.Name + "=true ;; esac")
	}
}

// collect reads a collected output after the procs have ended. The file is missing when
// a redirection before it failed, so errors of < are hidden.
func (b *bash) collect(out *bcode.ProcOutput, path string) {
	if out == nil || out.Kind != bcode.OutCollect {
		return
	}

	if !out.Var.Type().IsArray() {
		b.readString(out.Var.Name, path)
		return
	}

	b.print("local -a " + out.Var.Name + "=()")
	b.print("local L_line")
	b.print(`while IFS= read -r L_line || [ -n "$L_line" ]; do ` + appendInPlace(out.Var.Name, `"$L_line"`) + `; done 2>/dev/null < ` + path)

	if out.Text != nil {
		b.readString(out.Text.Name, path)
	}
}

// readString reads the whole file without its last newline, a command substitution would
// remove all of them
func (b *bash) readString(name, path string) {
	b.print("local " + name)
	b.print("IFS= read -r -d '' " + name + " 2>/dev/null < " + path)
	b.print(name + `="${` + name + `%$'\n'}"`)
}

func (b *bash) readLines(r *bcode.ReadLines) {
	fd := r.Pipe.Name
	line := r.Line.Name

	// read -u leaves the loop's stdin alone, a command in the body could read the pipe
	b.print("local " + line)
	b.print(`while IFS= read -r -u "$` + fd + `" ` + line + ` || [ -n "$` + line + `" ] || { ` + fd + `_eof=true; false; }; do`)

	b.loops++
	b.block(r.Block)
	b.loops--

	b.print("done")
}

// stop stops the active cmds from the index from on, the innermost first
func (b *bash) stop(from int) {
	for i := len(b.active) - 1; i >= from; i-- {
		a := b.active[i]

		words := []string{"lsh__cmdStop", `"$` + a.cmd.Pipe.Name + `"`}
		for _, p := range a.cmd.Procs {
			words = append(words, `"$`+pid(p)+`"`)
		}

		b.useFunc("cmdStop")
		b.print(strings.Join(words, " "))
		b.print(`rm -f "$` + a.files + `".*`)
	}
}

// stopForJump stops the cmds that a break or continue of the innermost loop leaves
func (b *bash) stopForJump() {
	for i, a := range b.active {
		if a.loops >= b.loops {
			b.stop(i)
			return
		}
	}
}
