//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// struct属性注入举例
func GetTrip(types string, num int) *Trip {
	panic(wire.Build(
		wire.Bind(new(Traffic), new(*Train)),
		NewDunHuang,
		NewTrain,
		wire.Struct(new(Trip), "TrafficTools", "Destination"),
	),
	)
}
