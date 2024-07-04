# cgroups使用方法

```
https://blog.csdn.net/RNGWGzZs/article/details/132301794

https://www.jianshu.com/p/910d00b6fe00

https://access.redhat.com/documentation/zh-cn/red_hat_enterprise_linux/7/html/resource_management_guide/index

https://segmentfault.com/a/1190000008125359

https://github.com/0voice/kernel_new_features/tree/main/cgroups

https://adtxl.com/index.php/archives/179.html

```

## 什么是cgroups
cgroups是 linux 内核提供的一种机制， 这种机制可以根据需求把一系列系统任务及其子任务整合(或分隔)到按资源划分等级的不同组内，从而为系统资源管理提供一个统一的框架。简单说， cgroups 就是可以控制、记录任务组所使用的物理资源。

本质上来说:  " cgroups 是内核附加在程序上的一系列钩子(hook) ,通过程序运行时对资源的调度触发相应的钩子以达到资源追踪和限制的目的 "。   

Hook原理

    Hook其本质就是劫持函数调用。但是由于处于Linux用户态，每个进程都有自己独立的进程空间，所以必须先注入到所要Hook的进程空间，修改其内存中的进程代码，替换其过程表的符号地址。

## 为什么使用cgroups？
cgroups可以做到对CPU、内存等等系统资源进行精细化管控！目前，很风靡的轻量级容器Docker及 k8s 中的 pod 就使用了 cgroups 提供的资源限制能力来完成 cpu、内存等部分资源的管控。

例如，如果你想在一台8核服务器上部署一个服务，其中你可以通过 cgroups 限制业务处理仅仅用掉完全够用的6个核，其余剩下的也可以另作他用。

## cgroups的用途作用
- Resource limitation: 限制资源使用.例:内存使用上限/cpu 的使用限制。
- Prioritization: 优先级控制. 例:CPU 利用/磁盘 IO 吞吐。
- Accounting: 一些审计或一些统计。
- Control: 挂起进程/恢复执行进程。
  
使用cgroups控制的子系统:

| 名称         | 作用                                                             |
|------------|----------------------------------------------------------------|
| blkio      | 对块设备的IO进行限制                                                    |
| cpu        | 限制cpu时间片分配                                                     |
| cpuacct    | 生成cgroup中的任务占用cpu资源的报告，与cpu挂载到同一目录                             |
| cpuset     | 给cgroup中的任务分配独立的cpu(多核)和内存节点。                                  |
| devices    | 限制设备文件创建和对设备文件的读写                                              |
| freezer    | 暂停、恢复cgroup中的任务                                                |
| memory     | 对cgroup中的任务的可用内存进行限制，并自动生成占用资源报告                               |
| perf_event | 允许perf观测cgroup中的task                                           |
| net_cls    | cgroup中的任务创建数据报文的类别标识符，让Linux流量控制器可以识别来自特定cgroup任务的数据包，并进行网络限制 |
| hugetlb    | 限制使用内存页的数量                                                     |
| pids       | 限制任务数量                                                         |
| rdma       | 限制远程直接数据存取                                                     |

## 子系统接口/参数
### cpu子系统：于限制进程的 CPU 利用率
cpu.shares：cpu比重分配。通过一个整数的数值来调节cgroup所占用的cpu时间。例如，有2个cgroup（假设为CPU1，CPU2），其中一个(CPU1)cpu.shares设定为100另外一个(CPU2)设为200，那么CPU2所使用的cpu时间将是CPU1所使用时间的2倍。cpu.shares 的值必须为2或者高于2

cpu.cfs_period_us：规定CPU的时间周期(单位是微秒)。最大值是1秒，最小值是1000微秒。如果在一个单CPU的系统内，要保证一个cgroup 内的任务在1秒的CPU周期内占用0.2秒的CPU时间，可以通过设置cpu.cfs_quota_us 为200000和cpu.cfs_period_us 为 1000000

cpu.cfs_quota_us：在单位时间内（即cpu.cfs_period_us设定值）可用的CPU最大时间（单位是微秒）。cpu.cfs_quota_us值可以大于cpu.cfs_period_us值，例如在一个双CPU的系统内，想要一个cgroup内的进程充分的利用2个CPU，可以设定cpu.cfs_quota_us为 200000 及cpu.cfs_period_us为 100000

当设定cpu.cfs_quota_us为-1时，表明不受限制，同时这也是默认值

### cpuacct子系统：统计各个 Cgroup 的 CPU 使用情况
cpuacct.stat：cgroup中所有任务的用户和内核分别使用CPU的时长

cpuacct.usage：cgroup中所有任务的CPU使用时长（纳秒）

cpuacct.usage_percpu：cgroup中所有任务使用的每个cpu的时间（纳秒）

### cpuset子系统：为一组进程分配指定的CPU和内存节点
cpuset.cpus：允许cgroup中的进程使用的CPU列表。如0-2,16代表 0,1,2,16这4个CPU

cpuset.mems：允许cgroup中的进程使用的内存节点列表。如0-2,16代表 0,1,2,16这4个可用节点

cpuset.memory_migrate：当cpuset.mems变化时内存页上的数据是否迁移（默认值0，不迁移；1，迁移）

cpuset.cpu_exclusive：cgroup是否独占cpuset.cpus 中分配的cpu 。（默认值0，共享；1，独占），如果设置为1，其他cgroup内的cpuset.cpus值不能包含有该cpuset.cpus内的值

cpuset.mem_exclusive：是否独占memory，（默认值0，共享；1，独占）

cpuset.mem_hardwall：cgroup中任务的内存是否隔离，（默认值0，不隔离；1，隔离，每个用户的任务将拥有独立的空间）

cpuset.sched_load_balance：cgroup的cpu压力是否会被平均到cpuset中的多个cpu上。（默认值1，启用负载均衡；0，禁用。）

### memory子系统：限制cgroup所能使用的内存上限
memory.limit_in_bytes：设定最大的内存使用量，可以加单位（k/K,m/M,g/G）不加单位默认为bytes

memory.soft_limit_in_bytes：和 memory.limit_in_bytes 的差异是，这个限制并不会阻止进程使用超过限额的内存，只是在系统内存不足时，会优先回收超过限额的进程占用的内存，使之向限定值靠拢。该值应小于memory.limit_in_bytes设定值

memory.stat：统计内存使用情况。各项单位为字节

memory.memsw.limit_in_bytes：设定最大的内存+swap的使用量

memory.oom_control：当进程出现Out of Memory时，是否进行kill操作。默认值0，kill；设置为1时，进程将进入睡眠状态，等待内存充足时被唤醒

memory.force_empty：当设置为0时，清空该group的所有内存页；该选项只有在当前group没有tasks才可以使用

### blkio子系统：限制cgroup对IO的使用
blkio.weight：设置权值，范围在[100, 1000]，属于比重分配，不是绝对带宽。因此只有当不同 Cgroup 争用同一个 阻塞设备时才起作用

blkio.weight_device：对具体设备设置权值。它会覆盖上面的选项值

blkio.throttle.read_bps_device：对具体的设备，设置每秒读磁盘的带宽上限

blkio.throttle.write_bps_device：对具体的设备，设置每秒写磁盘的带宽上限

blkio.throttle.read_iops_device：对具体的设备，设置每秒读磁盘的IOPS带宽上限

blkio.throttle.write_iops_device：对具体的设备，设置每秒写磁盘的IOPS带宽上限

### devices子系统：限定cgroup内的进程可以访问的设备
devices.allow：允许访问的设备。文件包括4个字段：type（设备类型）, major（主设备号）, minor（次设备号）, and access（访问方式）

#### type
- a — 适用所有设备，包括字符设备和块设备
- b — 块设备
- c — 字符设备
#### major, minor
- 9:*
- *:*
- 8:1
#### access
- r — 读
- w — 写
- m — 创建不存在的设备
- 
devices.deny：禁止访问的设备，格式同devices.allow

devices.list：显示目前允许被访问的设备列表

### freezer子系统：暂停或恢复任务
freezer.state：当前cgroup中进程的状态

FROZEN：挂起进程

FREEZING：进程正在挂起中

THAWED：激活进程

- 挂起进程时，会连同子进程一同挂起。
- 不能将进程移动到处于FROZEN状态的cgroup中。
- 只有FROZEN和THAWED可以被写进freezer.state中, FREEZING则不能



```
[root@localhost cgroup]# pwd
/sys/fs/cgroup
[root@localhost cgroup]# ls
blkio  cpu  cpuacct  cpu,cpuacct  cpuset  devices  freezer  hugetlb  memory  net_cls  net_cls,net_prio  net_prio  perf_event  pids  rdma  systemd
```

## cgroups资源控制实战

cgroups的基本信息有哪些？cgroups是如何进行资源管理、控制的呢？作为真正的资源控制层——OS，它同cgroups是怎样运作的？

### 基础知识:
#### pidstat:

pidstat 是 sysstat(Linux系统性能监控工具) 的一个命令，用于监控全部或指定进程的 CPU、内存、线程、设备、IO等系统资源的占用情况。

语法:

    pidstat [option] [ <时间间隔>] [次数]
    注：用户可以通过指定统计的次数和时间来获得所需的统计信息.

参数：

| 参数	     | 作用                    |
|---------|-----------------------|
| -u	     | 默认参数，显示各进程的 CPU 使用统计. |
| -r	     | 显示各进程的内存使用统计.         |
| -d	     | 显示各进程的 IO 使用情况.       |
| -p:ALL	 | 指定进程号,ALL 表示所有进程.     |
| -C      | 指定命令，这个字符串可以是正则表达式.   |
| -l	     | 显示命令名和所有参数.           |

安装\卸载（Centos7）

    yum install(remove) sysstat -y

示例: 我们可以查看Linux下的authserver服务资源使用情况

默认查看CPU资源（pidstat -p ALL -C authserver）
```
[root@localhost ~]# pidstat -p ALL -C authserver
Linux 4.18.0-348.el8.x86_64 (localhost.localdomain)     11/28/2023      _x86_64_        (8 CPU)

03:16:22 PM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
03:16:22 PM     0    344649    0.00    0.00    0.00    0.00    0.00     6  authserver
```

内存资源使用情况（pidstat -r -p ALL -C authserver）
```
[root@localhost ~]# pidstat -r -p ALL -C authserver
Linux 4.18.0-348.el8.x86_64 (localhost.localdomain)     11/28/2023      _x86_64_        (8 CPU)

03:18:41 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
03:18:41 PM     0    344649      0.00      0.00 1275440   38556   0.24  authserver
```

IO使用情况（pidstat -d -p ALL -C authserver）
```
[root@localhost ~]# pidstat -d -p ALL -C authserver
Linux 4.18.0-348.el8.x86_64 (localhost.localdomain)     11/28/2023      _x86_64_        (8 CPU)

03:19:38 PM   UID       PID   kB_rd/s   kB_wr/s kB_ccwr/s iodelay  Command
03:19:38 PM     0    344649      0.00      0.00      0.00       0  authserver
```

显示命令行参数（pidstat -l -p ALL -C authserver）
```
[root@localhost ~]# pidstat -l -p ALL -C authserver
Linux 4.18.0-348.el8.x86_64 (localhost.localdomain)     11/28/2023      _x86_64_        (8 CPU)

03:20:53 PM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
03:20:53 PM     0    344649    0.00    0.00    0.00    0.00    0.00     0  ./authserver serve -c ./etc/config.yaml 
03:20:53 PM     0    412504    0.00    0.00    0.00    0.00    0.00     0  pidstat -l -p ALL -C authserver
```

#### stress:
stress 是 Linux 的一个压力测试工具，可以对 CPU、 Memory、 IO、磁盘进行压力测
试。

语法:

    stress [ option ] [ARG] ]

参数:

| 参数               | 作用                                                                                                               |
|------------------|------------------------------------------------------------------------------------------------------------------|
| -c, --cpu N		    | 产生 N 个进程，每个进程都循环调用 sqrt 函数产生 CPU 压力.                                                                             |
| -i, --io N		     | 产生 N 个进程，每个进程循环调用 sync 将内存缓冲区内容写到磁盘上，产生 IO 压力.                                                                   |
| -m, --vm			      | 产生 N 个进程，每个进程循环调用 malloc/free 函数分配和释放内存。--vm-bytes B: 指定分配内存的大小。--vm-keep: 一直占用内存，区别于不断的释放和重新分配(默认是不断释放并重新分配内存)。 |
| -d, --hdd N		    | 产生 N 个不断执行 write 和 unlink 函数的进程（创建文件，写入内容，删除文件）--hdd-bytes B：指定文件大小                                              |
| -t, --timeout N	 | 在 N 秒后结束程序                                                                                                       |
| -q, --quiet：		   | 程序在运行的过程中不输出信息                                                                                                   |

安装\卸载（Centos7）:

    yum install(remove) stress -y

## 实操一: cgroups 信息查看

### cgroups 版本查看（cat /proc/filesystems | grep cgroup）
```
[root@localhost ~]# cat /proc/filesystems | grep cgroup
nodev   cgroup
nodev   cgroup2
```

### cgroups 子系统查看（cat /proc/cgroups）
```
[root@localhost ~]# cat /proc/cgroups
#subsys_name    hierarchy       num_cgroups     enabled
cpuset              8               9               1
cpu                 5               114             1
cpuacct             5               114             1
blkio               10              114             1
memory              9               265             1
devices             12              114             1
freezer             11              9               1
net_cls             3               9               1
perf_event          4               9               1
net_prio            3               9               1
hugetlb             6               9               1
pids                2               152             1
rdma                7               1               1
```

### cgroups 挂载信息查看（mount |grep cgroup 或  mount -t cgroup）

可以看到默认存储位置为 /sys/fs/cgroup。

```
[root@localhost ~]# mount |grep cgroup
tmpfs on /sys/fs/cgroup type tmpfs (ro,nosuid,nodev,noexec,seclabel,mode=755)
cgroup on /sys/fs/cgroup/systemd type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,xattr,release_agent=/usr/lib/systemd/systemd-cgroups-agent,name=systemd)
cgroup on /sys/fs/cgroup/pids type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,pids)
cgroup on /sys/fs/cgroup/net_cls,net_prio type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,net_cls,net_prio)
cgroup on /sys/fs/cgroup/perf_event type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,perf_event)
cgroup on /sys/fs/cgroup/cpu,cpuacct type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,cpu,cpuacct)
cgroup on /sys/fs/cgroup/hugetlb type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,hugetlb)
cgroup on /sys/fs/cgroup/rdma type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,rdma)
cgroup on /sys/fs/cgroup/cpuset type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,cpuset)
cgroup on /sys/fs/cgroup/memory type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,memory)
cgroup on /sys/fs/cgroup/blkio type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,blkio)
cgroup on /sys/fs/cgroup/freezer type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,freezer)
cgroup on /sys/fs/cgroup/devices type cgroup (rw,nosuid,nodev,noexec,relatime,seclabel,devices)
```

### 查看一个进程上的cgroups的信息

####  以当前 shell 进程为例,查看进程的 cgroup

    cat /proc/$$/cgroup
    注: "$$"当前shell脚本 —> 命令行解释器进程

```
[root@localhost ~]# cat /proc/$$/cgroup
12:devices:/user.slice
11:freezer:/
10:blkio:/user.slice
9:memory:/user.slice/user-0.slice/session-73.scope
8:cpuset:/
7:rdma:/
6:hugetlb:/
5:cpu,cpuacct:/user.slice
4:perf_event:/
3:net_cls,net_prio:/
2:pids:/user.slice/user-0.slice/session-73.scope
1:name=systemd:/user.slice/user-0.slice/session-73.scope
```

#### 比如 cpu 在 user.slice，我们可以找到这个目录，里面有对 init 进程的详细限制信息

    ll /sys/fs/cgroup/cpu/user.slice/

```
[root@localhost ~]# ll /sys/fs/cgroup/cpu/user.slice/
total 0
-rw-r--r--. 1 root root 0 Nov 28 15:33 cgroup.clone_children
-rw-r--r--. 1 root root 0 Nov 21 11:27 cgroup.procs
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.stat
-rw-r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage_all
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage_percpu
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage_percpu_sys
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage_percpu_user
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage_sys
-r--r--r--. 1 root root 0 Nov 28 15:33 cpuacct.usage_user
-rw-r--r--. 1 root root 0 Nov 28 15:16 cpu.cfs_period_us
-rw-r--r--. 1 root root 0 Nov 21 14:54 cpu.cfs_quota_us
-rw-r--r--. 1 root root 0 Nov 28 15:33 cpu.rt_period_us
-rw-r--r--. 1 root root 0 Nov 28 15:33 cpu.rt_runtime_us
-rw-r--r--. 1 root root 0 Nov 28 15:16 cpu.shares
-r--r--r--. 1 root root 0 Nov 28 15:33 cpu.stat
-rw-r--r--. 1 root root 0 Nov 28 15:33 notify_on_release
-rw-r--r--. 1 root root 0 Nov 28 15:33 tasks
```

## 实操二:使用 cgroups 对内存进行控制

### 创建内存的 cgroup，很简单我们进入到 cgroup 的内存控制目录： “/sys/fs/cgroup/memory”，我们创建目录 climits。

```
[root@localhost memory]# mkdir climits
[root@localhost memory]# ls climits/
cgroup.clone_children  memory.force_empty              memory.kmem.slabinfo                memory.kmem.tcp.usage_in_bytes  memory.memsw.failcnt             memory.move_charge_at_immigrate  memory.soft_limit_in_bytes  memory.use_hierarchy
cgroup.event_control   memory.kmem.failcnt             memory.kmem.tcp.failcnt             memory.kmem.usage_in_bytes      memory.memsw.limit_in_bytes      memory.numa_stat                 memory.stat                 notify_on_release
cgroup.procs           memory.kmem.limit_in_bytes      memory.kmem.tcp.limit_in_bytes      memory.limit_in_bytes           memory.memsw.max_usage_in_bytes  memory.oom_control               memory.swappiness           tasks
memory.failcnt         memory.kmem.max_usage_in_bytes  memory.kmem.tcp.max_usage_in_bytes  memory.max_usage_in_bytes       memory.memsw.usage_in_bytes      memory.pressure_level            memory.usage_in_bytes
```

    cgroup.event_control       #用于eventfd的接口
    memory.usage_in_bytes      #显示当前已用的内存
    memory.limit_in_bytes      #设置/显示当前限制的内存额度
    memory.failcnt             #显示内存使用量达到限制值的次数
    memory.max_usage_in_bytes  #历史内存最大使用量
    memory.soft_limit_in_bytes #设置/显示当前限制的内存软额度
    memory.stat                #显示当前cgroup的内存使用情况
    memory.use_hierarchy       #设置/显示是否将子cgroup的内存使用情况统计到当前cgroup里面
    memory.force_empty         #触发系统立即尽可能的回收当前cgroup中可以回收的内存
    memory.pressure_level      #设置内存压力的通知事件，配合cgroup.event_control一起使用
    memory.swappiness          #设置和显示当前的swappiness
    memory.move_charge_at_immigrate #设置当进程移动到其他cgroup中时，它所占用的内存是否也随着移动过去
    memory.oom_control         #设置/显示oom controls相关的配置
    memory.numa_stat           #显示numa相关的内存

这里只要在这个目录里创建新目录，内存限制文件就会自动在其中进行创建了。cgroups 文件系统会在创建文件目录的时候自动创建相应的配置文件。

### 删除 cgroup
```
rmdir /sys/fs/cgroup/cpu/climits
```

### 配置 cgroup 的策略为最大使用 20M 内存（cd /sys/fs/cgroup/memory/climits/）

    echo "20971520" > ./memory.limit_in_bytes

```
[root@localhost memory]# cat memory.limit_in_bytes
9223372036854771712
[root@localhost climits]# echo "20971520" > ./memory.limit_in_bytes
[root@localhost climits]# cat memory.limit_in_bytes
20971520 
```

### 启动 1 个消耗内存的进程，每个进程占用 50M 内存

    stress -m 1 --vm-bytes 50M

```
[root@localhost climits]# stress -m 1 --vm-bytes 50M
stress: info: [428542] dispatching hogs: 0 cpu, 0 io, 1 vm, 0 hdd

```

我们打开一个新窗口监控该内存使用情况:

    pidstat -r -C stress -p ALL 1 10000

```
[root@localhost ~]# pidstat -r -C stress -p ALL 1 10000
Linux 4.18.0-348.el8.x86_64 (localhost.localdomain)     11/28/2023      _x86_64_        (8 CPU)

03:44:30 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
03:44:31 PM     0    428542      0.00      0.00    7976    1124   0.01  stress
03:44:31 PM     0    428543  53064.36      0.00   59180   22408   0.14  stress

03:44:31 PM   UID       PID  minflt/s  majflt/s     VSZ     RSS   %MEM  Command
03:44:32 PM     0    428542      0.00      0.00    7976    1124   0.01  stress
03:44:32 PM     0    428543  50393.00      0.00   59180   40840   0.25  stress
```
另启与一个窗口，并且将进程 id 移动到我们的 cgroup 策略(climits):

    /sys/fs/cgroup/memory/climits/tasks
    echo pro_pid >> tasks

```
[root@localhost climits]# ps -ef|grep stress
root      428542  376083  0 15:42 pts/5    00:00:00 stress -m 1 --vm-bytes 50M
root      428543  428542 99 15:42 pts/5    00:05:09 stress -m 1 --vm-bytes 50M

[root@localhost climits]# cat tasks
[root@localhost climits]# echo 428542 >> tasks
[root@localhost climits]# cat tasks
428542
```
一旦设置了内存限制，将立即生效，并且当物理内存使用量达到limit的时候，memory.failcnt的内容会加1，但这时进程不一定就会被kill掉，内核会尽量将物理内存中的数据移到swap空间上去，如果实在是没办法移动了（设置的limit过小，或者swap空间不足），默认情况下，就会kill掉cgroup里面继续申请内存的进程。如果将交换内存memory.swappiness设置为0，内存达到阈值就将进程kill掉。

## 实操三:使用 cgroups 对 cpu 进行控制（/sys/fs/cgroup/cpu）

### 基础知识
#### 限制进程可使用的CPU百分比
设置 CPU 数字的单位都是微秒，用us表示。

    cpu.cfs_period_us:时间周期长度，取值范围为1毫秒到1秒。
    cfs_quota_us：当前cgroup在设置的周期长度内所能使用的CPU时间。

两个文件配合起来设置CPU的使用上限。

示例：

1.限制只能使用1个CPU（每250ms（周期）能使用250ms的CPU时间）。

# echo 250000 > cpu.cfs_quota_us
# echo 250000 > cpu.cfs_period_us

2.限制使用2个CPU（内核）（每500ms（周期）能使用1000ms的CPU时间，即使用两个内核）

# echo 1000000 > cpu.cfs_quota_us
# echo 500000 > cpu.cfs_period_us

3.限制使用1个CPU的20%（每50ms（周期）能使用10ms的CPU时间，即使用一个CPU核心的20%）

# echo 10000 > cpu.cfs_quota_us
# echo 50000 > cpu.cfs_period_us

除了cpu.cfs_quota_us和cpu.cfs_period_us，

2.cpu.shares

用来设置CPU的相对值，并且是针对所有的CPU（内核），默认值是1024，假如系统中有两个cgroup，分别是A和B，A的shares值是1024，B的shares值是512，那么A将获得1024/(1204+512)=66%的CPU资源，而B将获得33%的CPU资源。

shares有两个特点:

- 如果A不忙，没有使用到66%的CPU时间，那么剩余的CPU时间将会被系统分配给B，即B的CPU使用率可以超过33%
- 如果添加了一个新的cgroup C，且它的shares值是1024，那么A的限额变成了1024/(1204+512+1024)=40%，B的变成了20%。

综上，我们看到shares是一个绝对值，需要和其他cgroup的值进行比较才能得到自己的相对限额。

### 创建内存的 cgroup，很简单我们进入到 cgroup 的内存控制目录“/sys/fs/cgroup/cpu”，并且继续执行像在memory目录之中的动作，创建climits目录。

```
[root@localhost cpu]# cd /sys/fs/cgroup/cpu
[root@localhost cpu]# mkdir climits
[root@localhost cpu]# ls climits/
cgroup.clone_children  cpuacct.stat   cpuacct.usage_all     cpuacct.usage_percpu_sys   cpuacct.usage_sys   cpu.cfs_period_us  cpu.rt_period_us   cpu.shares  notify_on_release
cgroup.procs           cpuacct.usage  cpuacct.usage_percpu  cpuacct.usage_percpu_user  cpuacct.usage_user  cpu.cfs_quota_us   cpu.rt_runtime_us  cpu.stat    tasks
```

### 设置 cproup 的 cpu 使用率为 30%， cpu 使用率为："cfs_quota_us/cfs_period_us "

- cfs_period_us: 表示一个 cpu 带宽，单位为“us”。系统总 CPU 带宽 ，默认值 10,0000。
- cfs_quota_us: 表示 Cgroup 可以使用的 cpu 的带宽，单位为“us”。cfs_quota_us 为-1，表示使用的 CPU 不受 cgroup 限制。 cfs_quota_us 的最小值为1"ms"(1000us)，最大值为 1s。

因此，我们将 cfs_quota_us设置为 30000us 从理论上讲就可以限制 climits 控制的进程的 cpu 利用率最多是 30%。

    echo 30000 > cpu.cfs_quota_us

```
[root@localhost climits]# cd /sys/fs/cgroup/cpu/climits
[root@localhost climits]# cat cpu.cfs_quota_us
-1
[root@localhost climits]# echo 30000 > cpu.cfs_quota_us
[root@localhost climits]# cat cpu.cfs_quota_us
30000
```

### 新建一个窗口，并使用 stress 模拟一个任务， 并查看该cpu使用率

    stress -c 1

```
[root@localhost climits]# stress -c 1
stress: info: [447874] dispatching hogs: 1 cpu, 0 io, 0 vm, 0 hdd
```

打开另一个窗口终端监控CPU

    pidstat -C stress -p ALL 1 10000

```
[root@localhost ~]# pidstat -C stress -p ALL 1 10000
Linux 4.18.0-348.el8.x86_64 (localhost.localdomain)     11/28/2023      _x86_64_        (8 CPU)

04:17:55 PM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
04:17:56 PM     0    447874    0.00    0.00    0.00    0.00    0.00     3  stress
04:17:56 PM     0    447875  100.00    0.00    0.00    0.00  100.00     5  stress

04:17:56 PM   UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
04:17:57 PM     0    447874    0.00    0.00    0.00    0.00    0.00     3  stress
04:17:57 PM     0    447875   98.02    0.00    0.00    0.00   98.02     5  stress
```

另启与一个窗口，并且将进程 id 移动到我们的 cgroup 策略(climits):

    /sys/fs/cgroup/cpu/climits/tasks
    echo pro_pid >> tasks

```
[root@localhost climits]# pwd
/sys/fs/cgroup/cpu/climits
[root@localhost climits]# ps -ef|grep stress
root      447874  376083  0 16:15 pts/5    00:00:00 stress -c 1
root      447875  447874 99 16:15 pts/5    00:04:35 stress -c 1
[root@localhost climits]# cat tasks 
[root@localhost climits]# echo 447875 >> tasks
[root@localhost climits]# cat tasks 
447875
```

```
Average:      UID       PID    %usr %system  %guest   %wait    %CPU   CPU  Command
Average:        0    447874    0.00    0.00    0.00    0.00    0.00     -  stress
Average:        0    447875   30.00    0.00    0.00   69.88   30.00     -  stress
```
监控的 cpu 的使用率由 100%降低为 30%左右。

## 触发控制
当物理内存达到上限后，系统的默认行为是kill掉cgroup中继续申请内存的进程，那么怎么控制这样的行为呢？答案是配置memory.oom_control

这个文件里面包含了一个控制是否为当前cgroup启动OOM-killer的标识。如果写0到这个文件，将启动OOM-killer，当内核无法给进程分配足够的内存时，将会直接kill掉该进程；如果写1到这个文件，表示不启动OOM-killer，当内核无法给进程分配足够的内存时，将会暂停该进程直到有空余的内存之后再继续运行；同时，memory.oom_control还包含一个只读的under_oom字段，用来表示当前是否已经进入oom状态，也即是否有进程被暂停了。

    注意：root cgroup的oom killer是不能被禁用的

```
[root@localhost climits]# cat memory.oom_control
oom_kill_disable 0
under_oom 0
oom_kill 5
```

## 实操三:使用 cgroups 对 cpu 进行控制（/sys/fs/cgroup/cpuset）

### 创建内存的 cgroup，很简单我们进入到 cgroup 的内存控制目录“/sys/fs/cgroup/cpuset”，并且继续执行像在memory目录之中的动作，创建climits目录。

```
[root@localhost cpuset]# cd /sys/fs/cgroup/cpuset
[root@localhost cpuset]# mkdir climits
[root@localhost cpuset]# ls climits/
cgroup.clone_children  cpuset.cpu_exclusive  cpuset.effective_cpus  cpuset.mem_exclusive  cpuset.memory_migrate   cpuset.memory_spread_page  cpuset.mems                cpuset.sched_relax_domain_level  tasks
cgroup.procs           cpuset.cpus           cpuset.effective_mems  cpuset.mem_hardwall   cpuset.memory_pressure  cpuset.memory_spread_slab  cpuset.sched_load_balance  notify_on_release
```

## 小结:

至此我们成功的模拟了对 "内存" 和 "cpu" 的使用控制,而 docker 本质也是调用这些的 API来完成对资源的管理，只不过 docker 的易用性和镜像的设计更加人性化，所以 docker才能风靡全球。

## 常用可调参数

### tasks

列举绑定到某个cgroup的所有进程ID（PID）

### cgroup.procs

列举一个cgroup节点下的所有线程组ID

### notify_on_release

填 0或1，表示是否在cgroup中最后一个task退出时通知运行release agent，默认情况下是 0，表示不运行

### release_agent

指定 release agent执行脚本的文件路径（该文件在最顶层cgroup目录中存在），在这个脚本通常用于自动化umount无用的cgroup
