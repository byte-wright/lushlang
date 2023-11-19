package common

import "bytes"

func Map[S comparable, T comparable](s []S, f func(S) T) []T {
	t := []T{}

	for _, e := range s {
		t = append(t, f(e))
	}

	return t
}

func JoinComma[S comparable](s []S, f func(S) string) string {
	buf := bytes.NewBuffer(nil)

	for i := 0; i < len(s); i++ {
		if i > 0 {
			buf.WriteString(", ")
		}

		str := f(s[i])

		buf.WriteString(str)
	}

	return buf.String()
}
