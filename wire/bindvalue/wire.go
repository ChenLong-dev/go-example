//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 值绑定wire.Value举例,这里不用声明类型，但是类型只能出现一次。如果之前出现过，那只能自定义类型了
func GetTrip(types TrainTypes, num TrainNo) *Trip {
	panic(wire.Build(
		wire.Bind(new(Traffic), new(*Train)),
		wire.Value(8000),
		wire.Value("我要去旅行，旅行很好玩"),
		NewDunHuang,
		NewTrain,
		wire.Struct(new(Trip), "Msg", "Distance", "TrafficTools", "Destination"),
	),
	)
}
