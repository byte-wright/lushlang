package common

import "strconv"

type Location struct {
	File     string
	LineFrom int
	ColFrom  int
	LineTo   int
	ColTo    int
}

func (l *Location) String() string {
	return l.File + ":" + strconv.Itoa(l.LineFrom) + ":" + strconv.Itoa(l.ColFrom)
}
