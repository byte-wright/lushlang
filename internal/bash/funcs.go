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

	// The text never reaches $(( )): bash would evaluate it as an expression, a[$(cmd)]
	// runs cmd, a leading 0 means octal and too many digits wrap around. 19 digits are
	// compared to the int range in two parts that each fit an int. The digits are listed,
	// a range like 0-9 depends on the locale.
	register(&bFunc{
		name: "parseInt",
		code: `local L_s="$1" L_sign=""
L_ret0=0
L_ret1=false
case "$L_s" in
	-*) L_sign=-; L_s="${L_s#-}" ;;
	+*) L_s="${L_s#+}" ;;
esac
case "$L_s" in
	""|*[!0123456789]*) return ;;
esac
L_s="${L_s#"${L_s%%[!0]*}"}"
if [ -z "$L_s" ]; then
	L_ret1=true
	return
fi
if [ "${#L_s}" -gt 19 ]; then
	return
fi
if [ "${#L_s}" -eq 19 ]; then
	local L_max=854775807
	if [ "$L_sign" = - ]; then
		L_max=854775808
	fi
	if (( ${L_s:0:10} > 9223372036 || ${L_s:0:10} == 9223372036 && 10#${L_s:10} > L_max )); then
		return
	fi
fi
L_ret0="$L_sign$L_s"
L_ret1=true`,
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
