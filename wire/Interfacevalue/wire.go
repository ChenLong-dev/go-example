//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 值绑定wire.InterfaceValue,第一个参数是一个指向用户想要提供的接口的指针
// provider聚合举例
var TripProvider = wire.NewSet(
	wire.Bind(new(Traffic), new(*Train)),
	//第一个参数是一个指向用户想要提供的接口的指针。
	//第二个参数是实际的变量值，其类型实现
	wire.InterfaceValue(new(TripDistance), 8000),
	wire.InterfaceValue(new(TripMsg), "我要去旅行，旅行很好玩"),
	NewDunHuang,
	NewTrain,
	wire.Struct(new(Trip), "Msg", "Distance", "TrafficTools", "Destination"),
)

func GetTrip(types string, num int) *Trip {
	panic(wire.Build(
		TripProvider,
	),
	)
}
