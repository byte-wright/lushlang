package bash

import (
	"fmt"
	"regexp"
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
	// canFail is set while translating a function that can fail, and main
	canFail bool
	// active are the cmds whose bodies are around the current command, loops is the
	// number of loops around it, both in the current function
	active []*activeCmd
	loops  int
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
		b.canFail = gd.CanFail
		b.body(gd.Body)
		b.print("}")
	}

	b.print("function main() {")

	// main fails like a function, the last line ends the script with its error
	b.canFail = true
	b.body(b.p.Main)

	b.print("}")

	// generated variables start with L_: identifiers can't start with an uppercase
	// letter, and none of bash's own variables starts with L_
	// inside of functions, $@ are the function's arguments
	b.print(`
L_bool_0=true
L_bool_1=false
L_args=("$@")
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

	// before main, a script stops before it does anything when it can't run
	if b.usedFuncs["findUtf8"] {
		b.print("lsh__findUtf8")
	}

	b.print(`main "$@" || { printf '%s\n' "$L_err" >&2; exit 1; }`)
}

// body translates the block of a function. A function that can fail returns 1 when it
// fails, so every other way out of it returns 0, and not the status of its last command.
func (b *bash) body(block *bcode.Block) {
	b.block(block)

	if !b.canFail || len(block.Commands) > 0 && leaves(block.Commands[len(block.Commands)-1]) {
		return
	}

	b.indent++
	b.print("return 0")
	b.indent--
}

func leaves(cmd bcode.Command) bool {
	switch cmd.(type) {
	case *bcode.Return, *bcode.Fail:
		return true
	}

	return false
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
			// declare -g needs bash 4.2, and no local can hide a G_ name, so a plain
			// assignment reaches the global from any function
			scope := "local "
			if cmd.Global {
				scope = ""
			}

			b.print(scope + cmd.Var.Name + "=" + b.atom(cmd.Value))

		case *bcode.SetElement:
			// bash would add an element past the end, and leave a gap before it
			b.useFunc("checkIndex")
			b.print("lsh__checkIndex " + b.atom(cmd.Index) + " " + arrayLength(cmd.Array) + " " + b.location(cmd.Location))
			b.print(cmd.Array.Name + "[" + b.arith(cmd.Index) + "]=" + b.atom(cmd.Value))

		case *bcode.Func:
			call := b.call(cmd)
			if !cmd.Try {
				b.print(call)
			} else if len(b.active) == 0 {
				// L_err still holds the message of the failed call
				b.print(call + " || return 1")
			} else {
				b.print(call + " || {")
				b.indent++
				b.stop(0)
				b.print("return 1")
				b.indent--
				b.print("}")
			}

			b.results(cmd)

		case *bcode.Catch:
			// with nothing to run after a successful call, the else would be empty
			if len(cmd.Call.Return) == 0 && len(cmd.Then.Commands) == 0 {
				b.print("if ! " + b.call(cmd.Call) + "; then")
			} else {
				b.print("if " + b.call(cmd.Call) + "; then")
				b.indent++
				b.results(cmd.Call)
				b.indent--
				if len(cmd.Then.Commands) > 0 {
					b.block(cmd.Then)
				}
				b.print("else")
			}

			b.block(cmd.Block)
			b.print("fi")

		case *bcode.Fail:
			b.print("L_err=" + b.atom(cmd.Message))
			b.stop(0)
			b.print("return 1")

		case *bcode.If:
			b.print("if [ " + b.atom(cmd.Condition) + " == true ]; then")

			b.block(cmd.Block)

			if cmd.Else != nil {
				b.print("else")
				b.block(cmd.Else)
			}

			b.print("fi")

		case *bcode.While:
			if cmd.Condition == nil {
				b.print("while true; do")
			} else {
				// the check is in the condition list, so continue runs it too
				b.print("while")
				b.block(cmd.Check)
				b.indent++
				b.print("[ " + b.atom(cmd.Condition) + " == true ]")
				b.indent--
				b.print("do")
			}

			b.loops++
			b.block(cmd.Block)
			b.loops--

			b.print("done")

		case *bcode.ForEach:
			// the words of a for are expanded once, that reads the elements before the loop
			b.print("local " + cmd.Element.Name)
			b.print("for " + cmd.Element.Name + " in \"${" + cmd.Array.Name + "[@]}\"; do")

			b.loops++
			b.block(cmd.Block)
			b.loops--

			b.print("done")

		case *bcode.Break:
			b.stopForJump()
			b.print("break")

		case *bcode.Continue:
			b.stopForJump()
			b.print("continue")

		case *bcode.Return:
			for i, val := range cmd.Values {
				b.print("L_ret" + strconv.Itoa(i) + "=" + b.atom(val))
			}
			b.stop(0)

			if b.canFail {
				b.print("return 0")
			} else {
				b.print("return")
			}

		case *bcode.Cmd:
			b.cmd(cmd)

		case *bcode.ReadLines:
			b.readLines(cmd)

		case *bcode.Params:
			b.params(cmd.Vars)

		case *bcode.Code:
			b.useHelpersIn(cmd.Code)
			b.print(cmd.Code)

		case *bcode.Append:
			params := []string{}

			for _, p := range cmd.Elements {
				params = append(params, b.atom(p))
			}
			arr := cmd.Array.(*bcode.VarValue)

			if arr.Name == cmd.Target.Name {
				for _, p := range params {
					b.print(appendInPlace(arr.Name, p))
				}
			} else {
				b.print("local " + cmd.Target.Name + `=("${` + arr.Name + `[@]}" ` + strings.Join(params, " ") + `)`)
			}

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
		return "$(( " + b.arith(at.Left) + " + " + b.arith(at.Right) + " ))"

	case *bcode.Sub:
		return "$(( " + b.arith(at.Left) + " - " + b.arith(at.Right) + " ))"

	case *bcode.Minus:
		return "$(( -" + b.arith(at.Expression) + " ))"

	case *bcode.Mul:
		return "$(( " + b.arith(at.Left) + " * " + b.arith(at.Right) + " ))"

	case *bcode.Div:
		return "$(( " + b.arith(at.Left) + " / " + b.arith(at.Right) + " ))"

	case *bcode.Mod:
		return "$(( " + b.arith(at.Left) + " % " + b.arith(at.Right) + " ))"

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

	// the str helpers return in L_ret0, it's safe to read because bcode assigns
	// every Len, Index, Slice and Chars to a variable right away
	case *bcode.Chars:
		b.useFunc("strChars")
		b.print("lsh__strChars " + b.atom(at.Value))
		return `("${L_ret0[@]}")`

	case *bcode.Slice:
		from, to := b.atom(at.From), b.atom(at.To)
		if !at.Value.Type().IsArray() {
			b.useFunc("strSlice")
			b.print("lsh__strSlice " + b.atom(at.Value) + " " + from + " " + to + " " + b.location(at.Location))
			return `"$L_ret0"`
		}

		b.useFunc("checkSlice")
		b.print("lsh__checkSlice " + from + " " + to + " " + arrayLength(at.Value) + " " + b.location(at.Location))

		bounds := ":$(( " + b.arith(at.From) + " )):$(( " + b.arith(at.To) + " - " + b.arith(at.From) + " ))}"
		return "(\"${" + at.Value.Name + "[@]" + bounds + "\")"

	case *bcode.Len:
		if !at.Parameter.Type().IsArray() {
			b.useFunc("strLen")
			b.print("lsh__strLen " + b.atom(at.Parameter))
			return `"$L_ret0"`
		}
		return arrayLength(at.Parameter)

	case *bcode.Index:
		position := b.atom(at.Position)
		if !at.Value.Type().IsArray() {
			b.useFunc("strIndex")
			b.print("lsh__strIndex " + b.atom(at.Value) + " " + position + " " + b.location(at.Location))
			return `"$L_ret0"`
		}

		b.useFunc("checkIndex")
		b.print("lsh__checkIndex " + position + " " + arrayLength(at.Value) + " " + b.location(at.Location))
		return "\"${" + at.Value.Name + "[" + position + "]}\""

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

	case *bcode.ArgsValue:
		return `("${L_args[@]}")`

	case *bcode.ErrorValue:
		return `"$L_err"`

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

// bash 3.2 doesn't remove double quotes in $(( )), "$a" + 1 is a syntax error there
func (b *bash) arith(a bcode.Atom) string {
	if v, isVar := a.(*bcode.VarValue); isVar {
		return v.Name
	}
	return b.atom(a)
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

func appendInPlace(array, value string) string {
	// a+=(x) is quadratic on bash 3.2
	return array + "[${#" + array + "[@]}]=" + value
}

func arrayLength(v *bcode.VarValue) string {
	return "\"${#" + v.Name + "[@]}\""
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

func (b *bash) call(f *bcode.Func) string {
	return b.funcName(f) + " " + b.mkFuncParams(f)
}

// results copies the results of a call out of L_ret, the next call overwrites them
func (b *bash) results(f *bcode.Func) {
	for i, rpn := range f.Return {
		ret := &bcode.VarValue{Name: "L_ret" + strconv.Itoa(i), T: rpn.T}
		b.print("local " + rpn.Name + "=" + b.atom(ret))
	}
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
	if b.usedFuncs[name] {
		return
	}
	b.usedFuncs[name] = true

	f, isHelper := funcs[name]
	if isHelper {
		for _, u := range f.uses {
			b.useFunc(u)
		}
	}
}

var helperCall = regexp.MustCompile(`\blsh__([A-Za-z0-9]+)`)

// useHelpersIn finds the helpers that the bash body of an external function calls,
// like lsh__strLen in std.
func (b *bash) useHelpersIn(code string) {
	for _, m := range helperCall.FindAllStringSubmatch(code, -1) {
		_, isHelper := funcs[m[1]]
		if isHelper {
			b.useFunc(m[1])
		}
	}
}

func (b *bash) print(line string) {
	prefix := strings.Repeat("  ", b.indent)
	b.lines = append(b.lines, prefix+line)
}
