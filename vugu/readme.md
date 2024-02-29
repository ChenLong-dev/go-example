# Vugu使用方法

## 网址
```
开源地址：https://github.com/vugu/vugu
官网地址：https://www.vugu.org/
文档地址：https://www.vugu.org/doc
        https://www.play.vugu.org
```

## 开始
[simple](example/README.md)

### 1、安装
go get -u github.com/vugu/vgrun 或 go install github.com/vugu/vgrun@latest

vgrun -install-tools

或者

go install github.com/vugu/vugu/cmd/vugugen@latest

go install github.com/vugu/vgrouter/cmd/vgrgen@latest

```
$ go install github.com/vugu/vgrun@latest
go: downloading github.com/vugu/vgrun v0.0.0-20221010231011-b56916c1e8c2
go: downloading github.com/fsnotify/fsnotify v1.5.4
go: downloading github.com/gorilla/websocket v1.5.0

$ vgrun -install-tools
2024/02/27 21:12:34 Installing vugugen
2024/02/27 21:12:40 Installing vgrgen

```


### 2、创建一个空目录
```
$ mkdir example

$ cd example/

$ vgrun -new-from-example=simple .
2024/02/28 10:18:40 Running command: [git clone -q --depth=1 https://github.com/vugu-examples/simple .]
2024/02/28 10:18:42 Removing .git directory from example

$ ls
LICENSE  README.md  devserver.go  generate.go  go.mod  go.sum  root.vugu

```

### 3、运行开始一个开发服务
vgrun devserver.go 或 go run devserver.go

vgrun -1 devserver.go
```
$ go run devserver.go
2024/02/28 10:55:58 Starting HTTP Server at "127.0.0.1:8844"
```

### 4、用浏览器访问：http://127.0.0.1:8844/ 
```
$ go run devserver.go
2024/02/28 10:55:58 Starting HTTP Server at "127.0.0.1:8844"
WasmCompiler: Successful generate
WasmCompiler: Successful build

```
