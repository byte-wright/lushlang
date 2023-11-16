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
		b.print("function lsh__" + gd.Name + "() {")
		b.block(gd.Body)
		b.print("}")
	}

	b.print("function main() {")

	b.block(b.p.Main)

	b.print("}")

	b.print(`
lsh_bool_0=true
lsh_bool_1=false
`)

	ufncs := []string{}
	for fn := range b.usedFuncs {
		ufncs = append(ufncs, fn)
	}

	sort.Strings(ufncs)

	for _, fn := range ufncs {
		iFunc, isInternal := funcs[fn]
		if isInternal {
			b.print("lsh__" + fn + "() {")
			b.print(iFunc.code)
			b.print("}\n")

			continue
		}

		_, isDefined := b.p.FuncsByName[fn]
		if isDefined {
			// we added them already
			continue
		}

		panic("func not found " + fn)

	}

	b.print("main $@")
}

func (b *bash) block(block *bcode.Block) {
	b.indent++
	defer func() { b.indent-- }()

	for _, c := range block.Commands {
		switch cmd := c.(type) {
		case *bcode.Assignment:
			b.print("local " + cmd.Var.Name + "=" + b.atom(cmd.Value))

		case *bcode.Func:

			if cmd.Name == "append" {
				params := []string{}

				for _, p := range cmd.Parameters[1:] {
					params = append(params, b.atom(p))
				}
				arr := cmd.Parameters[0].(*bcode.VarValue)

				b.print(`if [ "${#` + arr.Name + `[@]}" -eq 0 ]; then`)
				b.print(`  local lsh__tmp_array=(` + strings.Join(params, " ") + `)`)
				b.print("else")
				b.print(`  local lsh__tmp_array=("${` + arr.Name + `[@]}" ` + strings.Join(params, " ") + `)`)
				b.print("fi")
				b.print("local " + cmd.Return[0].Name + `=("${lsh__tmp_array[@]}")`)

				continue
			}

			b.useFunc(cmd.Name)
			b.print("lsh__" + cmd.Name + " " + b.mkFuncParams(cmd))
			for i, rpn := range cmd.Return {
				b.print("local " + rpn.Name + "=\"$lsh__funcretparam_" + strconv.Itoa(i) + "\"")
			}

		case *bcode.If:
			b.print("if [ " + b.atom(cmd.Condition) + " == true ]; then")

			b.block(cmd.Block)

			b.print("fi")

		case *bcode.While:
			b.print("while [ " + b.atom(cmd.Condition) + " == true ]; do")

			b.block(cmd.Block)

			b.print("done")

		case *bcode.Return:
			for i, val := range cmd.Values {
				b.print("lsh__funcretparam_" + strconv.Itoa(i) + "=" + b.atom(val))
			}
			b.print("return")

		case *bcode.ExecVar:
			cmds := []string{b.atom(cmd.Command)}

			for _, p := range cmd.Parameters {
				cmds = append(cmds, b.atom(p))
			}

			cmdl := strings.Join(cmds, " ")

			b.print("output=$(" + cmdl + " 2> >(readarray -t " + cmd.Stderr.Name + "; printf '%s\n' \"${error_lines[@]}\"))")
			b.print("readarray -t " + cmd.Stdout.Name + " <<< \"$output\"")

		default:
			fmt.Printf("no valid statement in bash transpiler %T\n", c)
		}
	}
}

func (b *bash) atom(a bcode.Atom) string {
	switch at := a.(type) {
	case *bcode.Add:
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
		b.print("[ " + b.atom(at.Left) + " -eq " + b.atom(at.Right) + " ]")
		b.print("local lsh_i=\"lsh_bool_$?\"")
		return "${!lsh_i}"

	case *bcode.NotEqual:
		b.print("! [ " + b.atom(at.Left) + " -eq " + b.atom(at.Right) + " ]")
		b.print("local lsh_i=\"lsh_bool_$?\"")
		return "${!lsh_i}"

	case *bcode.LessThan:
		b.print("[ " + b.atom(at.Left) + " -lt " + b.atom(at.Right) + " ]")
		b.print("local lsh_i=\"lsh_bool_$?\"")
		return "${!lsh_i}"

	case *bcode.LessThanEqual:
		b.print("local lsh_i=\"lsh_bool_$?\"")
		b.print("[ " + b.atom(at.Left) + " -le " + b.atom(at.Right) + " ]")
		return "${!lsh_i}"

	case *bcode.GreaterThan:
		b.print("[ " + b.atom(at.Left) + " -gt " + b.atom(at.Right) + " ]")
		b.print("local lsh_i=\"lsh_bool_$?\"")
		return "${!lsh_i}"

	case *bcode.GreaterThanEqual:
		b.print("local lsh_i=\"lsh_bool_$?\"")
		b.print("[ " + b.atom(at.Left) + " -ge " + b.atom(at.Right) + " ]")
		return "${!lsh_i}"

	case *bcode.And:
		b.print("[ " + b.atom(at.Left) + " == true ] && [ " + b.atom(at.Right) + " == true ]")
		b.print("local lsh_i=\"lsh_bool_$?\"")
		return "${!lsh_i}"

	case *bcode.Not:
		b.print("[ " + b.atom(at.Expression) + " == false ]")
		b.print("local lsh_i=\"lsh_bool_$?\"")
		return "${!lsh_i}"

	case *bcode.Slice:
		switch tp := at.Value.Type().(type) {
		case *common.BasicType:
			if tp.Type == common.String {
				b.useFunc("substring")
				b.print("lsh__substring " + b.atom(at.Value) + " " + b.atom(at.From) + " " + b.atom(at.To))
				return "\"$lsh__funcretparam_0\""
			}

		case *common.ArrayType:
			return "(\"${" + at.Value.Name + "[@]:" + b.atom(at.From) + ":$(( " + b.atom(at.To) + "-" + b.atom(at.From) + " ))}\")"
		}
		panic(fmt.Sprintf("invalid slice type %T", at.Value.Type()))

	case *bcode.Len:
		if at.Parameter.Type().IsArray() {
			return "\"${#" + at.Parameter.Name + "[@]}\""
		}
		return "\"${#" + at.Parameter.Name + "}\""

	case *bcode.Index:
		return "\"${" + at.Value.Name + "[" + b.atom(at.Position) + "]}\""

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

func (b *bash) useFunc(name string) {
	b.usedFuncs[name] = true
}

func (b *bash) print(line string) {
	prefix := strings.Repeat("  ", b.indent)
	b.lines = append(b.lines, prefix+line)
}
