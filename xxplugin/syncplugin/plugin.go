//@Title:		plugin.go
//@Description:	连接器客户端插件原型定义
package syncplugin

import "context"

type Plugin interface {
	/*
		创建插件句柄接口
		入参是一个函数，插件内部将配置结构体指针传入即可获得配置（此配置和云端插件QueryPolicyCfg返回的结构体一致）
		返回		interface{} 	此接口是一个结构体句柄，该句柄中所有成员函数必须符合RPC的格式，所有成员函数会被注册到RPC服务中
		返回		error 			如果返回了错误信息即视作失败不会再次调用
	*/
	New(context.Context, func(interface{}) error) (interface{}, error)
	/*
		用来关闭New返回的句柄，暂时不使用免得麻烦
	*/
	Close(interface{})
}

var Plugins = make(map[string]Plugin)
