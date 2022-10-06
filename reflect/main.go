package main

import (
	"fmt"
	"reflect"
	"strings"
)

type T struct {
	A int    `json:"aaa" test:"testaaa"`
	B string `json:"bbb" test:"testbbb"`
}

// 动态调用函数（无参数）
func test01() {
	name := "Do01"
	t := &T{}
	reflect.ValueOf(t).MethodByName(name).Call(nil)
	fmt.Println("reflect.ValueOf(t):", reflect.ValueOf(t))
	fmt.Println("reflect.ValueOf(t).MethodByName(name):", reflect.ValueOf(t).MethodByName(name))
}

/*
$ go run main.go
hello chenlong's test01
reflect.ValueOf(t): &{0 }
reflect.ValueOf(t).MethodByName(name): 0x9f9b40
*/

func (t *T) Do01() {
	fmt.Println("hello chenlong's test01")
}

// 动态调用函数（有参数）
func test02() {
	name := "Do02"
	t := &T{}
	a := reflect.ValueOf(12345)
	b := reflect.ValueOf(" chenlong ")
	in := []reflect.Value{a, b}
	reflect.ValueOf(t).MethodByName(name).Call(in)
}

/*
$ go run main.go
hello chenlong  12345
*/

func (t *T) Do02(a int, b string) {
	fmt.Println("hello"+b, a)
}

// 处理返回值中的错误
// 返回值也是 Value 类型，对于错误，可以转为 interface 之后断言
func test03() {
	name := "Do03"
	t := &T{}
	ret := reflect.ValueOf(t).MethodByName(name).Call(nil)
	fmt.Println("ret:", ret)
	fmt.Println("ret[0]:", ret[0])
	fmt.Println("ret[1]", ret[1])
	fmt.Printf("strValue: %[1]v\nerrValue: %[2]v\nstrType: %[1]T\nerrType: %[2]T", ret[0], ret[1].Interface().(error))
}

/*
$ go run main.go
ret: [hello alexchen <error Value>]
ret[0]: hello alexchen
ret[1] a error in here!
strValue: hello alexchen
errValue: a error in here!
strType: reflect.Value
errType: *errors.errorString
*/

func (t *T) Do03() (string, error) {
	return "hello alexchen", fmt.Errorf("a error in here!")
}

// struct tag 解析
func test04() {
	t := T{
		A: 123,
		B: "xxxix",
	}
	tt := reflect.TypeOf(t)
	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		fmt.Println(i, ":", field)
		if json, ok := field.Tag.Lookup("json"); ok {
			fmt.Println("json:", json)
		}
		test := field.Tag.Get("test")
		fmt.Println("test:", test)
	}
}

/*
$ go run main.go
0 : {A  int json:"aaa" test:"testaaa" 0 [0] false}
json: aaa
test: testaaa
1 : {B  string json:"bbb" test:"testbbb" 8 [1] false}
json: bbb
test: testbbb
*/

// 类型转换和赋值
type T2 struct {
	C int    `newT:"CC"`
	D string `newT:"DD"`
}
type newT struct {
	CC int
	DD string
}

func test05() {
	t := T2{
		C: 123,
		D: "alexchen",
	}
	tt := reflect.TypeOf(t)
	tv := reflect.ValueOf(t)
	fmt.Println("tt:", tt)
	fmt.Println("tv:", tv)

	newT := &newT{}
	newTValue := reflect.ValueOf(newT)
	fmt.Println("newTValue:", newTValue)

	for i := 0; i < tt.NumField(); i++ {
		field := tt.Field(i)
		fmt.Println("field:", field)
		newTTag := field.Tag.Get("newT")
		fmt.Println("newTTag:", newTTag)
		tValue := tv.Field(i)
		fmt.Println("tValue:", tValue)
		newTValue.Elem().FieldByName(newTTag).Set(tValue)
		fmt.Println(i, "-newT:", newT)
	}
	fmt.Println("newT:", newT)
}

/*
$ go run main.go
tt: main.T2io
tv: {123 alexchen}
newTValue: &{0 }
field: {C  int newT:"CC" 0 [0] false}
newTTag: CC
tValue: 123
0 -newT: &{123 }
field: {D  string newT:"DD" 8 [1] false}
newTTag: DD
tValue: alexchen
1 -newT: &{123 alexchen}
newT: &{123 alexchen}
*/

// 通过 kind（）处理不同分支
func test06() {
	a := 1
	t := reflect.TypeOf(a)
	fmt.Println("t:", t, t.Kind())
	switch t.Kind() {
	case reflect.Int:
		fmt.Println("int")
	case reflect.String:
		fmt.Println("string")
	}
}

/*
$ go run main.go
t: int int
int
*/

// 判断实例是否实现了某接口
type IT interface {
	test1()
	//test2()
}

type T3 struct {
	A string
}

func (t *T3) test1() {

}

func test07() {
	t := &T3{}
	ITF := reflect.TypeOf((*IT)(nil)).Elem()
	fmt.Println("ITF:", ITF)
	tv := reflect.TypeOf(t)
	fmt.Println("tv:", tv)
	fmt.Println(tv.Implements(ITF))
}

/*
$ go run main.go
ITF: main.IT
tv: *main.T3
true
*/

func Print(t interface{}) string {
	tv := reflect.ValueOf(t)
	fmt.Println("tv:", tv)
	elem := tv.Elem()
	fmt.Printf("%v\n", elem)
	elemType := elem.Type()
	fmt.Printf("%v\n", elemType)
	numField := elem.NumField()
	fmt.Printf("%v\n", numField)

	var s string
	for i := 0; i < numField; i++ {
		fieldName := elemType.Field(i).Name

		field := elem.Field(i)
		fmt.Printf("%v\n", field)

		fieldType := field.Type()
		fieldValue := field.Interface()

		fmt.Printf("%d: %s %s = %v\n", i, fieldName, fieldType, fieldValue)

		if fieldType.Kind() == reflect.Slice {
			fmt.Printf("========= %d: %s %s = %v\n", i, fieldName, fieldType, fieldValue)
			continue
		}

		if len(s) == 0 {
			s = fmt.Sprintf("%s: %v=%v", elemType, fieldName, fieldValue)
		} else {
			s = strings.Join([]string{s, fmt.Sprintf("%v=%v", fieldName, fieldValue)}, ", ")
		}

	}
	return s
}

func test09() {
	type T struct {
		A int    `json:"aaa" test:"testaaa"`
		B string `json:"bbb" test:"testbbb"`
		C []byte
	}

	t := &T{
		A: 10,
		B: "xxx",
		C: []byte("yyy"),
	}

	//tt := reflect.TypeOf(t)
	//fmt.Println("tt:", tt)
	tv := reflect.ValueOf(t)
	fmt.Println("tv:", tv)
	elem := tv.Elem()
	fmt.Printf("%v\n", elem)
	elemType := elem.Type()
	fmt.Printf("%v\n", elemType)
	numField := elem.NumField()
	fmt.Printf("%v\n", numField)

	var s string
	for i := 0; i < numField; i++ {
		fieldName := elemType.Field(i).Name

		field := elem.Field(i)
		fmt.Printf("%v\n", field)

		fieldType := field.Type()
		fieldValue := field.Interface()

		fmt.Printf("%d: %s %s = %v\n", i, fieldName, fieldType, fieldValue)

		if fieldType.Kind() == reflect.Slice {
			fmt.Printf("========= %d: %s %s = %v\n", i, fieldName, fieldType, fieldValue)
			continue
		}

		if len(s) == 0 {
			s = fmt.Sprintf("%s: %v=%v", elemType, fieldName, fieldValue)
		} else {
			s = strings.Join([]string{s, fmt.Sprintf("%v=%v", fieldName, fieldValue)}, ", ")
		}

	}

	fmt.Println(t)
	fmt.Println("+++", s)
	fmt.Println("-----", Print(t))

}

func main() {
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	//test08()
	test09()

}
