// Code generated by 'yaegi extract errors'. DO NOT EDIT.

//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package stdlib

import (
	"errors"
	"reflect"
)

func init() {
	Symbols["errors/errors"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"As":     reflect.ValueOf(errors.As),
		"Is":     reflect.ValueOf(errors.Is),
		"New":    reflect.ValueOf(errors.New),
		"Unwrap": reflect.ValueOf(errors.Unwrap),
	}
}
