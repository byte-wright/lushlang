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
	register(&bFunc{
		name: "substring",
		code: `local i2=$(($3-$2))
lsh__funcretparam_0="${1:$2:$i2}"`,
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
