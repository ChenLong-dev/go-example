
`https://docs.nats.io/running-a-nats-service/configuration/securing_nats`

# mkcert使用方法

## 安装mkcert
```azure
https://github.com/FiloSottile/mkcert
```

## 使用说明
```azure
[root@localhost ~]# mkcert 192.168.0.110 example.com localhost 127.0.0.1 ::1
Note: the local CA is not installed in the system trust store.
Note: the local CA is not installed in the Firefox and/or Chrome/Chromium trust store.
Run "mkcert -install" for certificates to be trusted automatically 

Created a new certificate valid for the following names
 - "192.168.128.134"
 - "example.com"
 - "localhost"
 - "127.0.0.1"
 - "::1"

The certificate is at "./192.168.128.134+4.pem" and the key at "./192.168.128.134+4-key.pem"

It will expire on 7 September 2023
```

```azure
mkcert -cert-file ./ssl/server.crt -key-file ./ssl/server.key 192.168.0.110 example.com localhost 127.0.0.1 ::1
```