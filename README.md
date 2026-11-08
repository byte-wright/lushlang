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

  - math and comparison operators
  - int, string, bools and type inference
  - arrays (very slow)
  - functions with params and (multiple) return params
  - if, else if, else and for
  - `return` outside of functions ends the script
  - environment variables as `$NAME`, the script's arguments as `args()`
  - names start with a lowercase letter or `_`, so they can't clash with bash's
    variables like `PATH` and `IFS`
  - modules
  - global variables per package
  - the scripts need bash 3.2 or newer, so they also run with macOS's `/bin/bash`

Globals
-------

`var` declares a global at the top level of a file. It's visible in the whole package,
other packages can read it as `pkg.name`.

      import "strings"

      var host = strings.trimSuffix($DEPLOY_HOST, "/")
      var retries = 0

      func retry() {
        global retries = retries + 1
      }

  - every `var` gets its value once, before the top level of the script runs and after
    everything it reads or writes, also through the functions it calls. A cycle is a
    compile error, even through a branch that never runs
  - apart from that, imported packages come first, then main, each in source order
  - only `global x = ...` writes a global, and only inside its package. A plain
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
