package lock

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_Lock(t *testing.T) {
	Convey("测试文件锁", t, func() {
		Convey("测试文件锁---互斥\n", func() {
			filePath := "/tmp/lock.pid"
			var f1, f2 *Flock
			var err error

			defer func() {
				f1.Close()
				f2.Close()
			}()

			f1, err = NewFileLock(filePath)
			t.Logf("err:%v\n", err)
			So(err, ShouldBeNil)

			err = f1.Lock()
			t.Logf("err:%v\n", err)
			So(err, ShouldBeNil)

			f2, err = NewFileLock(filePath)
			t.Logf("err:%v\n", err)
			So(err, ShouldBeNil)

			err = f2.Lock()
			t.Logf("err:%v\n", err)
			So(err, ShouldNotBeNil)
		})
		Convey("测试文件锁---文件路径不存在\n", func() {
			filePath := "/tmp/test/lock.pid"

			f1, err := NewFileLock(filePath)
			t.Logf("err:%v\n", err)
			So(err, ShouldNotBeNil)
			So(f1, ShouldBeNil)
		})

	})
	Convey("测试目录锁", t, func() {
		Convey("测试目录锁---互斥\n", func() {
			dirPath, _ := os.Getwd()
			var f1, f2 *Flock
			var err error

			defer func() {
				f1.Close()
				f2.Close()
			}()

			f1, err = NewFileLock(dirPath)
			t.Logf("err:%v\n", err)
			So(err, ShouldBeNil)

			err = f1.Lock()
			t.Logf("err:%v\n", err)
			So(err, ShouldBeNil)

			f2, err = NewFileLock(dirPath)
			t.Logf("err:%v\n", err)
			So(err, ShouldBeNil)

			err = f2.Lock()
			t.Logf("err:%v\n", err)
			So(err, ShouldNotBeNil)
		})

		Convey("测试目录锁---目录路径不存在\n", func() {
			dirPath := "/tmp/test"

			f1, err := NewFileLock(dirPath)
			t.Logf("err:%v\n", err)
			So(err, ShouldNotBeNil)
			So(f1, ShouldBeNil)
		})
	})
}
