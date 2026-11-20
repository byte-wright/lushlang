lushlang
========

**Disclaimer**: This is a pet project.

A simple language that compiles to bash.

Why?
----

Bash can be a pita. THis is a try at a language that compiles to bash.

It has an intermediate representation, similar to assembly, from where it can be
translated into bash or other languages.

Its in no way optimized and quite rudimentary.

  - `//` line comments
  - math and comparison operators, `+=`, `-=`, `*=`, `/=` and `%=`
  - int, string, bools and type inference
  - a string plus an int or bool takes its text: `"n=" + n`. Only a string on the left
    does that, `1 + 2 + "x"` is an error, `"n=" + 1 + 2` is `n=12` and
    `"n=" + (1 + 2)` is `n=3`
  - `n = try parseInt(s)` reads an int: an optional `+` or `-` and decimal digits, a
    leading 0 doesn't make it octal. Anything else, also a number outside of the 64-bit
    range, is an [error](#errors)
  - strings are UTF-8: `len`, indexes and slices count characters (Unicode code points),
    whatever the user's locale. Bytes that aren't valid UTF-8 count as one character each.
    This needs a UTF-8 locale on the system, which nearly every system has. Without one,
    a script that uses these operations exits with code 2 before it starts
  - arrays (very slow). `a[i] = x` and `a[i] += x` set an element, an index out of range
    ends the script with exit code 2 like reading it. Arrays are values like strings:
    `b = a` and a call copy them, so changing an element only changes that variable
  - functions with params and (multiple) return params
  - errors with `fail`, `try` and `catch`, see [Errors](#errors)
  - if, else if, else
  - loops with `break` and `continue`, see [Loops](#loops)
  - `return` outside of functions ends the script with exit code 0
  - `println(a, b)` writes its arguments to stdout with a space between them and a newline
    at the end, `print(a, b)` does the same without the newline. `eprintln` and `eprint`
    write to stderr
  - environment variables as `$NAME`, the script's arguments as `args()`
  - commands with `cmd`, see [Commands](#commands)
  - names start with a lowercase letter or `_`, so they can't clash with bash's
    variables like `PATH` and `IFS`
  - packages: `import "strings"` from [std](#std), `import "./deploy"` and `import "../lib/log"` from
    the folder of the importing file, and `import "./util/strings" as ustrings` to use another
    name than the folder's
  - global variables per package
  - the scripts need bash 3.2 or newer, so they also run with macOS's `/bin/bash`

Loops
-----

      for { }                      // endless
      for i < 3 { }                // as long as the condition is true
      for i = 0; i < 3; i = i + 1 { }
      for x in ["a", "b"] { }      // x is the element
      for x, i in ["a", "b"] { }   // i is the index
      for c, i in "héllo" { }      // c is a character as a string, i its index
      for x in 7 { }               // 0 to 6
      for x in 3 until 7 { }       // 3 to 6

  - `break` leaves the innermost loop, `continue` goes on with its next iteration. In the
    `for init; cond; step` form, continue runs the step first, like in Go
  - the characters of a string count like `len` and `s[i]`, so `s[i] == c`
  - `for x in` reads the array, the string or the numbers once before the first iteration.
    Changing them in the body doesn't change what the loop goes over, and neither does
    changing `x` or `i`
  - `x` and `i` only exist in the loop. A variable, parameter or outer loop variable with
    the same name is a compile error, so a loop can't change one by accident

Errors
------

A function that can fail has `error` as its last result. `fail(message)` returns the error
to the caller, the message is a string:

      func half(n int) int, error {
        if n % 2 != 0 {
          fail("odd number: " + n)
        }
        return n / 2
      }

A call of such a function passes the error on with `try`, or catches it. The compiler
rejects a call that does neither, and a `try` or `catch` on a function that can't fail.

      func quarter(n int) int, error {
        h = try half(n)         // when half fails, quarter fails with its error
        return try half(h)
      }

      q = half(7) catch err {   // err is the message
        println("no half:", err)
        return
      }

  - `try` and `fail` only work in a function that can fail, so a package can't end the
    script behind its caller's back. At the top level they print the message to stderr and
    end the script with exit code 1
  - `try` goes in front of any call, also in an expression: `println(try half(4) + 1)`
  - `catch` handles the error of the statement's call, in `f() catch err { ... }` and
    `x, y = f() catch err { ... }`. When `f` fails, the block runs instead of the
    assignment and `x` and `y` keep their values. So a variable that's new in the
    statement needs a block that ends with `return`, `fail`, `break` or `continue`
  - `err` only exists in the block, like a loop variable, and `_` ignores it
  - a global of a package can't call a function that can fail
  - errors come from `fail`, `parseInt` and `cmd`. An index out of range still ends the script
    with exit code 2

Commands
--------

`cmd(name, args..., options...)` runs a command, the statement waits until it has ended.
The options say where its streams go, the ones with a type are the results, always in the
order stdout, stderr, exit. Without options, the command uses the script's own streams
and a nonzero exit is an error. [cr-cmd.md](cr-cmd.md) has the reasons.

      tag = try cmd("git", "describe", "--tags", "--always", stdout: string)
      try cmd("docker", "build", buildArgs..., ".", dir: "./api", $DOCKER_BUILDKIT: "1")

      if try cmd("test", "-f", path, exit: bool) { }

      out, errText, code = cmd("make", stdout: string, stderr: string, exit: int)

      for line in try cmd("docker", "logs", "-f", id, stdout: pipe, stderr: stdout) {
        if strings.hasPrefix(line, "ready") {
          break                // stops docker logs
        }
      }

      errText, code = cmd("docker", "build", ".", stdout: pipe, stderr: string, exit: int) {
        for line in stdout {   // the block runs while the command runs
          println(line)
        }
      }

      sum = try cmd("sha256sum", stdin: cmd("gzip", "-dc", "dump.gz", stdout: pipe), stdout: string)

  - `stdout:` and `stderr:` take `string` (without the last newline), `[]string` (the
    lines), `pipe`, `file(path)`, `appendFile(path)` or `_` to discard it. `stderr:` also
    takes `stdout`, like `2>&1`
  - `exit:` takes `int`, `bool` (0 is true, 1 is false, the rest is an error) or `_`, a
    nonzero exit is no error with all three
  - `stdin:` takes a string, which the command reads as it is, `file(path)` or a
    `cmd(..., stdout: pipe)`, which runs together with it like `a | b`. A command that
    feeds another one can't have results, and the statement fails when it fails, like with
    `pipefail`
  - `dir:` is the working directory, `$NAME:` sets an env var for this command only
  - a `cmd(..., stdout: pipe)` as an argument is a path to read its output, like `<(...)`
  - `xs...` passes the elements of an array as arguments
  - a pipe is read while the command runs, by a `for`, by the block after the call or by
    another command. When its reader stops early, or the block is left with `return`,
    `fail`, an error or `break`, the command gets SIGTERM. That's no error, with `exit: int`
    the code is 143
  - the error message has the command's name and the exit code, not its arguments, and
    the collected stderr: `git failed with exit code 128: fatal: No names found`
  - the command runs with `exec`, so `cmd("main")` doesn't run a function of the script
    and `cmd("cd", "/")` fails, builtins aren't commands

Globals
-------

`global x = ...` at the top level of a file defines a global. It's visible in the whole
package, other packages can read it as `pkg.name`.

      import "strings"

      global host = strings.trimSuffix($DEPLOY_HOST, "/")
      global retries = 0

      func retry() {
        global retries = retries + 1
      }

  - in main, the first `global x = ...` at the top level, outside of blocks, defines the
    global. It runs where it stands, like any other statement. Using the global in a
    statement before it, also through the functions that statement calls, is a compile
    error, even in a branch that never runs
  - a package has no statements, so its globals get their values before main runs. Each
    one comes after everything it reads or writes, also through the functions it calls,
    and a cycle is a compile error. Apart from that, imported packages come first, each in
    source order
  - every write needs `global` too, also `global x[i] = ...`, and only works inside the
    global's package. A plain `x = ...` with the name of a global is a compile error, so
    adding a global can't silently change what a function does
  - locals and parameters can't have the name of a global of their package

Std
---

`strings` works like the package of Go. The arguments are text, never patterns:

      hasPrefix(s, prefix) bool           trimPrefix(s, prefix) string
      hasSuffix(s, suffix) bool           trimSuffix(s, suffix) string
      contains(s, sub) bool               indexOf(s, sub) int          // -1 when it's missing
      cut(s, sep) string, string, bool    // before, after, found
      split(s, sep) []string              join(parts, sep) string
      replaceAll(s, old, new) string      trimSpace(s) string
      toUpper(s) string                   toLower(s) string

  - `split` splits at every separator, so `split("a,", ",")` is `["a", ""]`. An empty
    separator splits into the characters
  - `trimSpace` removes spaces, tabs, newlines, `\r`, `\v` and `\f`. `toUpper` and `toLower`
    only change the ASCII letters, `é` stays as it is

`files` reads and writes text files:

      exists(path) bool                   isFile(path) bool            isDir(path) bool
      read(path) string, error            // the whole text
      lines(path) []string, error         // without the newlines, like the []string of cmd
      write(path, text) error             // creates or replaces the file
      append(path, text) error            // adds to the end, creates the file

  - an error has the path and the reason: `can't read config.json: No such file or directory`
  - `write` and `append` don't add a newline. A new file gets its permissions from the umask,
    an existing one keeps its own
  - bash strings can't hold NUL bytes, `read` and `lines` drop them

`paths` has `base(path)` and `dir(path)`, they work like the `basename` and `dirname`
commands: `base("a/b/")` is `b` and `dir("/a")` is `/`. Both give `.` for an empty path.

Tests
-----

[test cases](tests) are defined in .yaml files with code, output and possible compiler errors.

Every test runs with the local bash and in Docker with bash 3.2 and 5.3, the oldest and
the newest supported version. The Docker runs are skipped when Docker isn't available,
`go test -run TestAll/local ./tests/` runs only the local ones.

Example
-------

      func sumGt(a int, b int) int, bool {
        return a + b, a > b
      }

      s, gt = sumGt(3, 6)
      println(s, gt)

prints:

    9 false

the compiled bash code:

    #!/bin/bash
    function lsh_main_sumGt() {
      local V_a="${1}"
      local V_b="${2}"
      local L_t3=$(( V_a + V_b ))
      local L_t4="$L_t3"
      [ "$V_a" -gt "$V_b" ]
      local L_i="L_bool_$?"
      local L_t5=${!L_i}
      local L_t6="$L_t5"
      L_ret0="$L_t4"
      L_ret1="$L_t6"
      return
    }
    function main() {
      lsh_main_sumGt 3 6
      local L_t1="$L_ret0"
      local L_t2="$L_ret1"
      local V_s="$L_t1"
      local V_gt="$L_t2"
      lsh__println "$V_s" "$V_gt"
      return 0
    }

    L_bool_0=true
    L_bool_1=false
    L_args=("$@")

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

    main "$@" || { printf '%s\n' "$L_err" >&2; exit 1; }
