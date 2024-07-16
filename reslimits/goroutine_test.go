package reslimits

import (
	"runtime"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_GoPool(t *testing.T) {
	Convey("测试 协程池 goroutine pool\n", t, func() {
		Convey("测试 协程池 goroutine pool---测试 协程池 受限制(10) \n", func() {
			limitGo := 3
			err := NewGoPool(limitGo)
			So(err, ShouldBeNil)
			for i := 0; i < 20; i++ {
				Add()
				go func() {
					defer Done()
					time.Sleep(time.Second)
					t.Log("go routine num: ", GetCurrentNum())
				}()
			}
			time.Sleep(500 * time.Millisecond)
			Wait()
			t.Logf("num:%d, defaultNum:%d, go routine num: %d", GetNum(), GetDefaultNum(), runtime.NumGoroutine())
		})
	})
}
