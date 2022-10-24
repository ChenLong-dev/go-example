package fileplugin

import (
	"file"
)

type T1 struct {
	A string
	B int
}

type T2 struct {
	C string
	D int
}

func NewPlugin() func() interface{} {
	return func() interface{} {
		val := T1{
			A: "T1",
			B: 1,
		}

		return T2{
			C: val.A + file.NewFile().E,
			D: val.B + file.NewFile().F,
		}
	}
}
