# token认证

# 执行docker-compose
```azure
$ docker-compose up
[+] Running 1/0
 - Container token-nats-1  Recreated                                                                       0.1s 
Attaching to token-nats-1
token-nats-1  | [1] 2022/10/06 05:00:55.542453 [INF] Starting nats-server
token-nats-1  | [1] 2022/10/06 05:00:55.542524 [INF]   Version:  2.9.2
token-nats-1  | [1] 2022/10/06 05:00:55.542527 [INF]   Git:      [6d81dde]
token-nats-1  | [1] 2022/10/06 05:00:55.542529 [INF]   Name:     NCU74TC4BQQ6IMI6CUG4P7RX7SJDFVEA7MITVMRLZVPYI2V
DL3XDAQSA
token-nats-1  | [1] 2022/10/06 05:00:55.542533 [INF]   Node:     003x2GU7
token-nats-1  | [1] 2022/10/06 05:00:55.542534 [INF]   ID:       NCU74TC4BQQ6IMI6CUG4P7RX7SJDFVEA7MITVMRLZVPYI2V
DL3XDAQSA
token-nats-1  | [1] 2022/10/06 05:00:55.542557 [INF] Using configuration file: ./nats-server.conf
token-nats-1  | [1] 2022/10/06 05:00:55.543043 [INF] Starting http monitor on 0.0.0.0:8222
token-nats-1  | [1] 2022/10/06 05:00:55.543116 [INF] Starting JetStream
token-nats-1  | [1] 2022/10/06 05:00:55.543556 [INF]     _ ___ _____ ___ _____ ___ ___   _   __  __
token-nats-1  | [1] 2022/10/06 05:00:55.543579 [INF]  _ | | __|_   _/ __|_   _| _ \ __| /_\ |  \/  |
token-nats-1  | [1] 2022/10/06 05:00:55.543581 [INF] | || | _|  | | \__ \ | | |   / _| / _ \| |\/| |
token-nats-1  | [1] 2022/10/06 05:00:55.543582 [INF]  \__/|___| |_| |___/ |_| |_|_\___/_/ \_\_|  |_|
token-nats-1  | [1] 2022/10/06 05:00:55.543583 [INF]
token-nats-1  | [1] 2022/10/06 05:00:55.543584 [INF]          https://docs.nats.io/jetstream
token-nats-1  | [1] 2022/10/06 05:00:55.543585 [INF]
token-nats-1  | [1] 2022/10/06 05:00:55.543586 [INF] ---------------- JETSTREAM ----------------
token-nats-1  | [1] 2022/10/06 05:00:55.543589 [INF]   Max Memory:      9.22 GB
token-nats-1  | [1] 2022/10/06 05:00:55.543591 [INF]   Max Storage:     175.06 GB
token-nats-1  | [1] 2022/10/06 05:00:55.543592 [INF]   Store Directory: "/tmp/nats/jetstream"
token-nats-1  | [1] 2022/10/06 05:00:55.543594 [INF] -------------------------------------------
token-nats-1  | [1] 2022/10/06 05:00:55.543990 [INF] Listening for client connections on 0.0.0.0:4222
token-nats-1  | [1] 2022/10/06 05:00:55.544272 [INF] Server is ready
```

## 执行客户端

#### 认证失败
```azure
$ nats pub hello woeld
nats: error: nats: Authorization Violation
```
#### 认证成功
```azure
$ nats -s "nats://mytoken@localhost:4222" pub hello world
13:04:54 Published 5 bytes to "hello"
```