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
  - strings are UTF-8: `len`, indexes and slices count characters (Unicode code points),
    whatever the user's locale. Bytes that aren't valid UTF-8 count as one character each.
    This needs a UTF-8 locale on the system, which nearly every system has. Without one,
    a script that uses these operations exits with code 2 before it starts
  - arrays (very slow)
  - functions with params and (multiple) return params
  - if, else if, else
  - loops with `break` and `continue`, see [Loops](#loops)
  - `return` outside of functions ends the script
  - environment variables as `$NAME`, the script's arguments as `args()`
  - names start with a lowercase letter or `_`, so they can't clash with bash's
    variables like `PATH` and `IFS`
  - packages: `import "strings"` from std, `import "./deploy"` and `import "../lib/log"` from
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
      for x in 7 { }               // 0 to 6
      for x in 3 until 7 { }       // 3 to 6

  - `break` leaves the innermost loop, `continue` goes on with its next iteration. In the
    `for init; cond; step` form, continue runs the step first, like in Go
  - `for x in` reads the array, or the numbers, once before the first iteration. Changing
    them in the body doesn't change what the loop goes over, and neither does changing `x`
    or `i`
  - `x` and `i` only exist in the loop. A variable, parameter or outer loop variable with
    the same name is a compile error, so a loop can't change one by accident

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
  - every write needs `global` too, and only works inside the global's package. A plain
    `x = ...` with the name of a global is a compile error, so adding a global can't
    silently change what a function does
  - locals and parameters can't have the name of a global of their package

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

    main "$@"
