package example

type ValStringer struct {
	val string
}

func (v ValStringer) String() string {
	return v.val
}
