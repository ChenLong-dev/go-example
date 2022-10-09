//@Title:		config.go
//@Description: 连接器客户端的配置定义
package config

var (
	CfgFile string
	Cfg     UserSyncConfig
)

type ServerParam struct {
	Host string `mapstructure:"Host"`
	Port int    `mapstructure:"Port"`
}

type ServerConfig struct {
	GrpcServer ServerParam `mapstructure:"GrpcServer"`
	HttpServer ServerParam `mapstructure:"HttpServer"`
}

type Logx struct {
	ServiceName string `mapstructure:"ServiceName"`
	Mode        string `mapstructure:"Mode"`
	Encoding    string `mapstructure:"Encoding"`
	Path        string `mapstructure:"Path"`
	Level       string `mapstructure:"Level"`
	KeepDays    int    `mapstructure:"KeepDays"`
	MaxBackups  int    `mapstructure:"MaxBackups"`
	MaxSize     int    `mapstructure:"MaxSize"`
}

type SyncServer struct {
	LdapSvr ServerParam `mapstructure:"LdapSvr"`
}

type UserSyncConfig struct {
	Server ServerConfig `mapstructure:"Server"`
	Logx   Logx         `mapstructure:"Logx"`
}
