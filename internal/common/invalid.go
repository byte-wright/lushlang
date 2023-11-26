package common

import "fmt"

func IncompleteSwitch(t any) {
	panic(fmt.Sprintf("unhandled type %T", t))
}
