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

	// Every ${s:i:1} counts the characters of the whole string, so one per character
	// would be quadratic. Cutting off chunks first makes that n²/300.
	register(&bFunc{
		name: "strChars",
		code: `local LC_ALL="$L_utf8"
local L_s="$1" L_c L_i L_n=0
L_ret0=()
while [ -n "$L_s" ]; do
	L_c="${L_s:0:300}"
	L_s="${L_s:300}"
	for (( L_i = 0; L_i < ${#L_c}; L_i++ )); do
		L_ret0[L_n]="${L_c:L_i:1}"
		L_n=$(( L_n + 1 ))
	done
done`,
		uses: []string{"findUtf8"},
	})

	// The text never reaches $(( )): bash would evaluate it as an expression, a[$(cmd)]
	// runs cmd, a leading 0 means octal and too many digits wrap around. 19 digits are
	// compared to the int range in two parts that each fit an int. The digits are listed,
	// a range like 0-9 depends on the locale.
	register(&bFunc{
		name: "parseInt",
		code: `local L_s="$1" L_sign=""
case "$L_s" in
	-*) L_sign=-; L_s="${L_s#-}" ;;
	+*) L_s="${L_s#+}" ;;
esac
case "$L_s" in
	""|*[!0123456789]*)
		L_err="\"$1\" is not an int"
		return 1
		;;
esac
L_s="${L_s#"${L_s%%[!0]*}"}"
if [ -z "$L_s" ]; then
	L_ret0=0
	return 0
fi
if [ "${#L_s}" -gt 19 ]; then
	L_err="\"$1\" is out of the int range"
	return 1
fi
if [ "${#L_s}" -eq 19 ]; then
	local L_max=854775807
	if [ "$L_sign" = - ]; then
		L_max=854775808
	fi
	if (( ${L_s:0:10} > 9223372036 || ${L_s:0:10} == 9223372036 && 10#${L_s:10} > L_max )); then
		L_err="\"$1\" is out of the int range"
		return 1
	fi
fi
L_ret0="$L_sign$L_s"
return 0`,
	})

	// fileError sets L_err for a file of std/files that can't be opened, $1 is read,
	// write or append. The reason is the end of bash's own message when it opens the
	// file again, the C locale makes it the same everywhere.
	register(&bFunc{
		name: "fileError",
		code: `local LC_ALL=C L_m
case "$1" in
	read)
		L_m=$( { : < "$2"; } 2>&1 )
		L_err="can't read $2"
		;;
	write)
		L_m=$( { : >> "$2"; } 2>&1 )
		L_err="can't write $2"
		;;
	*)
		L_m=$( { : >> "$2"; } 2>&1 )
		L_err="can't append to $2"
		;;
esac
if [ -n "$L_m" ]; then
	L_err="$L_err: ${L_m##*: }"
fi`,
	})

	// The temp dir holds the collected output and the FIFOs of cmd. The fds of the pipes
	// that lush reads start at 3.
	register(&bFunc{
		name: "cmdInit",
		code: `if [ -n "$L_cmd_dir" ]; then
	return
fi
L_cmd_dir=$(mktemp -d "${TMPDIR:-/tmp}/lush.XXXXXX") || exit 2
L_cmd_seq=0
L_cmd_fd=2
trap lsh__cmdExit EXIT`,
		uses: []string{"cmdExit"},
	})

	// A job that gets a signal right after its fork, before bash resets its traps, runs
	// this too. Its copy of L_cmd_fork is set, so it leaves the script's jobs and files
	// alone. The exec hides the messages of bash 3.2 about the stopped jobs, the script's
	// own messages are already printed.
	register(&bFunc{
		name: "cmdExit",
		code: `if [ -n "$L_cmd_fork" ]; then
	return
fi
exec 2>/dev/null
local L_p
for L_p in $(jobs -p); do
	kill -TERM "$L_p"
done
rm -rf "$L_cmd_dir"`,
	})

	// cmdStop stops the procs of a cmd whose body is left early, $1 is the fd of its pipe
	register(&bFunc{
		name: "cmdStop",
		code: `local L_fd="$1" L_p
shift
for L_p in "$@"; do
	kill -TERM "$L_p" 2>/dev/null
done
eval "exec $L_fd<&-"
L_cmd_fd=$(( L_cmd_fd - 1 ))
for L_p in "$@"; do
	wait "$L_p" 2>/dev/null
done`,
	})

	register(&bFunc{
		name: "print",
		code: `local i=1
for (( ; i<=$#; i++ )); do
	if [ $i -gt 1 ]; then
		printf " "
	fi
	printf -- "%s" "${!i}"
done`,
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

	register(&bFunc{
		name: "eprint",
		code: `lsh__print "$@" >&2`,
		uses: []string{"print"},
	})

	register(&bFunc{
		name: "eprintln",
		code: `lsh__println "$@" >&2`,
		uses: []string{"println"},
	})
}
