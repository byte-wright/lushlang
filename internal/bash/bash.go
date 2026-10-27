package bash

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/byte-wright/lush/internal/bcode"
	"github.com/byte-wright/lush/internal/common"
)

type bash struct {
	p         *bcode.Program
	lines     []string
	indent    int
	usedFuncs map[string]bool
}

func Translate(p *bcode.Program) string {
	b := &bash{
		p:         p,
		usedFuncs: map[string]bool{},
	}

	b.translate()

	return strings.Join(b.lines, "\n")
}

func (b *bash) translate() {
	b.print("#!/bin/bash")

	for _, gd := range b.p.Funcs {
		b.print("function lsh_" + gd.FullName() + "() {")
		b.block(gd.Body)
		b.print("}")
	}

	b.print("function main() {")

	b.block(b.p.Main)

	b.print("}")

	// generated variables start with L_: identifiers can't start with an uppercase
	// letter, and none of bash's own variables starts with L_
	b.print(`
L_bool_0=true
L_bool_1=false
`)

	ufncs := []string{}
	for fn := range b.usedFuncs {
		ufncs = append(ufncs, fn)
	}

	sort.Strings(ufncs)

	for _, fn := range ufncs {
		iFunc, isInternal := funcs[fn]
		if !isInternal {
			panic("func not found " + fn)
		}

		b.print("lsh__" + fn + "() {")
		b.print(iFunc.code)
		b.print("}\n")
	}

	b.print(`main "$@"`)
}

func (b *bash) block(block *bcode.Block) {
	b.indent++
	defer func() { b.indent-- }()

	// bash rejects empty function and if bodies
	if len(block.Commands) == 0 {
		b.print(":")
	}

	for _, c := range block.Commands {
		switch cmd := c.(type) {
		case *bcode.Assignment:
			b.print("local " + cmd.Var.Name + "=" + b.atom(cmd.Value))

		case *bcode.Func:
			b.print(b.funcName(cmd) + " " + b.mkFuncParams(cmd))
			for i, rpn := range cmd.Return {
				ret := &bcode.VarValue{Name: "L_ret" + strconv.Itoa(i), T: rpn.T}
				b.print("local " + rpn.Name + "=" + b.atom(ret))
			}

		case *bcode.If:
			b.print("if [ " + b.atom(cmd.Condition) + " == true ]; then")

			b.block(cmd.Block)

			if cmd.Else != nil {
				b.print("else")
				b.block(cmd.Else)
			}

			b.print("fi")

		case *bcode.While:
			b.print("while [ " + b.atom(cmd.Condition) + " == true ]; do")

			b.block(cmd.Block)

			b.print("done")

		case *bcode.Return:
			for i, val := range cmd.Values {
				b.print("L_ret" + strconv.Itoa(i) + "=" + b.atom(val))
			}
			b.print("return")

		case *bcode.ExecVar:
			cmds := []string{b.atom(cmd.Command)}

			for _, p := range cmd.Parameters {
				cmds = append(cmds, b.atom(p))
			}

			cmdl := strings.Join(cmds, " ")

			// stderr goes through a file because a process substitution runs in a
			// subshell and can't set our variables. The exit code must be read
			// right after the command substitution, local would overwrite $?.
			b.print("local L_stdout L_stderr")
			b.print("L_stderr=$(mktemp)")
			b.print("L_stdout=$(" + cmdl + " 2> \"$L_stderr\")")
			b.print("local " + cmd.Err.Name + "=$?")
			b.print("local -a " + cmd.Stdout.Name + " " + cmd.Stderr.Name)
			// printf instead of <<<, which would turn no output into one empty line
			b.print("readarray -t " + cmd.Stdout.Name + " < <(printf '%s' \"$L_stdout\")")
			b.print("readarray -t " + cmd.Stderr.Name + " < \"$L_stderr\"")
			b.print("rm -f \"$L_stderr\"")

		case *bcode.Params:
			b.params(cmd.Vars)

		case *bcode.Code:
			b.print(cmd.Code)

		case *bcode.Append:
			params := []string{}

			for _, p := range cmd.Elements {
				params = append(params, b.atom(p))
			}
			arr := cmd.Array.(*bcode.VarValue)

			b.print(`if [ "${#` + arr.Name + `[@]}" -eq 0 ]; then`)
			b.print(`  local L_tmp_array=(` + strings.Join(params, " ") + `)`)
			b.print("else")
			b.print(`  local L_tmp_array=("${` + arr.Name + `[@]}" ` + strings.Join(params, " ") + `)`)
			b.print("fi")
			b.print("local " + cmd.Target.Name + `=("${L_tmp_array[@]}")`)

		default:
			fmt.Printf("no valid statement in bash transpiler %T\n", c)
		}
	}
}

func (b *bash) atom(a bcode.Atom) string {
	switch at := a.(type) {
	case *bcode.Add:
		if at.Left.Type().Is(&common.BasicType{Type: common.String}) {
			// both atoms are quoted words, putting them side by side concatenates them
			return b.atom(at.Left) + b.atom(at.Right)
		}
		return "$(( " + b.atom(at.Left) + " + " + b.atom(at.Right) + " ))"

	case *bcode.Sub:
		return "$(( " + b.atom(at.Left) + " - " + b.atom(at.Right) + " ))"

	case *bcode.Minus:
		return "$(( -" + b.atom(at.Expression) + " ))"

	case *bcode.Mul:
		return "$(( " + b.atom(at.Left) + " * " + b.atom(at.Right) + " ))"

	case *bcode.Div:
		return "$(( " + b.atom(at.Left) + " / " + b.atom(at.Right) + " ))"

	case *bcode.Mod:
		return "$(( " + b.atom(at.Left) + " % " + b.atom(at.Right) + " ))"

	case *bcode.Equal:
		b.print("[ " + b.atom(at.Left) + " " + equalOp(at.Left) + " " + b.atom(at.Right) + " ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.NotEqual:
		b.print("! [ " + b.atom(at.Left) + " " + equalOp(at.Left) + " " + b.atom(at.Right) + " ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.LessThan:
		b.print("[ " + b.atom(at.Left) + " -lt " + b.atom(at.Right) + " ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.LessThanEqual:
		b.print("[ " + b.atom(at.Left) + " -le " + b.atom(at.Right) + " ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.GreaterThan:
		b.print("[ " + b.atom(at.Left) + " -gt " + b.atom(at.Right) + " ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.GreaterThanEqual:
		b.print("[ " + b.atom(at.Left) + " -ge " + b.atom(at.Right) + " ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.Not:
		b.print("[ " + b.atom(at.Expression) + " == false ]")
		b.print("local L_i=\"L_bool_$?\"")
		return "${!L_i}"

	case *bcode.Slice:
		from, to := b.atom(at.From), b.atom(at.To)
		b.useFunc("checkSlice")
		b.print("lsh__checkSlice " + from + " " + to + " " + length(at.Value) + " " + b.location(at.Location))

		// bash doesn't accept a quoted offset, "${a[@]:"$i"}" is a syntax error
		bounds := ":$(( " + from + " )):$(( " + to + " - " + from + " ))}"
		if at.Value.Type().IsArray() {
			return "(\"${" + at.Value.Name + "[@]" + bounds + "\")"
		}
		return "\"${" + at.Value.Name + bounds + "\""

	case *bcode.Len:
		return length(at.Parameter)

	case *bcode.Index:
		position := b.atom(at.Position)
		b.useFunc("checkIndex")
		b.print("lsh__checkIndex " + position + " " + length(at.Value) + " " + b.location(at.Location))

		if at.Value.Type().IsArray() {
			return "\"${" + at.Value.Name + "[" + position + "]}\""
		}
		return "\"${" + at.Value.Name + ":$(( " + position + " )):1}\""

	case *bcode.NumberValue:
		return strconv.Itoa(at.Value)

	case *bcode.StringValue:
		return "'" + strings.ReplaceAll(at.Value, "'", "'\\''") + "'"

	case *bcode.BoolValue:
		return strconv.FormatBool(at.Value)

	case *bcode.VarValue:
		if at.Type().IsArray() {
			return "(\"${" + at.Name + "[@]}\")"
		}
		return "\"$" + at.Name + "\""

	case *bcode.EnvVarValue:
		return "\"$" + at.Name + "\""

	case *bcode.ArrayValue:
		params := []string{}

		for _, p := range at.Values {
			params = append(params, b.atom(p))
		}

		return "(" + strings.Join(params, " ") + ")"

	default:
		fmt.Printf("no valid atom %T\n", at)
	}
	return "unknown"
}

// params binds the positional arguments. mkFuncParams passes an array as its
// length followed by its elements, so every later position depends on that length.
func (b *bash) params(vars []*bcode.VarValue) {
	pos := 1
	lengths := ""

	for _, v := range vars {
		at := strconv.Itoa(pos) + lengths

		if v.Type().IsArray() {
			length := "L_len_" + v.Name
			b.print("local " + length + "=\"${@:" + at + ":1}\"")
			b.print("local " + v.Name + "=(\"${@:" + at + "+1:" + length + "}\")")
			lengths += "+" + length
		} else if lengths == "" {
			b.print("local " + v.Name + "=\"${" + at + "}\"")
		} else {
			b.print("local " + v.Name + "=\"${@:" + at + ":1}\"")
		}

		pos++
	}
}

func length(v *bcode.VarValue) string {
	if v.Type().IsArray() {
		return "\"${#" + v.Name + "[@]}\""
	}
	return "\"${#" + v.Name + "}\""
}

func (b *bash) location(l *common.Location) string {
	return b.atom(&bcode.StringValue{Value: l.String()})
}

// equalOp picks the test operator for == and !=, -eq only accepts integers.
func equalOp(v bcode.Value) string {
	if v.Type().Is(&common.BasicType{Type: common.Int}) {
		return "-eq"
	}
	return "="
}

func (b *bash) mkFuncParams(f *bcode.Func) string {
	params := []string{}

	for _, p := range f.Parameters {
		if p.Type().IsArray() {
			switch v := p.(type) {
			case *bcode.VarValue:
				params = append(params, "${#"+v.Name+"[@]}")
				params = append(params, "\"${"+v.Name+"[@]}\"")
			default:
				panic("arrays must always be var values")
			}
		} else {
			params = append(params, b.atom(p))
		}
	}

	return strings.Join(params, " ")
}

// funcName gives helpers a double underscore, a namespace never starts with one, so
// user functions can't replace them.
func (b *bash) funcName(f *bcode.Func) string {
	if f.Namespace == "" {
		b.useFunc(f.Name)
		return "lsh__" + f.Name
	}

	return "lsh_" + f.FullName()
}

func (b *bash) useFunc(name string) {
	b.usedFuncs[name] = true
}

func (b *bash) print(line string) {
	prefix := strings.Repeat("  ", b.indent)
	b.lines = append(b.lines, prefix+line)
}
