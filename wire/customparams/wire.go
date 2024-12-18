//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 自定义类型举例
func GetTrip(types TrainTypes, num TrainNo) *Trip {
	panic(wire.Build(
		NewDunHuang,
		NewTrain,
		NewTrip),
	)
}
