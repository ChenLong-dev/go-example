package lion

import "fmt"

type Lion struct {
	Age int
}

func (l *Lion) Breathe() {
	fmt.Println("a lion breathe ...")
}

func (l *Lion) Walk() {
	fmt.Println("a lion walk ...")
}
