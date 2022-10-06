# tls连接

## 生成证书和私钥
```azure
[root@localhost ~]# mkcert 192.168.0.110 example.com localhost 127.0.0.1 ::1
Note: the local CA is not installed in the system trust store.
Note: the local CA is not installed in the Firefox and/or Chrome/Chromium trust store.
Run "mkcert -install" for certificates to be trusted automatically 

Created a new certificate valid for the following names
 - "192.168.0.110"
 - "example.com"
 - "localhost"
 - "127.0.0.1"
 - "::1"

The certificate is at "./192.168.0.110+4.pem" and the key at "./192.168.0.110+4-key.pem"

It will expire on 7 September 2023
```

## docker-compose执行
```azure
docker-compose up -d
```

## 执行客户端

#### 认证失败
```azure
$ nats pub hello world
nats: error: x509: certificate signed by unknown authority
```

#### 认证成功(ca证书未被加入本机信任区)
```azure
$ nats pub hello world --tlsca ./certs/rootCA.pem
17:17:47 Published 5 bytes to "hello"
```

#### 将ca证书加入到本机的信任列表中去
```azure
mkcert --install
```

#### 认证成功
```azure
$ nats pub hello world 
17:21:41 Published 5 bytes to "hello"
```