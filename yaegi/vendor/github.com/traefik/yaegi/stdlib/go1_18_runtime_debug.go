// Code generated by 'yaegi extract runtime/debug'. DO NOT EDIT.

//go:build go1.18 && !go1.19
// +build go1.18,!go1.19

package stdlib

import (
	"reflect"
	"runtime/debug"
)

func init() {
	Symbols["runtime/debug/debug"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"FreeOSMemory":    reflect.ValueOf(debug.FreeOSMemory),
		"ParseBuildInfo":  reflect.ValueOf(debug.ParseBuildInfo),
		"PrintStack":      reflect.ValueOf(debug.PrintStack),
		"ReadBuildInfo":   reflect.ValueOf(debug.ReadBuildInfo),
		"ReadGCStats":     reflect.ValueOf(debug.ReadGCStats),
		"SetGCPercent":    reflect.ValueOf(debug.SetGCPercent),
		"SetMaxStack":     reflect.ValueOf(debug.SetMaxStack),
		"SetMaxThreads":   reflect.ValueOf(debug.SetMaxThreads),
		"SetPanicOnFault": reflect.ValueOf(debug.SetPanicOnFault),
		"SetTraceback":    reflect.ValueOf(debug.SetTraceback),
		"Stack":           reflect.ValueOf(debug.Stack),
		"WriteHeapDump":   reflect.ValueOf(debug.WriteHeapDump),

		// type definitions
		"BuildInfo":    reflect.ValueOf((*debug.BuildInfo)(nil)),
		"BuildSetting": reflect.ValueOf((*debug.BuildSetting)(nil)),
		"GCStats":      reflect.ValueOf((*debug.GCStats)(nil)),
		"Module":       reflect.ValueOf((*debug.Module)(nil)),
	}
}
