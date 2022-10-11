# 如何构建 Go 应用的 Docker 镜像
在部署 Go 应用时，我们通常会使用 Docker 镜像来部署，那么如何构建一个 Go 应用的 Docker 镜像呢？镜像构建过程中有没有什么最佳实践呢？

这正是本文想要讲解的内容。总的来说，本文会包含 Dockerfile 编写、镜像构建、多阶段构建、交叉编译以及使用 Makefile 简化构建流程等知识点。

## 创建一个简单的 Go 应用
为了说明整个镜像构建流程，让我们先从一个简单的 Go REST 应用开始。

### 该应用主要有以下功能：

- 访问 /，返回  Hello, Docker! <3；

- 访问 /ping ，返回 JSON 字符串 {"Status":"OK"}；

- 可以通过环境变量设置 HTTP_PORT，默认值为8080。

应用源码地址在 https://github.com/nodejh/docker-go-server-ping ，你可以直接下载使用，也可以按照下面的步骤从零开始编写代码。

### 方式一：下载源码
```
$ git clone git@github.com:nodejh/docker-go-server-ping.git
```
安装依赖模块：
```
$ go mod download
```

### 方式二：从零编写 Go 应用
#### 新建一个 docker-go-server-ping 目录，然后初始化 Go 模块：
```
$ mkdir docker-go-server-ping && cd docker-go-server-ping
$ go mod init github.com/nodejh/docker-go-server-ping
```

#### 安装 Echo 模块：
```
$ go get github.com/labstack/echo/v4
```

#### 接下来创建一个 main.go 文件，并实现一个简单的 Go 服务：
```
package main

import (
    "net/http"
    "os"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

func main() {

    e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.GET("/", func(c echo.Context) error {
        return c.HTML(http.StatusOK, "Hello, Docker! <3")
    })

    e.GET("/ping", func(c echo.Context) error {
        return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
    })

    httpPort := os.Getenv("HTTP_PORT")
    if httpPort == "" {
        httpPort = "8080"
    }

    e.Logger.Fatal(e.Start(":" + httpPort))
}
```
#### 接下来可能需要执行 go mod tidy 来确保 go.mod 和源码中的模块一致：
```
$ go mod tidy
```

#### 测试 Go 应用
让我们启动我们的 Go 应用并确保它正常运行。进入项目目录并通过 go run 命令执行源码：
```
$ go run main.go 

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.6.1
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```
让我们对应用进行一个简单的测试，打开一个新的终端，使用 curl 命令或在浏览器打开 http://localhost:8080/ 进行测试。以 curl 为例：

```
$ curl http://localhost:8080/
Hello, Docker! <3
```
可以看到应用正常返回了，正如开头描述的那样。

确定服务器正在运行并且可以访问后，我们就可以继续针对应用构建 Docker 镜像了。

## 为 Go 应用创建一个 Dockerfile

Dockerfile 是 Docker 镜像的描述文件，是一个文本文件。当我们执行 docker build 构建镜像时，Docker 就会读取 Dockerfile 中的指令来创建 Docker 镜像。

### 从零创建 Dockerfile
让我们先来看一下创建 Dockerfile 的详细过程。

在项目根目录中创建一个名为 Dockerfile 的文件并在编辑器中打开。

添加到 Dockerfile 的第一行是 # syntax 解析器指令，这是一个可选项，表示 Docker 构建器在解析 Dockerfile 时使用什么语法。解析器指令必须在 Dockerfile 其他任何注释、空格或指令之前，并且应该是 Dockerfile 的第一行。建议使用 docker/dockerfile:1 ，它始终指向版本 1 语法的最新版本。
```
# syntax=docker/dockerfile:1
```

接下来在 Dockerfile 中再添加一行，告诉 Docker 我们的应用使用什么基础镜像：
```
FROM golang:1.18.7-alpine3.16
```


这里我们使用了 Golang 官方镜像 中的 1.18.7-alpine3.16 版本作为基础镜像，alpine 是专门为容器设计的小型 Linux 发行版。使用基础镜像的好处是，基础镜像中内置了 Go 运行环境和工具，我们就不用自己再去安装了。

为了更好地在镜像中管理我们的应用程序，让我们在镜像中创建一个工作目录，之后源码或编译产物都存放在该目录中：
```
WORKDIR /app
```
接下来我们就需要在镜像中编译 Go 应用，这样做是为了保证编译和最终运行的环境一致。

通常我们编译 Go 应用的第一步是安装依赖，所以要先把 go.mod 和 go.sum 复制到镜像中：
```
COPY go.mod ./
COPY go.sum ./
```
COPY 命令可以把文件复制到镜像中，这里的 ./ 对应的目录就是上一个命令 WORKDIR 指定的 /app 目录。

然后通过 RUN 命令在镜像中执行 go mod download 安装依赖，这与我们在本机直接运行命令的作用完全相同，不同的是这次会将依赖安装在镜像中：
```
RUN go mod download
```
此时我们已经有了一个基于 golang:1.18.7 的镜像，并安装了 Go 应用所需的依赖。下一步要做的就是把源码复制到镜像中：
```
COPY *.go ./
```
该 COPY 命令使用通配符将本机当前目录（即 Dockerfile 所在目录）中后缀为 .go 的文件全部复制到镜像中。

接下来我们就可以编译 Go 应用了，依然使用熟悉的 RUN 命令：
```
RUN go build -o /docker-go-server-ping
```
该命令的结果就是产生一个名为 docker-go-server-ping 的二进制文件，并存放在镜像的系统根目录中。当然你也可以将二进制文件放在其他任何位置，根目录在这里没有特殊意义，只是它的路径较短且保持了可读性，使用起来更方便。

现在，剩下要做的就是告诉 Docker 当我们用该镜像来启动容器时要执行什么命令，这时可以使用 CMD 命令：
```
CMD [ "/docker-go-server-ping" ]
```
完整的 Dockerfile
```
FROM golang:1.18.7-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-go-server-ping

EXPOSE 8080

CMD [ "/docker-go-server-ping" ]
```
### 构建镜像
Dockerfile 编写完成后就可以使用 docker build 命令来构建镜像了。

#### 构建镜像
让我们进入 Dockerfile 所在目录构建镜像，并通过可选的 --tag 给镜像定义一个方便阅读和识别的名字和标签，格式为 <镜像名称>:<标签>，默认是标签 latest：
```
$ docker build --tag docker-go-server-ping:latest .
[+] Building 28.0s (12/12) FINISHED
 => [internal] load build definition from Dockerfile                                                       0.0s 
 => => transferring dockerfile: 248B                                                                       0.0s 
 => [internal] load .dockerignore                                                                          0.0s 
 => => transferring context: 2B                                                                            0.0s 
 => [internal] load metadata for docker.io/library/golang:1.18.7-alpine3.16                                0.0s 
 => [1/7] FROM docker.io/library/golang:1.18.7-alpine3.16                                                  0.0s 
 => [internal] load build context                                                                          0.0s 
 => => transferring context: 81B                                                                           0.0s 
 => CACHED [2/7] WORKDIR /app                                                                              0.0s 
 => CACHED [3/7] COPY go.mod ./                                                                            0.0s 
 => CACHED [4/7] COPY go.sum ./                                                                            0.0s 
 => [5/7] RUN go mod download                                                                             19.8s 
 => [6/7] COPY *.go ./                                                                                     0.0s 
 => [7/7] RUN go build -o /docker-go-server-ping                                                           6.6s 
 => exporting to image                                                                                     1.4s 
 => => exporting layers                                                                                    1.3s 
 => => writing image sha256:6df657b11bba27fe9b272d03e1dcae6744c1667d146d7ac556ba0757cc107677               0.0s 
 => => naming to docker.io/library/docker-go-server-ping:latest                                            0.0s
```
构建完成后可以通过 docker image ls （或 docker images 简写）来查看镜像列表：
```
$ docker image ls|grep ping
docker-go-server-ping             latest              6df657b11bba   About an hour ago   422MB
```
#### 为镜像设置标签
我们也可以通过 docker image tag 命令来为镜像设置新的标签，例如：
```
docker image tag docker-go-server-ping:latest docker-go-server-ping:v2
```
这时通过 docker image ls 就可以看到 docker-go-server-ping 镜像的两个标签：
```
$ docker image ls|grep ping
docker-go-server-ping             latest              6df657b11bba   2 hours ago     422MB
docker-go-server-ping             v2                  6df657b11bba   2 hours ago     422MB
```
我们还可以通过 docker image rm （简写为 docker rmi）删除镜像：
```
$ docker rmi -f docker-go-server-ping:v2
Untagged: docker-go-server-ping:v2
```
这时再查看镜像列表，v2 版本的镜像已被删除，只剩下 latest 版本了：
```
$ docker image ls|grep ping
docker-go-server-ping             latest              6df657b11bba   2 hours ago     422MB
```

### 单元测试
既然本文主要将 Go 的 Docker 镜像，这样就顺便简单说明如何使用 dockertest 对 Go 应用进行单元测试。

dockertest 可以在 Docker 容器中启动 Go 应用镜像并执行测试用例。

相关测试用例可以参考 main_test.go 。

在 main_test.go 中我们使用了 docker-go-server-ping:latest 镜像来运行 Go 应用：
```
$ docker run -p 8080:8080 docker-go-server-ping:latest

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.9.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```
所以在执行 go test 之前，需要先构建 docker-go-server-ping:latest 镜像：
```
go test
```

## 多阶段构建
可能你已经注意到了，docker-go-server-ping 镜像的大小有 407MB，这着实有点大，并且镜像中还有全套的 Go 工具、Go 应用的依赖等，但实际我们应用运行时不需要这些文件，只需要编译后的二进制文件。那么能不能减小镜像的体积呢？

要减小镜像体积，我们可以使用多阶段构建。Docker 在 17.05 版本以后，新增了多阶段构建功能。多阶段构建实际上是允许一个 Dockerfile 中出现多个 FROM 指令。通常我们使用多阶段构建时，会先使用一个（或多个）镜像来构建一些中间制品，然后再将中间制品放入另一个最新且更小的镜像中，这个最新的镜像就只包含最终需要的构建制品。

多阶段构建 Dockerfile
我们先创建一个多阶段构建的 Dockerfile，名为 Dockerfile.multistage，文件名中的 multistage 没有特殊含义，只是为了和之前的 Dockerfile 作区分，下面是完整的 Dockerfile：
```
##
## Build
##
FROM golang:1.18.7-alpine3.16 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-go-server-ping

##
## Deploy
##
FROM scratch

WORKDIR /

COPY --from=build /docker-go-server-ping /docker-go-server-ping

EXPOSE 8080

ENTRYPOINT ["/docker-go-server-ping"]
```

在 Dockerfile.multistage 中使用了两次 FROM 指令，分别对应两个构建阶段。第一个阶段构建的 FROM 指令依然使用 golang:1.18.7-alpine 作为基础镜像，并将该阶段命名为 build。第二个构建阶段的 FROM 指令使用 scratch 作为基础镜像，告诉 Docker 接下来从一个全新的基础镜像开始构建，scratch 镜像是 Docker 项目预定义的最小的镜像。第二阶段构建主要是将上个阶段中编译好的二进制文件复制到新的镜像中。

在 Go 应用中，多阶段构建非常常见，可以减小镜像的体积、节省大量的存储空间。

在 Dockerfile.multistage 中需要额外关注的是 RUN 指令，这里使用到了交叉编译。

### 交叉编译
交叉编译是指在一个平台上生成另一个平台的可执行程序。

在其他编程语言中进行交叉编译可能要借助第三方工具，但 Go 内置了交叉编译工具，使用起来非常方便，通常设置 CGO_ENABLED、GOOS 和 GOARCH 这几个环境变量就够了。

#### CGO_ENABLED
默认值是 1，即默认开启 cgo，允许在 Go 代码中调用 C 代码。

- 当 CGO_ENABLED=1 进行编译时，会将文件中引用 libc 的库（比如常用的 net 包）以动态链接的方式生成目标文件；

- 当 CGO_ENABLED=0 进行编译时，则会把在目标文件中未定义的符号（如外部函数）一起链接到可执行文件中。

所以交叉编译时，我们需要将 CGO_ENABLED 设置为 0。

#### GOOS 和 GOARCH
GOOS 是目标平台的操作系统，如 linux、windows，注意 macOS 的值是 darwin。默认是当前操作系统。

GOARCH 是目标平台的 CPU 架构，如 amd64、arm、386 等。默认值是当前平台的 CPU 架构。

Go 支持的所有操作系统和 CPU 架构可以查看 syslist.go 。

我们可以使用 go env 命令获取当前 GOOS 和 GOARCH 的值。例如我当前的操作系统是 linux：
```
$ go env
set GOARCH=amd64
set GOOS=linux
```
所以在本文的多阶段构建 Dockerfile.multistage 中，构建命令是：
```
RUN CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o /docker-go-server-ping
```

### 构建镜像
由于我们现在有两个 Dockerfile，所以我们必须告诉 Docker 我们要使用新的 Dockerfile 进行构建。
```
$ docker build -t docker-go-server-ping:v1 -f Dockerfile.multistage .
[+] Building 14.1s (13/13) FINISHED
 => [internal] load build definition from Dockerfile.multistage                                            0.0s 
 => => transferring dockerfile: 444B                                                                       0.0s 
 => [internal] load .dockerignore                                                                          0.0s 
 => => transferring context: 2B                                                                            0.0s 
 => [internal] load metadata for docker.io/library/golang:1.18.7-alpine3.16                                0.0s 
 => [build 1/7] FROM docker.io/library/golang:1.18.7-alpine3.16                                            0.0s 
 => [internal] load build context                                                                          0.0s 
 => => transferring context: 114B                                                                          0.0s 
 => CACHED [build 2/7] WORKDIR /app                                                                        0.0s 
 => CACHED [build 3/7] COPY go.mod ./                                                                      0.0s 
 => CACHED [build 4/7] COPY go.sum ./                                                                      0.0s 
 => CACHED [build 5/7] RUN go mod download                                                                 0.0s 
 => CACHED [build 6/7] COPY *.go ./                                                                        0.0s 
 => [build 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-go-server-ping              12.9s 
 => [stage-1 1/2] COPY --from=build /docker-go-server-ping /docker-go-server-ping                          0.1s 
 => exporting to image                                                                                     0.2s 
 => => exporting layers                                                                                    0.2s 
 => => writing image sha256:6de21360b16454d4ea8c99eb048b7a167be78b06bb1cbf3a4294a3b4c79e274e               0.0s 
 => => naming to docker.io/library/docker-go-server-ping:v1                                                0.0s 
```
构建完成后，你会发现 docker-go-server-ping:multistage 只有不到 8MB，比 docker-go-server-ping:latest 小了几十倍。
```
$ docker images | grep ping
docker-go-server-ping             v1                  6de21360b164   2 minutes ago   7.44MB
docker-go-server-ping             latest              6df657b11bba   2 hours ago     422MB
```

### 运行 Go 镜像
现在我们有了 Go 应用的镜像，接下来就可以运行 Go 镜像查看应用程序是否正常运行。

要在 Docker 容器中运行镜像，我们可以使用 docker run 命令，参数是镜像名称：
```
$ docker run -p 8080:8080 docker-go-server-ping:v1

   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.9.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:8080
```
可以看到 Go 应用成功启动了。

让我们再打开一个新的终端，通过 curl 向 Go 服务器发起一个请求：
```
$ curl http://localhost:8080/
Hello, Docker! <3
```


## 使用 Makefile 简化构建流程
在前面的步骤中，我们使用到了非常多的命令，维护起来非常麻烦，这时我们就可以使用 make 来简化构建流程。

make 是一个自动化构建工具，会在根据当前目录下名为 Makefile（或 makefile）的文件来执行相应的构建任务。

所以让我们先创建一个 Makefile 文件，内容如下：
```
APP=docker-go-server-ping

all: clean test build-docker-multistage

test:
	go test -v *.go

run: clean
	go build -o ${APP}
	./${APP}

deps:
	go mod download

clean:
	go clean

build-docker:
	docker build -t ${APP}:latest -f Dockerfile .

build-docker-multistage:
	docker build -t ${APP}:multistage -f Dockerfile.multistage .
```

接下来就可以通过 make 命令进行测试或构建了。

例如：

- make：执行 all 中定义的命令

- make test：执行单元测试

- make build-docker：构建 Docker 镜像

- make build-docker-multistage：多阶段构建镜像，构建的镜像通常用于生产环境

当然你也可以在 Makefile 中定义其他命令。
```
$ make build-docker-multistage
docker build -t docker-go-server-ping:multistage -f Dockerfile.multistage .
[+] Building 0.1s (13/13) FINISHED
 => [internal] load build definition from Dockerfile.multistage                                                                                                                                                                    0.0s 
 => => transferring dockerfile: 43B                                                                                                                                                                                                0.0s 
 => [internal] load .dockerignore                                                                                                                                                                                                  0.0s 
 => => transferring context: 2B                                                                                                                                                                                                    0.0s 
 => [internal] load metadata for docker.io/library/golang:1.18.7-alpine3.16                                                                                                                                                        0.0s 
 => [build 1/7] FROM docker.io/library/golang:1.18.7-alpine3.16                                                                                                                                                                    0.0s 
 => [internal] load build context                                                                                                                                                                                                  0.0s 
 => => transferring context: 114B                                                                                                                                                                                                  0.0s 
 => CACHED [build 2/7] WORKDIR /app                                                                                                                                                                                                0.0s 
 => CACHED [build 3/7] COPY go.mod ./                                                                                                                                                                                              0.0s 
 => CACHED [build 4/7] COPY go.sum ./                                                                                                                                                                                              0.0s 
 => CACHED [build 5/7] RUN go mod download                                                                                                                                                                                         0.0s 
 => CACHED [build 6/7] COPY *.go ./                                                                                                                                                                                                0.0s 
 => CACHED [build 7/7] RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /docker-go-server-ping                                                                                                                                0.0s 
 => CACHED [stage-1 1/2] COPY --from=build /docker-go-server-ping /docker-go-server-ping                                                                                                                                           0.0s 
 => exporting to image                                                                                                                                                                                                             0.0s 
 => => exporting layers                                                                                                                                                                                                            0.0s 
 => => writing image sha256:6de21360b16454d4ea8c99eb048b7a167be78b06bb1cbf3a4294a3b4c79e274e                                                                                                                                       0.0s 
 => => naming to docker.io/library/docker-go-server-ping:multistage                                                                                                                                                                0.0s
```


## 总结
在本文中，我们首先开发了一个简单的 Go REST 服务应用，然后针对该应用详细讲解了如何构建 Docker 镜像。要构建镜像首先需要编写 Dockerfile，但基础的 Dockerfile 体积过大，所以我们又学习了如何通过多阶段构建减小镜像体积。在多阶段构建时，由于构建机和部署服务器可能存在操作系统和 CPU 架构的差异，又学习了如何通过交叉编译构建出可在其他平台直接使用的二进制文件。最后由于整个构建流程涉及命令比较多，真实 Go 项目可能构建流程会更复杂，所以学习了如何通过 Makefile 简化构建流程。

## 参考
- Docker docks - Build your Go image (https://docs.docker.com/language/golang/build-images/#meet-the-example-application)

- olliefr/docker-gs-ping (https://github.com/olliefr/docker-gs-ping)

- nodejh/docker-go-server-ping (https://github.com/nodejh/docker-go-server-ping)
