package reslimits

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/process"
)

var gTestLimitsStop string

func printInfo(perfix string) {
	pid := getPid()
	p, err := process.NewProcess(int32(pid))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	cpuUsage, err := p.CPUPercent()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	memoryInfo, err := p.MemoryInfo()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	rss := memoryInfo.RSS / MB

	fmt.Printf("[%s] ==> pid:%d, rss:%dMB, cores:%d, usage:%0.2f%%, defnum:%d, curnum:%d\n", perfix, pid, rss, runtime.NumCPU(), cpuUsage, GetDefaultNum(), GetCurrentNum())
}

func rmdir(limitType string) (err error) {
	var dir string
	switch limitType {
	case MemoryType:
		dir = getMemoryDir()
	case CpuType:
		dir = getCpuDir()
	case CpuSetType:
		dir = getCpuSetDir()
	default:
		return fmt.Errorf("not supported the type! type:%s", limitType)

	}

	if !checkPathIsExist(dir, false) {
		fmt.Printf("%s is not dir\n", dir)
		return nil
	}
	// rmdir dir
	command := exec.Command("rmdir", dir)
	err = command.Start()
	if err != nil {
		return
	}
	fmt.Printf("rmdir %s is successful!\n", dir)
	return
}

func SetLimitsStop(stop string) {
	fmt.Println("[SetLimitsStop] set: ", stop)
	gTestLimitsStop = stop
}

func TestLimitMemory() {
	fmt.Printf("-> TestLimitMemory is starting! child pid is %d\n", os.Getpid())
	size := 0
	blocks := make([][4 * MB]byte, 0)
	for i := 0; ; i++ {
		time.Sleep(time.Second)
		if gTestLimitsStop == "stop" {
			fmt.Printf("[TestLimitMemory] stop ...\n")
			return
		}
		printInfo(fmt.Sprintf("TestLimitMemory-%d", i+1))

		b := [4 * MB]byte{}
		blocks = append(blocks, b)
		size += len(b)
		// fmt.Printf("[men-%d] --> blocks' len:%d, size:%dMB\n", i+1, len(blocks), size/(1024*1024))
	}
}

func TestLimitCpu() {
	fmt.Printf("-> TestLimitCpu is starting! child pid is %d\n", os.Getpid())
	for {
		if gTestLimitsStop == "stop" {
			return
		}
		// time.Sleep(1 * time.Nanosecond)
	}
}

func TestLimitsGoNum() {
	fmt.Printf("-> TestLimitsGoNum is starting! child pid is %d\n", os.Getpid())
	for {
		for i := 0; i < 100; i++ {
			Add()
			go func() {
				defer Done()
				time.Sleep(500 * time.Millisecond)
				//fmt.Println("go routine num: ", GetCurrentNum())
			}()
		}
		//time.Sleep(5 * time.Second)
		if gTestLimitsStop == "stop" {
			return
		}
	}
}

func TestLimits(limitMem, limitCpu, limitGo int, limitCpuCores string) {
	if limitMem != 0 {
		Add()
		go func() {
			defer Done()
			TestLimitMemory()
		}()
	}

	if limitCpu != 0 || limitCpuCores != "" {
		Add()
		go func() {
			defer Done()
			TestLimitCpu()
		}()
	}

	if limitGo != 0 {
		Add()
		go func() {
			defer Done()
			TestLimitsGoNum()
		}()
	}
}
