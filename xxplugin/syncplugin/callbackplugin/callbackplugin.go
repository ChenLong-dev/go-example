package callbackplugin

import (
	"context"
	"fmt"
	"xxplugin/syncplugin"
)

const (
	PluginType = "callback"
)

type CallbackPlugin struct {
}

func (c *CallbackPlugin) New(ctx context.Context, dec func(interface{}) error) (interface{}, error) {
	//p := &Params{}
	//err := dec(p)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, err
	//}
	//fmt.Println("p:", p)
	return newCallback()
}

// 关闭相关连接句柄，当前是http请求，非长连接无需关闭
func (c *CallbackPlugin) Close(interface{}) {
}

// 初始化同步插件
func init() {
	syncplugin.Plugins[PluginType] = &CallbackPlugin{}
	fmt.Printf("type:%s register is success!\n", PluginType)
}
