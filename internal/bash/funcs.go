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
		name: "trimPrefix",
		code: `lsh__funcretparam_0="${1#"$2"}"`,
	})

	register(&bFunc{
		name: "hasPrefix",
		code: `if [[ "$1" == "$2"* ]]; then
	lsh__funcretparam_0=true
else
	lsh__funcretparam_0=false
fi`,
	})

	register(&bFunc{
		name: "trimSuffix",
		code: `lsh__funcretparam_0="${1%"$2"}"`,
	})

	register(&bFunc{
		name: "hasSuffix",
		code: `if [[ "$1" == *"$2" ]]; then
	lsh__funcretparam_0=true
else
	lsh__funcretparam_0=false
fi`,
	})

	register(&bFunc{
		name: "indexOf",
		code: `local pos=$(awk -v a="$1" -v b="$2" 'BEGIN{print index(a,b)}')
if [ "$pos" -eq 0 ]; then
	lsh__funcretparam_0=-1
else
	lsh__funcretparam_0=$((pos - 1))
fi`,
	})

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

	register(&bFunc{
		name: "join",
		code: `local i=2
local v=""
local di=$(($1+2))
for (( i=2;i<$di;i++ )); do
	if [ $i -eq 2 ]; then
		v="${!i}"
	else
		v="$v${!di}${!i}"
	fi
done
lsh__funcretparam_0="${v}"`,
	})
}
