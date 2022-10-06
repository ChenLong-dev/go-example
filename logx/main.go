package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/zeromicro/go-zero/core/logx"
	"os"
	"test.logx/example"
)

func main() {
	c := logx.LogConf{
		ServiceName: "test.logx",
		Mode:        "console",
		Encoding:    "plain",
		TimeFormat:  "2006/01/02 15:04:05",
		Path:        "logs",
		Level:       "info",
		KeepDays:    7,
		MaxBackups:  3,
		MaxSize:     500,
	}
	logx.MustSetup(c)

	fmt.Fprintln(os.Stderr, color.HiRedString("[x] Remote target required"))

	example.FieldExample()
}
