//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 支持传参,参数可多，但是类型只能出现一次，要不然会报错。
// wire.go的文件名随意取
func GetTrip(types string, num int) *Trip {
	panic(wire.Build(NewDunHuang, NewTrain, NewTrip))
}
