package bash

type bFunc struct {
	name string
	code string
	uses []string
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

	// Strings are UTF-8 and count characters, whatever the user's locale. Only the
	// str helpers switch to a UTF-8 locale: a global LC_ALL would reach the commands
	// the script runs and change what $LC_ALL reads. When LC_ALL is exported, the
	// local copy is exported too, so these helpers must not run commands.
	register(&bFunc{
		name: "findUtf8",
		code: `local L_c L_t=$'\xc3\xa9'
for L_c in C.UTF-8 C.utf8 en_US.UTF-8 UTF-8 "${LC_ALL:-${LC_CTYPE:-$LANG}}"; do
	if [ -n "$L_c" ]; then
		# hides bash's warning for a locale that doesn't exist
		local LC_ALL="$L_c" 2>/dev/null
		if [ "${#L_t}" -eq 1 ]; then
			L_utf8="$L_c"
			return
		fi
	fi
done
printf 'string operations need a UTF-8 locale, none found (tried C.UTF-8, C.utf8, en_US.UTF-8, UTF-8 and the current locale)\n' >&2
exit 2`,
	})

	register(&bFunc{
		name: "strLen",
		code: `local LC_ALL="$L_utf8"
L_ret0="${#1}"`,
		uses: []string{"findUtf8"},
	})

	register(&bFunc{
		name: "strIndex",
		code: `local LC_ALL="$L_utf8"
lsh__checkIndex "$2" "${#1}" "$3"
L_ret0="${1:$2:1}"`,
		uses: []string{"findUtf8", "checkIndex"},
	})

	register(&bFunc{
		name: "strSlice",
		code: `local LC_ALL="$L_utf8"
lsh__checkSlice "$2" "$3" "${#1}" "$4"
L_ret0="${1:$2:$(( $3 - $2 ))}"`,
		uses: []string{"findUtf8", "checkSlice"},
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
