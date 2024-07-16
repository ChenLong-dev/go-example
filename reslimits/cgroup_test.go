package reslimits

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

// pidstat -r -C reslimits.test -p ALL 1 10000
func Test_CgroupMem(t *testing.T) {
	Convey("测试cgroup资源限制", t, func() {
		Convey("测试cgroup资源限制---内存限制60MB\n", func() {

			limitMem := 60
			limitCpu := 0
			limitGo := 0
			limitCpuCores := ""

			err := NewGoPool(limitGo)
			t.Log(err)
			So(err, ShouldBeNil)

			err = NewCgroup(limitMem, limitCpu, limitCpuCores)
			t.Log(err)
			So(err, ShouldBeNil)

			printInfo("Test_CgroupMem")
			go func() {
				for i := 0; i < 20; i++ {
					printInfo(fmt.Sprintf("Test_CgroupMem-%d", i+1))
					time.Sleep(1 * time.Second)
				}
				SetLimitsStop("stop")
				fmt.Println("recode is over ...")
			}()

			TestLimits(limitMem, limitCpu, limitGo, limitCpuCores)

			for {
				if gTestLimitsStop == "stop" {
					t.Log("exist test memory ....")
					break
				}
				time.Sleep(1 * time.Second)
			}
			err = rmdir(MemoryType)
			So(err, ShouldBeNil)
		})
	})
}

// pidstat -C reslimits.test -p ALL 1 10000
func Test_CgroupCpuUsage(t *testing.T) {
	Convey("测试cgroup资源限制", t, func() {
		Convey("测试cgroup资源限制---cpu限制60%\n", func() {
			limitMem := 0
			limitCpu := 60
			limitGo := 0
			limitCpuCores := ""

			err := NewGoPool(limitGo)
			t.Log(err)
			So(err, ShouldBeNil)

			err = NewCgroup(limitMem, limitCpu, limitCpuCores)
			t.Log(err)
			So(err, ShouldBeNil)

			printInfo("Test_CgroupCpuUsage")
			go func() {
				for i := 0; i < 20; i++ {
					printInfo(fmt.Sprintf("Test_CgroupCpuUsage-%d", i+1))
					time.Sleep(1 * time.Second)
				}
				SetLimitsStop("stop")
				fmt.Println("recode is over ...")
			}()

			TestLimits(limitMem, limitCpu, limitGo, limitCpuCores)

			for {
				if gTestLimitsStop == "stop" {
					t.Log("exist test cpu's usage ....")
					break
				}
				time.Sleep(1 * time.Second)
			}
			err = rmdir(CpuType)
			So(err, ShouldBeNil)

		})
	})
}

func Test_CgroupCpuCores(t *testing.T) {
	Convey("测试cgroup资源限制\n", t, func() {
		Convey("测测试cgroup资源限制---cpu's core 校验 \n", func() {
			err := checkCpuCores([]string{"1"})
			So(err, ShouldBeNil)
			err = checkCpuCores([]string{"0", "1", "2"})
			So(err, ShouldBeNil)
			err = checkCpuCores([]string{"0", "100", "2"})
			So(err, ShouldNotBeNil)
			t.Log(err)
		})
		Convey("测测试cgroup资源限制---测试 cgroup cpu 和 cpuset(多核心设置） 互斥校验 \n", func() {
			limitMem := 0
			limitCpu := 60
			limitCpuCores := "0,1"
			err := NewCgroup(limitMem, limitCpu, limitCpuCores)
			t.Log(err)
			So(err, ShouldNotBeNil)
		})

		Convey("测测试cgroup资源限制---测试 cgroup cpuset 限制(60%, 0) \n", func() {
			limitMem := 0
			limitCpu := 60
			limitGo := 0
			limitCpuCores := "0"

			err := NewGoPool(limitGo)
			t.Log(err)
			So(err, ShouldBeNil)

			err = NewCgroup(limitMem, limitCpu, limitCpuCores)
			So(err, ShouldBeNil)
			printInfo("Test_CgroupCpuCores")
			go func() {
				for i := 0; i < 20; i++ {
					printInfo(fmt.Sprintf("Test_CgroupCpuCores-%d", i+1))
					time.Sleep(1 * time.Second)
				}
				SetLimitsStop("stop")
				fmt.Println("recode is over ...")
			}()

			TestLimits(limitMem, limitCpu, limitGo, limitCpuCores)

			for {
				if gTestLimitsStop == "stop" {
					t.Logf("exist test cpu's core ....")
					break
				}
				time.Sleep(1 * time.Second)
			}

			err = rmdir(CpuSetType)
			So(err, ShouldBeNil)
		})
	})
}
