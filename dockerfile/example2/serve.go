package main

import (
	"example2/config"
	"fmt"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"os/signal"
	"syscall"
)

var serveCommand = &cobra.Command{
	Use:   "cron",
	Short: "Start the timer jobs",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO：init logx
		if err := initLogx(config.Cfg.Logx); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer logx.Close()

		//TODO: start demo ...
		logx.Info("timerjobs start ...")
		// 优雅关闭grpc连接
		//
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
		go func() {
			select {
			case sig := <-c:
				{
					logx.Infof("Got %s signal. Aborting...", sig)
					os.Exit(1)
				}
			}
		}()

		logx.Info(config.Cfg)

		select {}
	},

	PreRun: func(cmd *cobra.Command, args []string) {
		if err := InitConfig(config.CfgFile, &config.Cfg); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// init ...
func init() {
	serveCommand.Flags().StringVarP(&config.CfgFile, "config", "c", "", "config file(required)")
	serveCommand.MarkFlagRequired("config")

	rootCmd.AddCommand(serveCommand)
}

func initLogx(cfg config.Logx) error {
	c := logx.LogConf{
		ServiceName: cfg.ServiceName,
		Mode:        cfg.Mode,
		Encoding:    cfg.Encoding,
		TimeFormat:  "2006/01/02 15:04:05",
		Path:        cfg.Path,
		Level:       cfg.Level,
		KeepDays:    cfg.KeepDays,
		MaxBackups:  cfg.MaxBackups,
		MaxSize:     cfg.MaxSize,
	}
	logx.MustSetup(c)
	logx.Info("logx init is success ...")

	// 禁用stat日志
	logx.DisableStat()
	//logx.SetLevel(logx.ErrorLevel)
	//logx.Infow("扩展日志输出的字段，添加了 uid 字段记录请求的用户的 uid", logx.Field("uid", uuid.New().String()))

	// 扩展其他第三方日志库，通过 logx.SetWriter 来进行设置
	//writer := logrusx.NewLogrusWriter(func(logger *logrus.Logger) {
	//	logger.SetFormatter(&logrus.JSONFormatter{})
	//})
	//logx.SetWriter(writer)
	return nil
}
