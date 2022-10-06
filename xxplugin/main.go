package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.uber.org/atomic"
	router "xxplugin/route"
	"xxplugin/syncplugin"
	"xxplugin/syncplugin/callbackplugin"
)

type server struct {
	*router.Router              //函数路由器
	close          func()       //用来关闭句柄的函数
	last           atomic.Int64 //最后访问时间
}

func tfunc(p interface{}) interface{} {
	return p
}

func test01() {
	pi := tfunc(&callbackplugin.Params{
		Name: "Sangfor",
	})
	arg, err := json.Marshal(pi)
	fmt.Println("arg:", arg, err)

	var v callbackplugin.Params
	err = json.Unmarshal(arg, &v)
	fmt.Println("v:", v, err)

}

func test() {
	p, ok := syncplugin.Plugins[callbackplugin.PluginType]
	if !ok {
		fmt.Printf("ok:%v\n", ok)
		return
	}

	pi := tfunc(&callbackplugin.Params{
		Name: "Sangfor",
	})
	//arg, err := json.Marshal(&callbackplugin.Params{
	//	Name: "Sangfor",
	//})
	arg, err := json.Marshal(pi)
	dec := func(v interface{}) error {
		return json.Unmarshal(arg, v)
	}
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	ctx := context.TODO()
	//handler, err := p.New(ctx, func(i interface{}) error {
	//	return json.Unmarshal(arg, i)
	//})
	handler, err := p.New(ctx, nil)
	srv := &server{
		Router: router.New(),
		close: func() {
			p.Close(handler)
		},
	}
	err = srv.RegisterName("", handler)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	res, err := srv.Call(ctx, "", "CallBack", dec)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}
	resp, ok := res.(*callbackplugin.Result)
	if !ok {
		fmt.Printf("ok:%v\n", ok)
		return
	}
	fmt.Printf("resp:%v\n", resp)
}

func main() {
	fmt.Println("xxxxxxxxxxxyyy")
	//test()
	test01()
}
