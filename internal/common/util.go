package common

func Map[S comparable, T comparable](s []S, f func(S) T) []T {
	t := []T{}

	for _, e := range s {
		t = append(t, f(e))
	}

	return t
}
