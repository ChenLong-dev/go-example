package main

import "fmt"

type Person struct {
	Name   string
	Age    int
	Gender int
	Height int

	Country string
	City    string
}

// NewPerson
/*
传统方式
*/
func NewPerson(name string, age, gender, height int, country, city string) *Person {
	return &Person{
		Name:    name,
		Age:     age,
		Gender:  gender,
		Height:  height,
		Country: country,
		City:    city,
	}
}

func test01() {
	person := NewPerson("dongxiaojian", 18, 1, 180, "china", "Beijing")
	fmt.Printf("%+v\n", person)
}

//
/*
options模式
*/
type Options func(*Person)

func WithPersonProperty(name string, age, gender, height int) Options {
	return func(p *Person) {
		p.Name = name
		p.Age = age
		p.Gender = gender
		p.Height = height
	}
}

func WithRegional(country, city string) Options {
	return func(p *Person) {
		p.Country = country
		p.City = city
	}
}

func NewPersonOptions(opt ...Options) *Person {
	p := new(Person)
	p.Country = "china"
	p.City = "beijing"
	fmt.Printf("--- opt:%v\n", opt)
	for i, o := range opt {
		fmt.Printf("---[%d] o:%v\n", i, o)
		o(p)
		fmt.Printf("---[%d] p:%v\n", i, p)
	}
	return p
}

func test02() {
	// 默认值方式
	person := NewPersonOptions(WithPersonProperty("dongxiaojian", 18, 1, 180))
	fmt.Printf("%+v\n", person)

	// 设置值
	person2 := NewPersonOptions(WithPersonProperty("dongxiaojian", 18, 1, 180), WithRegional("china", "shenzhen"))
	fmt.Printf("%+v\n", person2)
}

func main() {
	test01()
	fmt.Println("====================================")
	test02()
}
