package reslimits

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

const (
	groupName           = "impauth"
	cgroupRoot          = "/sys/fs/cgroup"
	memoryLimitFile     = "memory.limit_in_bytes"
	cpuPeriodFile       = "cpu.cfs_period_us"
	cpuLimitFile        = "cpu.cfs_quota_us" // 表示 Cgroup 可以使用的 cpu 的带宽，单位为“us”。cfs_quota_us 为-1，表示使用的 CPU 不受 cgroup 限制。 cfs_quota_us 的最小值为1"ms"(1000us)，最大值为 1s
	cpusetCpusLimitFile = "cpuset.cpus"
	cpusetMemsLimitFile = "cpuset.mems"
	tasksFile           = "tasks"
	MB                  = 1024 * 1024
	MS                  = 1000
	MemoryType          = "memory"
	CpuType             = "cpu"
	CpuSetType          = "cpuset"
)

func getMemoryDir() string {
	// /sys/fs/cgroup/memory/climits
	return filepath.Join(cgroupRoot, MemoryType, groupName)
}

func getMemoryLimitFile() string {
	// /sys/fs/cgroup/memory/climits/memory.limit_in_bytes
	return filepath.Join(cgroupRoot, MemoryType, groupName, memoryLimitFile)
}

func getMemoryTasksFile() string {
	// /sys/fs/cgroup/memory/climits/tasks
	return filepath.Join(cgroupRoot, MemoryType, groupName, tasksFile)
}

func getCpuDir() string {
	// /sys/fs/cgroup/cpu/climits
	return filepath.Join(cgroupRoot, CpuType, groupName)
}

func getCpuPeriodFile() string {
	// /sys/fs/cgroup/cpu/climits/cpu.cfs_period_us
	return filepath.Join(cgroupRoot, CpuType, groupName, cpuPeriodFile)
}

func getCpuLimitFile() string {
	// /sys/fs/cgroup/cpu/climits/cpu.cfs_quota_us
	return filepath.Join(cgroupRoot, CpuType, groupName, cpuLimitFile)
}

func getCpuTasksFile() string {
	// /sys/fs/cgroup/cpu/climits/tasks
	return filepath.Join(cgroupRoot, CpuType, groupName, tasksFile)
}

func getCpuSetDir() string {
	// /sys/fs/cgroup/cpuset/climits
	return filepath.Join(cgroupRoot, CpuSetType, groupName)
}

func getCpuSetCpusLimitFile() string {
	// /sys/fs/cgroup/cpuset/climits/cpuset.cpus
	return filepath.Join(cgroupRoot, CpuSetType, groupName, cpusetCpusLimitFile)
}

func getCpuSetMemsLimitFile() string {
	// /sys/fs/cgroup/cpuset/climits/cpuset.mems
	return filepath.Join(cgroupRoot, CpuSetType, groupName, cpusetMemsLimitFile)
}

func getCpuSetTasksFile() string {
	// /sys/fs/cgroup/cpuset/climits/tasks
	return filepath.Join(cgroupRoot, CpuSetType, groupName, tasksFile)
}

func limitMem(limit, pid int) (err error) {
	if limit <= 0 {
		return fmt.Errorf("param is error! limit:%d", limit)
	}
	if ok := checkPathIsExist(cgroupRoot, false); !ok {
		return fmt.Errorf("dir [%s] is not exist", cgroupRoot)
	}
	if ok := checkPathIsExist(getMemoryDir(), true); !ok {
		err = fmt.Errorf("%s is not exist", getMemoryDir())
		return
	}
	if err = whiteFile(getMemoryLimitFile(), fmt.Sprintf("%d", limit*MB)); err != nil {
		return
	}
	if err = whiteFile(getMemoryTasksFile(), fmt.Sprintf("%d", pid)); err != nil {
		return
	}
	fmt.Printf("memory setting is successful! pid:%d, mem:%d MiB\n", pid, limit)
	return
}

func limitCpu(limit, pid int) (err error) {
	if limit <= 0 || limit > 100 {
		return fmt.Errorf("param is error! limit:%d", limit)
	}
	if ok := checkPathIsExist(cgroupRoot, false); !ok {
		return fmt.Errorf("dir [%s] is not exist", cgroupRoot)
	}
	if ok := checkPathIsExist(getCpuDir(), true); !ok {
		return fmt.Errorf("[%s] is not exist", getCpuDir())
	}
	if err = whiteFile(getCpuPeriodFile(), fmt.Sprintf("%d", 100*MS)); err != nil {
		return
	}
	if err = whiteFile(getCpuLimitFile(), fmt.Sprintf("%d", limit*MS)); err != nil {
		return
	}
	if err = whiteFile(getCpuTasksFile(), fmt.Sprintf("%d", pid)); err != nil {
		return
	}
	fmt.Printf("cpu usage setting is successful! pid:%d, cpu's uage:%d%%\n", pid, limit)
	return
}

func limitCpuCores(limit string, pid int) (err error) {
	if limit == "" {
		err = fmt.Errorf("param is error")
		return
	}
	cpus := strings.Split(limit, ",")
	if len(cpus) == 0 {
		return fmt.Errorf("param is error! limit:%+v", cpus)
	}
	if ok := checkPathIsExist(cgroupRoot, false); !ok {
		return fmt.Errorf("dir [%s] is not exist", cgroupRoot)
	}
	if err = checkCpuCores(cpus); err != nil {
		return
	}
	if ok := checkPathIsExist(getCpuSetDir(), true); !ok {
		return fmt.Errorf("[%s] is not exist", getCpuSetDir())
	}
	if err = whiteFile(getCpuSetCpusLimitFile(), limit); err != nil {
		return
	}
	if err = whiteFile(getCpuSetMemsLimitFile(), "0"); err != nil {
		return
	}
	if err = whiteFile(getCpuSetTasksFile(), fmt.Sprintf("%d", pid)); err != nil {
		return
	}
	fmt.Printf("cpu cores setting is successful! pid:%d, cpuset:%s\n", pid, limit)
	return
}

func checkCpuCores(cpus []string) (err error) {
	cNum := runtime.NumCPU()
	for _, cpu := range cpus {
		c := -1
		c, err = strconv.Atoi(cpu)
		if err != nil {
			return
		}
		if c > cNum {
			return fmt.Errorf("cpuset is over cpu num! cpu:%s, cpunum:%d", cpu, cNum)
		}
	}
	return
}

func whiteFile(path string, value string) (err error) {
	if err = os.WriteFile(path, []byte(value), 0755); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func getPid() int {
	return os.Getpid()
}

func checkPathIsExist(path string, isCreate bool) (ok bool) {
	// 判断目录是否存在
	if _, err := os.Stat(path); os.IsNotExist(err) {
		if isCreate {
			if err = os.MkdirAll(path, 0o755); err != nil {
				fmt.Printf("mkdir all is failed! err:%v", err)
				return false
			}
			return true
		}
		return false
	}
	return true
}

func NewCgroup2(limit interface{}, lType string) (err error) {
	pid := getPid()
	if pid <= 0 {
		return fmt.Errorf("pid <= 0, pid:%d", pid)
	}
	switch lType {
	case MemoryType:
		return limitMem(limit.(int), pid)
	case CpuType:
		return limitCpu(limit.(int), pid)
	case CpuSetType:
		return limitCpuCores(limit.(string), pid)
	default:
		err = fmt.Errorf("not supported the %s", lType)
		return
	}
}

func NewCgroup(limitMem, limitCpu int, limitCpuCores string) (err error) {
	pid := getPid()
	if pid <= 0 {
		return fmt.Errorf("pid <= 0, pid:%d", pid)
	}
	fmt.Printf("[NewCgroup] pid:%d, limitMem:%d, limitCpu:%d, limitCpuCores:%s\n", pid, limitMem, limitCpu, limitCpuCores)

	if ok := checkPathIsExist(cgroupRoot, false); !ok {
		return fmt.Errorf("path [%s] is not exist", cgroupRoot)
	}

	if limitMem > 0 {
		if ok := checkPathIsExist(getMemoryDir(), true); !ok {
			err = fmt.Errorf("%s is not exist", getMemoryDir())
			return
		}
		if err = whiteFile(getMemoryLimitFile(), fmt.Sprintf("%d", limitMem*MB)); err != nil {
			return
		}
		if err = whiteFile(getMemoryTasksFile(), fmt.Sprintf("%d", pid)); err != nil {
			return
		}
		fmt.Printf("memory setting is successful! pid:%d, mem:%d MiB\n", pid, limitMem)
	}

	// cpu使用率限制和cpu指定多核心不能同时设置， cpu使用率限制和cpu指定单核心能同时设置。limitCpuSet = "1,2,3" 或 limitCpu = 60 和 limitCpuSet = "1"
	cpus := strings.Split(limitCpuCores, ",")
	if limitCpu > 0 && limitCpuCores != "" {
		if len(cpus) > 1 {
			return fmt.Errorf("limitCpu is not 0 and limitCpuSet is multiple, limitCpu:%d, limitCpuCores:%s",
				limitCpu, limitCpuCores)
		}
	}

	if limitCpu > 0 {
		//if err = rmdir(CpuType); err != nil {
		//	return
		//}
		if limitCpu > 100 {
			return fmt.Errorf("limitCpu is greater than 100, limitCpu:%d", limitCpu)
		}
		if ok := checkPathIsExist(getCpuDir(), true); !ok {
			return fmt.Errorf("[%s] is not exist", getCpuDir())
		}
		if err = whiteFile(getCpuPeriodFile(), fmt.Sprintf("%d", 100*MS)); err != nil {
			return
		}
		if err = whiteFile(getCpuLimitFile(), fmt.Sprintf("%d", limitCpu*MS)); err != nil {
			return
		}
		if err = whiteFile(getCpuTasksFile(), fmt.Sprintf("%d", pid)); err != nil {
			return
		}
		fmt.Printf("cpu usage setting is successful! pid:%d, cpu's uage:%d%%\n", pid, limitCpu)
	}

	if len(cpus) > 0 && cpus[0] != "" {
		if err = checkCpuCores(cpus); err != nil {
			return
		}
		//if err = rmdir(CpuSetType); err != nil {
		//	return
		//}
		if ok := checkPathIsExist(getCpuSetDir(), true); !ok {
			return fmt.Errorf("[%s] is not exist", getCpuSetDir())
		}
		if err = whiteFile(getCpuSetCpusLimitFile(), limitCpuCores); err != nil {
			return
		}
		if err = whiteFile(getCpuSetMemsLimitFile(), "0"); err != nil {
			return
		}
		if err = whiteFile(getCpuSetTasksFile(), fmt.Sprintf("%d", pid)); err != nil {
			return
		}
		fmt.Printf("cpu cores setting is successful! pid:%d, cpuset:%s\n", pid, limitCpuCores)
	}
	return
}
