package bash

type bFunc struct {
	name string
	code string
}

var funcs = map[string]*bFunc{}

func register(f *bFunc) {
	funcs[f.name] = f
}

func init() {
	// both exit with 2, like a panic in Go
	register(&bFunc{
		name: "checkIndex",
		code: `if (( $1 < 0 || $1 >= $2 )); then
	printf '%s index out of range [%s] with length %s\n' "$3" "$1" "$2" >&2
	exit 2
fi`,
	})

	register(&bFunc{
		name: "checkSlice",
		code: `if (( $1 < 0 || $1 > $2 || $2 > $3 )); then
	printf '%s slice bounds out of range [%s:%s] with length %s\n' "$4" "$1" "$2" "$3" >&2
	exit 2
fi`,
	})

	register(&bFunc{
		name: "println",
		code: `local i=1
for (( ; i<=$#; i++ )); do
	if [ $i -gt 1 ]; then
		printf " "
	fi
	printf -- "%s" "${!i}"
done
printf "\n"`,
	})
}
