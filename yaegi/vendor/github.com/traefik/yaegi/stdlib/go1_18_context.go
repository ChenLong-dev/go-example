// Code generated by 'yaegi extract context'. DO NOT EDIT.

//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package stdlib

import (
	"context"
	"reflect"
	"time"
)

func init() {
	Symbols["context/context"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Background":       reflect.ValueOf(context.Background),
		"Canceled":         reflect.ValueOf(&context.Canceled).Elem(),
		"DeadlineExceeded": reflect.ValueOf(&context.DeadlineExceeded).Elem(),
		"TODO":             reflect.ValueOf(context.TODO),
		"WithCancel":       reflect.ValueOf(context.WithCancel),
		"WithDeadline":     reflect.ValueOf(context.WithDeadline),
		"WithTimeout":      reflect.ValueOf(context.WithTimeout),
		"WithValue":        reflect.ValueOf(context.WithValue),

		// type definitions
		"CancelFunc": reflect.ValueOf((*context.CancelFunc)(nil)),
		"Context":    reflect.ValueOf((*context.Context)(nil)),

		// interface wrapper definitions
		"_Context": reflect.ValueOf((*_context_Context)(nil)),
	}
}

// _context_Context is an interface wrapper for Context type
type _context_Context struct {
	IValue    interface{}
	WDeadline func() (deadline time.Time, ok bool)
	WDone     func() <-chan struct{}
	WErr      func() error
	WValue    func(key any) any
}

func (W _context_Context) Deadline() (deadline time.Time, ok bool) {
	return W.WDeadline()
}
func (W _context_Context) Done() <-chan struct{} {
	return W.WDone()
}
func (W _context_Context) Err() error {
	return W.WErr()
}
func (W _context_Context) Value(key any) any {
	return W.WValue(key)
}
