package main

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
)

type Dog struct {
	A string
	B int
}
type Snake struct {
	A string
	B int
	C float64
}
type APIContent interface {
}

func testReflect() {
	log.Println(" dog content---------------------------")
	dog := &Dog{}
	dogObj := reflect.TypeOf(dog)
	log.Println("dog type:", dogObj)
	log.Println("dog kind:", dogObj.Kind())
	log.Println("dog type kind:", dogObj.Kind() == reflect.Ptr)
	dogObj = dogObj.Elem()
	log.Println("model elem dog:", dogObj)
	// apicontent
	log.Println(" snake content---------------------------")
	snake := &Snake{}
	snakeObj := reflect.TypeOf(snake)
	log.Println("snake type:", snakeObj)
	log.Println("snake kind:", snakeObj.Kind())
	log.Println("snake type kind", snakeObj.Kind() == reflect.Ptr)
	snakeObj = snakeObj.Elem()
	log.Println("model elem snake :", snakeObj)

	log.Println("copy---------------------------")

	jsonText := `{
		"A":"user1",
		"B":45,
		"C":2.3
	}`
	contentValue1 := reflect.New(snakeObj)
	contentSrc := contentValue1.Interface().(APIContent)
	err := json.Unmarshal([]byte(jsonText), &contentSrc)
	log.Println("Err:", err)
	contentValue2 := reflect.New(dogObj)
	contentDst := contentValue2.Interface().(APIContent)
	log.Println("api contentSrc:", contentSrc)
	log.Println("api contentDst:", contentDst)

	log.Println("copyField err:", copyField(contentDst, contentSrc, []string{}))

	log.Println("contentSrc:", contentSrc)
	log.Println("api contentDst:", contentDst)
}

type R struct {
	A string
	B string
	C int
	D int
}

type R2 struct {
	B string
	C int
}

func (r *R) fu() {
	fmt.Println("xxxxx")
}

func testReflect2() {

	a := &R{
		A: "chenlong",
		B: "alexchen",
		C: 10,
		D: 20,
	}
	b := &R2{}
	log.Println("copyField err:", copyField(b, a, []string{}))
	log.Println("a:", a)
	log.Println("b:", b)
}

func test08() {
	//testReflect()
	testReflect2()
}

/*
$ go run main.go interface-model.go
2022/09/18 21:52:21  dog content---------------------------
2022/09/18 21:52:21 dog type: *main.Dog
2022/09/18 21:52:21 dog kind: ptr
2022/09/18 21:52:21 dog type kind: true
2022/09/18 21:52:21 model elem dog: main.Dog
2022/09/18 21:52:21  snake content---------------------------
2022/09/18 21:52:21 snake type: *main.Snake
2022/09/18 21:52:21 snake kind: ptr
2022/09/18 21:52:21 snake type kind true
2022/09/18 21:52:21 model elem snake : main.Snake
2022/09/18 21:52:21 copy---------------------------
2022/09/18 21:52:21 Err: <nil>
2022/09/18 21:52:21 api contentSrc: &{user1 45 2.3}
2022/09/18 21:52:21 api contentDst: &{ 0}
2022/09/18 21:52:21 copyField err: <nil>
2022/09/18 21:52:21 contentSrc: &{user1 45 2.3}
2022/09/18 21:52:21 api contentDst: &{user1 0}
*/

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
func copyField(to interface{}, from interface{}, excepts []string) error {
	toVal := reflect.ValueOf(to)
	if toVal.Kind() == reflect.Ptr {
		toVal = toVal.Elem()
	}
	fromVal := reflect.ValueOf(from)
	if fromVal.Kind() == reflect.Ptr {
		fromVal = fromVal.Elem()
	}
	// to fileld
	toType := toVal.Type()
	fieldNum := toType.NumField()
	for i := 0; i < fieldNum; i++ {
		toField := toType.Field(i)
		if excepts != nil && stringInSlice(toField.Name, excepts) {
			continue
		}
		toValField := toVal.Field(i)
		if !toValField.CanSet() {
			continue
		}
		if fromValField := fromVal.FieldByName(toField.Name); fromValField.IsValid() && fromValField.Type() == toValField.Type() {
			toValField.Set(fromValField)
			continue
		}
		if fromFunc := fromVal.Addr().MethodByName(toField.Name); fromFunc.IsValid() &&
			fromFunc.Type().NumOut() >= 1 &&
			fromFunc.Type().Out(0) == toValField.Type() &&
			fromFunc.Type().NumIn() == 0 {
			res := fromFunc.Call(make([]reflect.Value, 0))
			if len(res) > 1 {
				last := res[len(res)-1]
				if last.CanInterface() && !last.IsNil() {
					if err, ok := last.Interface().(error); ok {
						return err
					}
				}

			}
			toValField.Set(res[0])
			continue
		}
	}
	// to func

	toVal = toVal.Addr()
	toType = toVal.Type()
	funcNum := toType.NumMethod()
	for i := 0; i < funcNum; i++ {
		// method from type
		toMethod := toType.Method(i)
		if !strings.HasPrefix(toMethod.Name, "Set") {
			// only SetXXX methods
			continue
		}

		name := strings.TrimPrefix(toMethod.Name, "Set")
		// skip excepts
		if excepts != nil && stringInSlice(name, excepts) {
			continue
		}

		// func from value
		toFunc := toVal.MethodByName(toMethod.Name)
		argType := toFunc.Type().In(0)

		// from field
		if fromValField := fromVal.FieldByName(name); fromValField.IsValid() && fromValField.Type() == argType {
			res := toFunc.Call([]reflect.Value{fromValField})
			if len(res) > 0 {
				last := res[len(res)-1]
				if last.CanInterface() && !last.IsNil() {
					if err, ok := last.Interface().(error); ok {
						return err
					}
				}

			}
			continue
		}
		// from func

		if fromFunc := fromVal.Addr().MethodByName(name); fromFunc.IsValid() &&
			fromFunc.Type().NumOut() >= 1 &&
			fromFunc.Type().Out(0) == argType &&
			fromFunc.Type().NumIn() == 0 {
			res := fromFunc.Call(make([]reflect.Value, 0))
			if len(res) > 1 {
				last := res[len(res)-1]

				if last.CanInterface() && !last.IsNil() {
					if err, ok := last.Interface().(error); ok {
						return err
					}
				}

			}

			res = toFunc.Call([]reflect.Value{res[0]})
			if len(res) > 0 {
				last := res[len(res)-1]
				if last.CanInterface() && !last.IsNil() {
					if err, ok := last.Interface().(error); ok {
						return err
					}
				}

			}
			continue
		}

	}
	return nil
}
