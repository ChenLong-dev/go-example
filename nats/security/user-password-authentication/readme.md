# 用户名密码认证

### 执行docker-compose
```azure
$ docker-compose up
[+] Running 2/2
 - Network user-password_default   Created                                                                 0.9s 
 - Container user-password-nats-1  Created                                                                 0.1s 
Attaching to user-password-nats-1
user-password-nats-1  | [1] 2022/10/06 05:09:55.907635 [INF] Starting nats-server
user-password-nats-1  | [1] 2022/10/06 05:09:55.907705 [INF]   Version:  2.9.2
user-password-nats-1  | [1] 2022/10/06 05:09:55.907708 [INF]   Git:      [6d81dde]
user-password-nats-1  | [1] 2022/10/06 05:09:55.907710 [INF]   Name:     NCDO6GIGVEU746HJOEVNIE4VWH5LZQWFIVWVNRU
KWVFOJTXD5ICIMRNA
user-password-nats-1  | [1] 2022/10/06 05:09:55.907716 [INF]   Node:     BIosDyuN
user-password-nats-1  | [1] 2022/10/06 05:09:55.907717 [INF]   ID:       NCDO6GIGVEU746HJOEVNIE4VWH5LZQWFIVWVNRU
KWVFOJTXD5ICIMRNA
user-password-nats-1  | [1] 2022/10/06 05:09:55.907719 [WRN] Plaintext passwords detected, use nkeys or bcrypt  
user-password-nats-1  | [1] 2022/10/06 05:09:55.907744 [INF] Using configuration file: nats-server.conf
user-password-nats-1  | [1] 2022/10/06 05:09:55.908302 [INF] Starting http monitor on 0.0.0.0:8222
user-password-nats-1  | [1] 2022/10/06 05:09:55.908356 [INF] Starting JetStream
user-password-nats-1  | [1] 2022/10/06 05:09:55.908800 [INF]     _ ___ _____ ___ _____ ___ ___   _   __  __     
user-password-nats-1  | [1] 2022/10/06 05:09:55.908823 [INF]  _ | | __|_   _/ __|_   _| _ \ __| /_\ |  \/  |    
user-password-nats-1  | [1] 2022/10/06 05:09:55.908825 [INF] | || | _|  | | \__ \ | | |   / _| / _ \| |\/| |    
user-password-nats-1  | [1] 2022/10/06 05:09:55.908825 [INF]  \__/|___| |_| |___/ |_| |_|_\___/_/ \_\_|  |_|    
user-password-nats-1  | [1] 2022/10/06 05:09:55.908826 [INF]
user-password-nats-1  | [1] 2022/10/06 05:09:55.908827 [INF]          https://docs.nats.io/jetstream
user-password-nats-1  | [1] 2022/10/06 05:09:55.908828 [INF]
user-password-nats-1  | [1] 2022/10/06 05:09:55.908828 [INF] ---------------- JETSTREAM ----------------        
user-password-nats-1  | [1] 2022/10/06 05:09:55.908831 [INF]   Max Memory:      9.22 GB
user-password-nats-1  | [1] 2022/10/06 05:09:55.908833 [INF]   Max Storage:     175.06 GB
user-password-nats-1  | [1] 2022/10/06 05:09:55.908834 [INF]   Store Directory: "/tmp/nats/jetstream"
user-password-nats-1  | [1] 2022/10/06 05:09:55.908834 [INF] -------------------------------------------        
user-password-nats-1  | [1] 2022/10/06 05:09:55.909213 [INF] Listening for client connections on 0.0.0.0:4222   
user-password-nats-1  | [1] 2022/10/06 05:09:55.909404 [INF] Server is ready
```
## 执行客户端

#### 认证失败
```azure
$ nats pub hello world
nats: error: nats: Authorization Violation
```
#### 认证成功
```azure
$ nats pub hello world --user alexchen --password mypassword
13:24:21 Published 5 bytes to "hello"
```

