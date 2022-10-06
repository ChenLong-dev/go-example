package dog

import "fmt"

type Dog struct {
	Age int
}

func (d *Dog) Breathe() {
	fmt.Println("a dog breathe ...")
}

func (d *Dog) Walk() {
	fmt.Println(" a dog walk ...")
}



