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
      local a="${1}"
      local b="${2}"
      local var_3=$(( "$a" + "$b" ))
      local var_4="$var_3"
      [ "$a" -gt "$b" ]
      local lsh_i="lsh_bool_$?"
      local var_5=${!lsh_i}
      local var_6="$var_5"
      lsh__funcretparam_0="$var_4"
      lsh__funcretparam_1="$var_6"
      return
    }
    function main() {
      lsh_main_sumGt 3 6
      local var_1="$lsh__funcretparam_0"
      local var_2="$lsh__funcretparam_1"
      local s="$var_1"
      local gt="$var_2"
      lsh__println "$s" "$gt"
    }

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

    main "$@"
