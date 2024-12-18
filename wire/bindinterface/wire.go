//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// 我们应该依赖接口，而不是实现。返回数据的时候返回实现而不是接口，这是在 Golang 中的最佳实践
// 接口绑定举例
func GetTrip(types string, num int) *Trip {
	panic(wire.Build(
		wire.Bind(new(Traffic), new(*Train)),
		NewDunHuang,
		NewTrain,
		NewTrip,
	),
	)
}
