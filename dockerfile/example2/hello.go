package main

/*
goctl docker -go hello.go
docker build -t hello:v1 -f Dockerfile .
docker run -it --rm hello:v1

*/

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

func InitConfig(cfgFile string, cfg interface{}) (err error) {
	if cfgFile == "" {
		err = fmt.Errorf("error: missing config file, cfgFile:%s", cfgFile)
		fmt.Println(err)
		return
	}

	viper.SetConfigFile(cfgFile) // 指定配置文件路径
	viper.AutomaticEnv()
	viper.AllowEmptyEnv(true)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err = viper.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to load config file %s: %v", cfgFile, err)
	}
	if err = viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("failed to marshal config file %s: %v", cfgFile, err)
	}
	return nil
}
