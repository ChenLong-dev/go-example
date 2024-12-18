

# Wire
[Wire简介](https://zhuanlan.zhihu.com/p/449115603)

简介：
wire是 Google 开源的一个依赖注入工具。
是一个代码生成器，执行命令直接生成需要的代码，帮我们自动生成代码，组装依赖
记住一句话，注入各种依赖，又懒得用手维护的代码，它可以自动生成。就这点功能

安装wire命令：
```bash
go install github.com/google/wire/cmd/wire@latest
```

安装wire包：
```bash
go get github.com/google/wire/cmd/wire
```


使用：
```bash
$ wire -h
Usage: C:\Users\SHSZ\go\bin\wire.exe <flags> <subcommand> <subcommand args>

Subcommands:
        check            print any Wire errors found
        commands         list all command names
        diff             output a diff between existing wire_gen.go files and what gen would generate
        flags            describe all known top-level flags
        gen              generate the wire_gen.go file for each package
        help             describe subcommands and their syntax
        show             describe all top-level provider sets

```
```bash
$ wire

    wire.go 
        --------------------------
        // +build wireinject

        package main
        --------------------------
    wire_gen.go

    Provider（构造器）,Injector（注入器）

    用了wire，编译或者运行的时候:
        go run .(和原生的go run main.go的区别)
```

### 目录结构
```bash
$ tree
.
|-- Interfacevalue
|   |-- main.go
|   `-- wire.go
|-- bindinterface
|   |-- main.go
|   `-- wire.go
|-- bindvalue
|   |-- main.go
|   `-- wire.go
|-- commonparams
|   |-- main.go
|   `-- wire.go
|-- customparams
|   |-- main.go
|   `-- wire.go
|-- go.mod
|-- go.sum
|-- readme.md
|-- simple
|   |-- main.go
|    `-- wire.go
`-- structattrinjectors
    |-- main.go
    `-- wire.go


7 directories, 24 files
```
[1、简单应用](./simple)

[2,传参应用-普通传参](./commonparams)

[2,传参应用-自定义传参](./customparams)

[3,接口绑定](./bindinterface)

[4,struct属性注入](./structattrinjectors)

[5,值绑定，provider聚合-wire.Value实践](./bindvalue)

[6,值绑定，provider聚合-wire.InterfaceValue实践](./interfacevalue)
