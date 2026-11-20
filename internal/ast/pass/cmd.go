package astpass

import (
	"github.com/byte-wright/lush/internal/ast"
	"github.com/byte-wright/lush/internal/common"
)

var pipeType = &common.PipeType{}

// cmd checks a cmd after the ones in its arguments and stdin, whose types are known then.
func (t *setTypesPass) cmd(c *ast.Cmd) {
	t.cmdArgs(c)
	t.cmdOptions(c)

	c.ReturnTypes = cmdResults(c)
	c.CanFail = c.Exit == ast.ExitFails || c.Exit == ast.ExitBool
	for _, src := range sources(c) {
		c.CanFail = c.CanFail || src.CanFail
	}

	if isPipe(c.Stdout) && isPipe(c.Stderr) {
		t.errorf(c.Location, "lush reads one pipe per cmd, stdout: and stderr: can't both be pipe")
	}

	switch c.Use {
	case ast.UseSource:
		t.source(c)
		// its error is the one of the cmd it feeds
		return

	case ast.UseLoop:
		if len(c.ReturnTypes) > 0 {
			t.errorf(c.Location, "a cmd in a for can't have results, get them with the block after the call: cmd(...) { for line in %v { } }", c.Pipe())
		}

	default:
		switch {
		case c.Body != nil && c.Pipe() == "":
			t.errorf(c.Location, "the block after cmd reads its pipe, it needs stdout: pipe or stderr: pipe")

		case c.Body == nil && c.Pipe() != "":
			t.errorf(c.Location, "the pipe of cmd isn't read, read it with for line in cmd(...), in the block after the call or as stdin: of another cmd")

		case c == t.statementCmd && len(c.ReturnTypes) > 0:
			t.errorf(c.Location, "the results of cmd aren't used, assign them or discard the streams with _")

		case c != t.statementCmd && len(c.ReturnTypes) == 0:
			t.errorf(c.Location, "cmd doesn't return a value")
			c.ReturnTypes = unknownTypes
			return
		}
	}

	t.failingCall("cmd", c, c.Try, c.CanFail, c.Location)
}

func isPipe(target *ast.Target) bool {
	return target != nil && target.Kind == ast.PipeTarget
}

// sources are the cmds that feed c, as stdin or as arguments
func sources(c *ast.Cmd) []*ast.Cmd {
	srcs := []*ast.Cmd{}

	for _, arg := range c.Args {
		inner, isCmd := arg.Value.(*ast.Cmd)
		if isCmd && inner.Use == ast.UseSource {
			srcs = append(srcs, inner)
		}
	}

	if c.Stdin != nil && c.Stdin.Kind == ast.InputCmd {
		srcs = append(srcs, c.Stdin.Value.(*ast.Cmd))
	}

	return srcs
}

func (t *setTypesPass) cmdArgs(c *ast.Cmd) {
	for i, arg := range c.Args {
		inner, isCmd := arg.Value.(*ast.Cmd)
		if isCmd && inner.Use == ast.UseSource {
			switch {
			case i == 0:
				t.errorf(arg.Location, "the name of the command can't be a pipe")
			case arg.Spread:
				t.errorf(arg.Location, "argument %v of cmd is a pipe, it can't be spread with ...", i+1)
			}
			continue
		}

		tps := arg.Value.Types()
		if !known(tps) {
			continue
		}

		if len(tps) != 1 {
			t.errorf(arg.Location, "argument %v of cmd must be one value, not %v", i+1, printTypes(tps))
			continue
		}

		tp := tps[0]
		switch {
		case i == 0 && arg.Spread:
			t.errorf(arg.Location, "the name of the command can't be spread with ...")

		case i == 0 && !tp.Is(stringType):
			t.errorf(arg.Location, "the name of the command must be a string, not %v", tp.Print())

		case arg.Spread && !tp.IsArray():
			t.errorf(arg.Location, "argument %v of cmd is %v, only an array can be spread with ...", i+1, tp.Print())

		case !arg.Spread && tp.IsArray():
			t.errorf(arg.Location, "argument %v of cmd is %v, spread it with ... to pass its elements", i+1, tp.Print())

		case !arg.Spread && !isOneOf(tp, []common.Type{stringType, intType, boolType}):
			t.errorf(arg.Location, "argument %v of cmd can't be of type %v", i+1, tp.Print())
		}
	}
}

func (t *setTypesPass) cmdOptions(c *ast.Cmd) {
	for _, target := range []*ast.Target{c.Stdout, c.Stderr} {
		if target != nil && target.Path != nil {
			t.expectType(target.Location, "the path of a file", target.Path, stringType)
		}
	}

	if c.Stdin != nil {
		switch c.Stdin.Kind {
		case ast.InputString:
			t.expectType(c.Location, "stdin", c.Stdin.Value, stringType)
		case ast.InputFile:
			t.expectType(c.Location, "the path of a file", c.Stdin.Value, stringType)
		}
	}

	if c.Dir != nil {
		t.expectType(c.Location, "dir", c.Dir, stringType)
	}

	for _, env := range c.Env {
		if env.Name[0] >= '0' && env.Name[0] <= '9' || env.Name == "_" {
			t.errorf(env.Location, "$%v is not an environment variable", env.Name)
		}

		tps := env.Value.Types()
		if known(tps) && (len(tps) != 1 || !isOneOf(tps[0], []common.Type{stringType, intType, boolType})) {
			t.errorf(env.Location, "the value of $%v must be a string, int or bool, not %v", env.Name, printTypes(tps))
		}
	}
}

func cmdResults(c *ast.Cmd) []common.Type {
	tps := []common.Type{}

	for _, target := range []*ast.Target{c.Stdout, c.Stderr} {
		switch {
		case target == nil:
		case target.Kind == ast.CollectString:
			tps = append(tps, stringType)
		case target.Kind == ast.CollectLines:
			tps = append(tps, stringArrayType)
		}
	}

	switch c.Exit {
	case ast.ExitInt:
		tps = append(tps, intType)
	case ast.ExitBool:
		tps = append(tps, boolType)
	}

	return tps
}

// source checks a cmd that feeds another one. It isn't a statement of its own, so it has
// no results.
func (t *setTypesPass) source(c *ast.Cmd) {
	if c.Stderr != nil && (c.Stderr.Kind == ast.CollectString || c.Stderr.Kind == ast.CollectLines) {
		t.errorf(c.Stderr.Location, "stderr: %v isn't possible in a cmd that feeds another one, it has no results", c.Stderr.Print())
	}

	if c.Exit == ast.ExitInt || c.Exit == ast.ExitBool {
		t.errorf(c.Location, "exit: %v isn't possible in a cmd that feeds another one, it has no results", c.Exit.Print())
	}

	if c.Try {
		t.errorf(c.Location, "a cmd that feeds another one takes no try, the outer cmd fails when it fails")
	}
}

func (t *setTypesPass) BeforeCmdBody(c *ast.Cmd) {
	if c.Pipe() != "" {
		t.loopVar(c.Body, &ast.AssignmentTarget{Name: c.Pipe(), Location: c.Location}, pipeType, "pipe")
	}
}

func (t *setTypesPass) AfterCmdBody(c *ast.Cmd) {
	t.block = c.Body.Parent
}

// cmdBreaks reports whether a break in the block after a cmd leaves the loop around it
func cmdBreaks(exps ...ast.Expression) bool {
	for _, exp := range exps {
		c, isCmd := exp.(*ast.Cmd)
		if isCmd && c.Body != nil && breaks(c.Body) {
			return true
		}
	}

	return false
}
