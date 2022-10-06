# bcrypt加密的密码配置

## 用nats客户端生成bcrypt密码
```azure
password: 1234567890qwertyuiopasdfghjkl
    
$ nats server passwd
? Enter password [? for help] *****************************
? Reenter password [? for help] *****************************
$2a$11$CPcnFq9kXMO1z6qYLr2hZuyYDpGLxzPM8fymTAxdzmgglpzlvVIHi
```

## 将生成的加密密码配置到服务端
```azure
jetstream: enabled
http_port: 8222

authorization {
    user: alexchen
    password: "$2a$11$CPcnFq9kXMO1z6qYLr2hZuyYDpGLxzPM8fymTAxdzmgglpzlvVIHi"
}
```

## 执行docker-compose

## 执行客户端

#### 认证失败
```azure
$ nats pub hello world
nats: error: nats: Authorization Violation
```
#### 认证成功
```azure
$ nats pub hello world --user alexchen --password 1234567890qwertyuiopasdfghjkl
13:35:09 Published 5 bytes to "hello"
```