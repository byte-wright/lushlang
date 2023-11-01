package bash

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/byte-wright/lush/internal/bcode"
)

type bash struct {
	p      *bcode.Program
	lines  []string
	indent int
}

func Translate(p *bcode.Program) string {
	b := &bash{
		p: p,
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

lsh__println() {
	local i=1
	for (( ; i<=$#; i++ )); do
		if [ $i -gt 1 ]; then
			printf " "
		fi
		printf -- "%s" "${!i}"
	done
	printf "\n"
}


lsh__join() {
	local i=2
	local v=""
	local di=$(($1+2))
	for (( i=2;i<$di;i++ )); do
		if [ $i -eq 2 ]; then
			v="${!i}"
		else
			v="$v${!di}${!i}"
		fi
	done
	lsh__funcretparam="${v}"
}

lsh__hasPrefix() {
	if [[ "$1" == "$2"* ]]; then
		lsh__funcretparam=true
	else
		lsh__funcretparam=false
	fi
}

lsh__hasSuffix() {
	if [[ "$1" == *"$2" ]]; then
		lsh__funcretparam=true
	else
		lsh__funcretparam=false
	fi
}

lsh__trimPrefix() {
	lsh__funcretparam="${1#"$2"}"
}

lsh__trimSuffix() {
	lsh__funcretparam="${1%"$2"}"
}

lsh__len() {
	lsh__funcretparam=${#1}
}

lsh__indexOf() {
	local pos=$(awk -v a="$1" -v b="$2" 'BEGIN{print index(a,b)}')
	if [ "$pos" -eq 0 ]; then
	lsh__funcretparam=-1
	else
		lsh__funcretparam=$((pos - 1))
	fi
}

lsh__substring() {
	local i2=$(($3-$2))
	lsh__funcretparam="${1:$2:$i2}"
}

`)

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
			fmt.Println(fmt.Sprintf("no valid statement %T", c))
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
		b.print("lsh__substring " + b.atom(at.Value) + " " + at.From.Print() + " " + at.To.Print())
		return "\"$lsh__funcretparam\""

	case *bcode.Index:
		return "\"${" + at.Value.Name + "[" + b.atom(at.Position) + "]}\""

	case *bcode.NumberValue:
		return strconv.Itoa(at.Value)

	case *bcode.StringValue:
		return "'" + at.Value + "'"

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

func (b *bash) print(line string) {
	prefix := strings.Repeat("  ", b.indent)
	b.lines = append(b.lines, prefix+line)
}
