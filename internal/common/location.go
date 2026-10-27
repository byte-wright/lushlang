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
	return l.File + " Ln:" + strconv.Itoa(l.LineFrom) + " Col:" + strconv.Itoa(l.ColFrom)
}
