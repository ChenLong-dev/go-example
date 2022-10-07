# 授权（能干啥）

**NATS支持对每一个用户进行主题级别的授权，即允许谁发布或订阅那个主题**
```js
# 官方文档地址
https://docs.nats.io/running-a-nats-service/configuration/securing_nats/authorization
```

## 有四种方式进行配置

### 方式一：变量配置（Variables）
```
authorization {
    default_permissions = {
        publish = "SANDBOX.*"
        subscribe = ["PUBLIC.>", "_INBOX.>"]
    }
    ADMIN = {
        publish = ">"
        subscribe = ">"
    }
    REQUESTOR = {
        publish = ["req.a", "req.b"]
        subscribe = "_INBOX.>"
    }
    RESPONDER = {
        publish = ["req.a", "req.b"]
        subscribe = "_INBOX.>"
    }
    users = [
        {user: admin, password:$ADMIN_PASS, permissions:$ADMIN}
        {user: client, password:$CLIENT_PASS, permissions:$REQUESTOR}
        {user: service, password:$SERVICE_PASS, permissions:$RESPONDER}
        {user: other, password:$OTHER_PASS}
    ]
}
```
注：
- 配置角色，如：ADMIN，REQUESTOR，RESPONDER
- admin 有$ADMIN授权，可以在任何主题上进行发布和订阅，用“>”来匹配任何主题
- client 是一个$REQUESTOR授权，可以在主题req.a和req.b中发布request，可以订阅主题_INBOX.>中的任何一个response
- service是一个对主题req.a和req.b的$REQUESTOR授权
- _INBOX.> 是用来做req和resp模式的消息回复
- user中没有指定角色就使用默认配置$default_permissions
- admin has ADMIN permissions and can publish/subscribe on any subject. We use the wildcard > to match any subject.
- client is a REQUESTOR and can publish requests on subjects req.a or req.b, and subscribe to anything that is a response (_INBOX.>). 
- service is a RESPONDER to req.a and req.b requests, so it needs to be able to subscribe to the request subjects and respond to client's that can publish requests to req.a and req.b. The reply subject is an inbox. Typically inboxes start with the prefix _INBOX. followed by a generated string. The _INBOX.> subject matches all subjects that begin with _INBOX..
- other has no permissions granted and therefore inherits the default permission set.

### 方式二：指定允许或拒绝的主题（Allow/Deny Specified）
```
authorization: {
    users = [
        {
            user: admin
            password: secret
            permissions: {
                publish: ">"
                subscribe: ">"
            }
        }
        {
            user: test
            password: test
            permissions: {
                publish: {
                    deny: ">"
                },
                subscribe: {
                    allow: "client.>"
                }
            }
        }
    ]
}
```

### 方式三：response允许（allow_responses）
```
authorization: {
    users: [
        { user: a, password: a },
        { user: b, password: b, permissions: {subscribe: "q", allow_responses: true } },
        { user: c, password: c, permissions: {subscribe: "q", allow_responses: { max: 5, expires: "1m" } } }
        { user: d, password: d, permissions: {subscribe: "q", publish: "x", allow_responses: true } }
    ]
}
```
- 用户a没有任何限制。 
- 用户b可以监听q请求，并且只能发布一次以回复主题。当allow_responses被设置时，所有其他发布主题都被隐式拒绝。 
- 用户c可以监听q请求，但最多只能返回5条应答消息，应答主题最多可以发布1分钟。 
- 用户d具有与用户b相同的行为，只是它也可以显式地发布到主题x，这将覆盖allow_responses的隐式拒绝。

### 方式四：队列授权（Queue Permissions）
```
users = [
  {
    user: "a", password: "a", permissions: {
      sub: {
        allow: ["foo queue"]
     }
  }
  {
    user: "b", password: "b", permissions: {
      sub: {
        # Allow plain subscription foo, but only v1 groups or *.dev queue groups
        allow: ["foo", "foo v1", "foo v1.>", "foo *.dev"]

        # Prevent queue subscriptions on prod groups
        deny: ["> *.prod"]
     }
  }
]
```

## 启动nats-server
```azure
$ docker-compose up
```

## 测试验证

#### 认证失败
```azure
$ nats pub hello world
nats: error: nats: Authorization Violation
```

```azure
$ nats pub SANDBOX.hi world --user client --password client
11:16:58 Unexpected NATS error from server nats://127.0.0.1:4222: nats: Permissions Violation for Publish to "SA
NDBOX.hi"
nats: error: nats: Permissions Violation for Publish to "SANDBOX.hi"
```

#### 认证成功
admin用户是允许发送到所有主题
```azure
$ nats pub hello world --user admin --password admin                         
11:14:07 Published 5 bytes to "hello"
```

```azure
$ nats pub SANDBOX.hi world --user other --password other
11:16:33 Published 5 bytes to "SANDBOX.hi"
```

```azure
$ nats pub req.a world --user client --password client
11:18:05 Published 5 bytes to "req.a"
```