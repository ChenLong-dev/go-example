package main

import (
	"flag"
	"fmt"
	"testserver/internal/config"
	"testserver/internal/handler"
	"testserver/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 统一错误处理
	//httpx.SetErrorHandler(func(err error) (int, any) {
	//	switch e := err.(type) {
	//	case *errorx.Error:
	//		return http.StatusOK, errorx.Fail(e)
	//	default:
	//		return http.StatusInternalServerError, nil
	//	}
	//})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
