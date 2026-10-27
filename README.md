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
  - modules

Tests
-----

[test cases](tests) are defined in .yaml files with code, output and possible compiler errors.

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
      local L_t3=$(( "$V_a" + "$V_b" ))
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
