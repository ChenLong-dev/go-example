package main

import "fmt"

type Trip struct {
	Msg          string
	TrafficTools *Train
	Destination  *DunHuang
}

func NewTrip(trafficTools *Train, destination *DunHuang) *Trip {
	return &Trip{
		TrafficTools: trafficTools,
		Destination:  destination,
	}
}

func (t *Trip) CircleOfFriends() {
	fmt.Printf("我坐%d次%s%s去%s旅游了，好开心\n", t.TrafficTools.TrainNo, t.TrafficTools.Types, t.TrafficTools.Desc, t.Destination.Name)
}

type Train struct {
	Desc    string
	Types   string
	TrainNo int
}

func NewTrain(types string, trainNo int) *Train {
	return &Train{Desc: "火车", Types: types, TrainNo: trainNo}
}

type DunHuang struct {
	Name string
}

func NewDunHuang() *DunHuang {
	return &DunHuang{
		Name: "敦煌",
	}
}

func main() {
	GetTrip("绿皮", 1228).CircleOfFriends()
	//我坐1228次绿皮火车去敦煌旅游了，好开心
}
