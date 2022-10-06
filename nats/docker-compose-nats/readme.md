
# Docker Compose + NATS: Microservices Development Made Easy


Full example: https://gist.github.com/wallyqs/7f72efdc3fd6371364f8b28cbe32c5ee

https://nats.io/blog/docker-compose-plus-nats/

自动生成dockerfile文件
```azure
goctl docker -go main.go
```

生成docker镜像
```azure
docker build -t hello:v1 -f Dockerfile .
```


执行docker镜像
```azure
docker run -it --rm hello:v1
```

docker-compose生成镜像
```azure
docker-compose -f build.yml build
```

docker-compose开始自行镜像
```azure
docker-compose -f build.yml up -d
```

测试
```azure
$ curl 127.0.0.1:8080/createTask 
Task scheduled in 529.689µs
Response: Done!
$ curl 127.0.0.1:8080/
Basic NATS based microservice example v0.0.1
$ curl 127.0.0.1:8080
Basic NATS based microservice example v0.0.1
$ curl 127.0.0.1:8080/healthz
OK
```
