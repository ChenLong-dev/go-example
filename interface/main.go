package main

import (
	"fmt"
	"interface/dog"
	"interface/ianimal"
	"interface/lion"
)

func Call(a ianimal.Animal) {
	a.Walk()
}

func Print(a ianimal.Animal) {
	l, ok := a.(*dog.Dog)
	if ok {
		fmt.Println(l)
	} else {
		fmt.Println("a is not of type lion")
	}
}

func main() {
	var a ianimal.Animal
	a = &lion.Lion{Age: 10}
	a.Breathe()
	Call(a)

	b := &dog.Dog{Age: 20}
	Print(b)
	b.Breathe()
}
