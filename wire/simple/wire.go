//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func GetTrip() *Trip {
	panic(wire.Build(NewDunHuang, NewAirplane, NewTrip))
}
