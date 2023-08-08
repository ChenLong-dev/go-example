package main

import (
	"flag"
	"fmt"
	"strings"
)

func FlagDemo() {
	//命令行是：test.exe -u root -p root123 -h localhost -port 8080
	var user, pwd, host, port string
	flag.StringVar(&user, "u", "user", "用户名")
	flag.StringVar(&pwd, "p", "123456", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "IP地址")
	flag.StringVar(&port, "port", "80", "端口号")

	flag.Parse() //解析注册的flag，必须

	fmt.Println("flag解析命令行参数如下：")
	fmt.Printf("user = %v \n", user)
	fmt.Printf("pwd = %v \n", pwd)
	fmt.Printf("host = %v \n", host)
	fmt.Printf("port = %v \n", port)
}

func FlagDemo2() {
	//命令行是：test.exe -u root -p root123 -h localhost -port 8080
	paramstr := "-u root -p root123  -h localhost -port 80 "
	paramlist := strings.Split(paramstr, " ")
	fmt.Println(paramlist)

	var user, pwd, host, port string
	var cmdl = flag.NewFlagSet("", flag.ExitOnError)
	cmdl.StringVar(&user, "u", "user", "用户名")
	cmdl.StringVar(&pwd, "p", "123456", "密码")
	cmdl.StringVar(&host, "h", "127.0.0.1", "IP地址")
	cmdl.StringVar(&port, "port", "80", "端口号")

	cmdl.Parse(paramlist) //解析注册的flag，必须

	fmt.Println("flag解析命令行参数如下：")
	fmt.Printf("user = %v \n", user)
	fmt.Printf("pwd = %v \n", pwd)
	fmt.Printf("host = %v \n", host)
	fmt.Printf("port = %v \n", port)
}

func main() {
	fmt.Println("xxxxxxx")
	FlagDemo2()
}
