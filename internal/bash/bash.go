package bash

import (
	"fmt"
	"log"
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
		b.print("lsh__" + fn + "() {")
		b.print(funcs[fn].code)
		b.print("}\n")
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
				log.Println("append is a nop when not using the result")
				continue
			}

			b.useFunc(cmd.Name)
			b.print("lsh__" + cmd.Name + " " + b.mkFuncParams(cmd))

		case *bcode.If:
			b.print("if [ " + b.atom(cmd.Condition) + " == true ]; then")

			b.block(cmd.Block)

			b.print("fi")

		case *bcode.While:
			b.print("while [ " + b.atom(cmd.Condition) + " == true ]; do")

			b.block(cmd.Block)

			b.print("done")

		default:
			fmt.Printf("no valid statement %T\n", c)
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
		case *bcode.BasicType:
			if tp.Type == common.String {
				b.useFunc("substring")
				b.print("lsh__substring " + b.atom(at.Value) + " " + b.atom(at.From) + " " + b.atom(at.To))
				return "\"$lsh__funcretparam\""
			}

		case *bcode.ArrayType:
			return "(\"${" + at.Value.Name + "[@]:" + b.atom(at.From) + ":$(( " + b.atom(at.To) + "-" + b.atom(at.From) + " ))}\")"
		}
		panic(fmt.Sprintf("invalid slice type %T", at.Value.Type()))

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

	case *bcode.Func:
		if at.Name == "append" {
			params := []string{}

			for _, p := range at.Parameters[1:] {
				params = append(params, b.atom(p))
			}
			arr := at.Parameters[0].(*bcode.VarValue)

			b.print(`if [ "${#` + arr.Name + `[@]}" -eq 0 ]; then`)
			b.print(`  local lsh__tmp_array=(` + strings.Join(params, " ") + `)`)
			b.print("else")
			b.print(`  local lsh__tmp_array=("${` + arr.Name + `}" ` + strings.Join(params, " ") + `)`)
			b.print("fi")

			return `("${lsh__tmp_array[@]}")`
		}

		b.useFunc(at.Name)
		b.print("lsh__" + at.Name + " " + b.mkFuncParams(at))
		return "\"$lsh__funcretparam\""

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
