//@Title:		route.go
//@Description:		函数路由模块的接口和结构体定义
package router

import (
	"context"
	"sync"
)

//路由模块不会对外部有任何依赖，内部也没有任何并发或者网络，故不作抽象

type Router struct {
	serviceMap sync.Map
}

//rcvr上的成员函数必须符合RPC的格式，同时全部会被注册到name上(name如果重复就会导致接口的覆盖)
func (r *Router) RegisterName(name string, rcvr interface{}) error {
	_, service, err := newService(rcvr)
	if err != nil {
		return err
	}

	r.serviceMap.Store(name, service)
	return nil
}

//去注册，即使不存在也不会报错
func (r *Router) UnRegister(name string) (interface{}, bool) {
	v, ok := r.serviceMap.LoadAndDelete(name)
	if ok {
		return v.(*service).Val(), true
	} else {
		return nil, false
	}
}

/*
调用函数的接口
serviceName 服务名
methodName 	方法名
param 		参数，必须是使用codec序列化的函数
返回			string 	执行结果是通过codec序列化后的字符串，如果有需要也可以返回
			error 	错误信息
*/
func (r *Router) Call(ctx context.Context, serviceName, methodName string, dec func(interface{}) error) (interface{}, error) {
	svci, ok := r.serviceMap.Load(serviceName)
	if !ok {
		return nil, ErrNoSuchService
	}
	svc := svci.(*service)

	return svc.Call(ctx, methodName, dec)
}

// New ...
func New() *Router {
	return &Router{}
}
