package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlConfig MysqlConfig
	Auth        Auth
	//Telemetry   Telemetry
}

type MysqlConfig struct {
	DataSource     string
	ConnectTimeout int64
}

type Auth struct {
	Secret string
	Expire int64
}

// https://github.com/Mikaelemmmm/go-zero-looklook/blob/main/doc/chinese/12-%E9%93%BE%E8%B7%AF%E8%BF%BD%E8%B8%AA.md
//type Telemetry struct {
//	Name     string  `json:",optional"`
//	Endpoint string  `json:",optional"`
//	Sampler  float64 `json:",default=1.0"`
//	Batcher  string  `json:",default=jaeger,options=jaeger|zipkin"`
//}
