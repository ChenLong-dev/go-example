# Portainer

[Portainer](https://www.portainer.io/)

## 简单安装portainer
### docker run
```
$ docker volume create portainer_data
$ docker run --name portainer -d -p 8000:8000 -p 9000:9000 -v /var/run/docker.sock:/var/run/docker.sock -v portainer_data:/data portainer/portainer
```

```
Ensure that you have started the Portainer container with the following Docker flag:

-v "/var/run/docker.sock:/var/run/docker.sock" (Linux).

or

-v \\.\pipe\docker_engine:\\.\pipe\docker_engine (Windows).
```

### docker-compose
```
version: "3.8"
services:
  portainer:
    image: portainer/portainer:latest
    container_name: portainer
    ports:
      - "9000:9000"
    volumes:
      - /home/app/portainer/data:/data
      - /var/run/docker.sock:/var/run/docker.sock
```

web页面登录http://ip:9000，初始密码设置（sangfor@123）

### 查看容器运行命令
```
docker pull cucker/get_command_4_run_container
```
命令格式：
```
docker run --rm -v /var/run/docker.sock:/var/run/docker.sock cucker/get_command_4_run_container [容器名称]/[容器ID]
```

可以将其封装为一个别名：
```
echo "alias get_run_command='docker run --rm -v /var/run/docker.sock:/var/run/docker.sock cucker/get_command_4_run_container'" >> ~/.bashrc \
&& \
. ~/.bashrc
```
这样就可以这样用了：get_run_command [容器名称]/[容器ID]

![img.png](img.png)

