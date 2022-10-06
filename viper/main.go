package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"strings"
)

// 调用AutomaticEnv函数，开启环境变量读取
func test01() {
	fmt.Println(viper.Get("path"))
	//开始读取环境变量，如果没有调用这个函数，则下面无法读取到path的值
	viper.AutomaticEnv()
	//会从环境变量读取到该值，注意不用区分大小写
	fmt.Println(viper.Get("path"))
}

// 使用BindEnv绑定某个环境变量
func test02() {
	//将p绑定到环境变量PATH,注意这里第二个参数是环境变量，这里是区分大小写的
	viper.BindEnv("p", "PATH")
	fmt.Println("--p:", viper.Get("p")) //通过p可以读取PATH的值
	//错误绑定方式，path为小写，无法读取到PATH的值(不会出错)
	viper.BindEnv("pp", "path")
	fmt.Println("==pp:", viper.Get("pp")) //通过p可以读取PATH的值
}

// 使用函数SetEnvPrefix可以为所有环境变量设置一个前缀，这个前缀会影响AutomaticEnv和BindEnv函数
func test03() {
	os.Setenv("TEST_PATH", "test")
	viper.SetEnvPrefix("test")
	viper.AutomaticEnv()
	//无法读取path的值，因为此时加上前缀，viper会去读取TEST_PATH这个环境变量的值
	fmt.Println(viper.Get("path"))      //输出:test
	fmt.Println(viper.Get("test_path")) //输出：nil
}

// 环境变量大多是使用下划号(_)作为分隔符的，如果想替换，可以使用SetEnvKeyReplacer函数
func test04() {
	//设置一个环境变量
	os.Setenv("USER_NAME", "test")
	//将下线号替换为-和.
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_", "@", "_"))
	//读取环境变量
	viper.AutomaticEnv()
	fmt.Println(viper.Get("user.name")) //通过.访问, 输出:test
	fmt.Println(viper.Get("user-name")) //通过-访问, 输出:test
	fmt.Println(viper.Get("user_name")) //原来的下划线也可以访问, 输出:test
	fmt.Println(viper.Get("user@name")) //通过-访问, 输出:test
}

// 默认的情况下，如果读取到的环境变量值为空(注意，不是环境变量不存在，而是其值为空)，
// 会继续向优化级更低数据源去查找配置，如果想阻止这一行为，让空的环境变量值有效，则可以
// 调用AllowEmptyEnv函数
func test05() {
	viper.SetDefault("username", "admin")
	viper.SetDefault("password", "123456")
	//默认是AllowEmptyEnv(false)，这里设置为true
	viper.AllowEmptyEnv(true)
	viper.BindEnv("username")
	os.Setenv("USERNAME", "")
	fmt.Println(viper.Get("username")) //输出为空，因为环境变量USERNAME空
	fmt.Println(viper.Get("password")) //输出：123456
}

// viper可以和解析命令行库相关flag库一起工作，从命令行读取配置，其内置了对pflag库的支持，
// 同时也留有接口让我们可以支持扩展其他的flag库
func test06() {
	pflag.Int("port", 8080, "server http port")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	fmt.Println(viper.GetInt("port")) //输出8080
}

// 如果我们没有使用pflag库，但又想让viper帮我们读取命令行参数
type myFlag struct {
	f *flag.Flag
}

func (m *myFlag) HasChanged() bool {
	return false
}
func (m *myFlag) Name() string {
	return m.f.Name
}
func (m *myFlag) ValueString() string {
	return m.f.Value.String()
}
func (m *myFlag) ValueType() string {
	return "string"
}
func NewMyFlag(f *flag.Flag) *myFlag {
	return &myFlag{f: f}
}
func test07() {
	flag.String("username", "defaultValue", "usage")
	m := NewMyFlag(flag.CommandLine.Lookup("username"))
	viper.BindFlagValue("myFlagValue", m)
	flag.Parse()
	fmt.Println(viper.Get("myFlagValue"))
}

// 配置文件读取
type MySQL struct {
	Host     string
	DbName   string
	Username string
	Password string
	Charset  string
}
type Redis struct {
	Host string
	Port string
}
type Config struct {
	MySQL MySQL
	Redis Redis
}

func test08() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	var config Config
	viper.Unmarshal(&config)
	fmt.Println(config.MySQL.Username)
	fmt.Println(config.Redis.Host)
	fmt.Println(config)
}

// 写入文件
func test09() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	fmt.Println(viper.AllKeys())

	viper.WriteConfigAs("new-server.yaml")            // 直接写入，有内容就覆盖，没有文件就新建
	err := viper.SafeWriteConfigAs("new-server.yaml") // 因为该配置文件已经存在，所以会报错
	if err != nil {
		fmt.Println(err)
	}
	//viper.WriteConfig() // 将当前配置写入“viper.AddConfigPath()”和“viper.SetConfigName”设置的预定义路径
	//viper.SafeWriteConfig()
	//viper.WriteConfigAs("./config/1.yaml")
	//viper.SafeWriteConfigAs("./config/2.yaml") // 因为该配置文件写入过，所以会报错
	//viper.SafeWriteConfigAs("./config/3.yaml")
}

func test10() {
	viper.SetConfigType("yaml")
	var yamlExample = []byte(`
Hacker: true
name: steve
hobbies:
- skateboarding
- snowboarding
- go
clothing:
  jacket: leather
  trousers: denim
age: 35
eyes : brown
beard: true
`)

	viper.ReadConfig(bytes.NewBuffer(yamlExample))
	fmt.Println(viper.Get("name"))
	viper.WriteConfigAs("new1-server.yaml")
}

// map比较
func CompDrainageTaskMap(data1 map[string][]common.DrainageTaskInfo, data2 map[string][]common.DrainageTaskInfo) bool {
	keySlice := make([]string, 0)
	dataSlice1 := make([]interface{}, 0)
	dataSlice2 := make([]interface{}, 0)

	if len(data1) != len(data2) {
		return false
	}
	for key, value := range data1 {
		keySlice = append(keySlice, key)
		dataSlice1 = append(dataSlice1, value)
	}
	for _, key := range keySlice {
		if data, ok := data2[key]; ok {
			dataSlice2 = append(dataSlice2, data)
		} else {
			return false
		}
	}
	dataStr1, _ := json.Marshal(dataSlice1)
	dataStr2, _ := json.Marshal(dataSlice2)

	return string(dataStr1) == string(dataStr2)
}

func main() {
	//test01()
	//test02()
	//test03()
	//test04()
	//test05()
	//test06()
	//test07()
	//test08()
	//test09()
	test10()
}
