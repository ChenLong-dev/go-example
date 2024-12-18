package main

import "fmt"

//举例说明:
//  这个例子完成一个旅游，旅游需要依赖一个交通工具和旅游目的地，旅行完再发一个朋友圈

// 旅行结构体
type Trip struct {
	Msg          string
	TrafficTools *Airplane //交通工具，举例飞机
	Destination  *DunHuang //目的地，举例敦煌
}

// 旅游的provider
func NewTrip(trafficTools *Airplane, destination *DunHuang) *Trip {
	return &Trip{
		TrafficTools: trafficTools,
		Destination:  destination,
	}
}

// 旅行，发朋友圈的操作
func (t *Trip) CircleOfFriends() {
	fmt.Printf("我坐%s去%s旅游了，好开心\n", t.TrafficTools.Name, t.Destination.Name)
}

// 交通工具相关
type Airplane struct {
	Name string
}

func NewAirplane() *Airplane {
	return &Airplane{
		Name: "波音飞机",
	}
}

// 目的地相关
type DunHuang struct {
	Name string
}

func NewDunHuang() *DunHuang {
	return &DunHuang{
		Name: "敦煌",
	}
}

// 手动依赖注入
func HandleDI() {
	dh := NewDunHuang()
	air := NewAirplane()
	trip := NewTrip(air, dh)
	trip.CircleOfFriends()
}

func main() {
	//调用手动依赖注入
	HandleDI()
	//调用wire生成的依赖注入
	GetTrip().CircleOfFriends()
}
