//@Title:		validator_test.go
//@Description:
package validator

import (
	. "github.com/smartystreets/goconvey/convey"
	"reflect"
	"testing"
	"time"
)

// TestGetTag ...
func TestGetTag(t *testing.T) {
	Convey("getTag：参数类型错误", t, func() {
		Convey("getTag: nil类型", func() {
			_, err := getTag(nil, []string{"a"}, "key")
			//t.Logf("err:%v", err)
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "param err, tag: [<nil>], fields: [[a]]")
		})
		Convey("getTag: 时间类型", func() {
			_, err := getTag(reflect.TypeOf(time.Time{}), []string{"a"}, "key")
			//t.Logf("err:%v", err)
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "param err, tag: [<nil>], fields: [[a]]")
		})
		Convey("getTag: 非结构体类型", func() {
			_, err := getTag(reflect.TypeOf(1), []string{"a"}, "key")
			//t.Logf("err:%v", err)
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "param err, tag: [<nil>], fields: [[a]]")
		})
	})

	Convey("getTag：field错误", t, func() {
		type St1 struct {
			A int `key:"tag"`
		}
		Convey("getTag: 取出正确tag", func() {
			tag, err := getTag(reflect.TypeOf(&St1{}), []string{"A"}, "key")
			//t.Logf("tag:%s", tag)
			So(err, ShouldBeNil)
			So(tag, ShouldEqual, "tag")
		})
		Convey("getTag: field与结构不符", func() {
			_, err := getTag(reflect.TypeOf(&St1{}), []string{"A", "Z"}, "key")
			//t.Logf("err:%v", err)
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "param err, tag: [<nil>], fields: [[Z]]")
		})
		Convey("getTag: 错误的字段名", func() {
			_, err := getTag(reflect.TypeOf(&St1{}), []string{"Z"}, "key")
			//t.Logf("err:%v", err)
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "param err, tag: [validator.St1], fields: [[]]")
		})
	})

	Convey("getTag：一些异常情况", t, func() {
		type St1 struct {
			A int `key:"tag"`
		}
		type St2 struct {
			A int
		}
		Convey("getTag: 指定tag不存在的情况", func() {
			val, err := getTag(reflect.TypeOf(&St1{}), []string{"A"}, "x")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "")
		})

		Convey("getTag: 没有tag的情况", func() {
			val, err := getTag(reflect.TypeOf(&St2{}), []string{"A"}, "x")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "")
		})
	})

	Convey("getTag：复杂结构", t, func() {
		type St1 struct {
			A int `key:"tag"`
		}
		type St2 struct {
			S1 St1
		}

		type St3 struct {
			St1
		}

		type St4 struct {
			S1s [3]St1
		}

		type St5 struct {
			PS1s [3]*St1
		}

		type St6 struct {
			S1s []St1
		}

		type St7 struct {
			PS1s []*St1
		}

		type St8 struct {
			MS1s map[string]St1
		}
		type St9 struct {
			PMS1s map[string]*St1
		}

		type St10 struct {
			S2 St2
		}
		Convey("getTag: 子结构体", func() {
			val, err := getTag(reflect.TypeOf(St2{}), []string{"S1", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")
		})
		Convey("getTag: 内嵌结构体", func() {
			val, err := getTag(reflect.TypeOf(St3{}), []string{"St1", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")
		})
		Convey("getTag: 结构体数组", func() {
			val, err := getTag(reflect.TypeOf(St4{}), []string{"S1s", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")

			val, err = getTag(reflect.TypeOf(St5{}), []string{"PS1s", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")
		})

		Convey("getTag: 结构体切片", func() {
			val, err := getTag(reflect.TypeOf(St6{}), []string{"S1s", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")

			val, err = getTag(reflect.TypeOf(St7{}), []string{"PS1s", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")
		})
		Convey("getTag: 结构体映射", func() {
			val, err := getTag(reflect.TypeOf(St8{}), []string{"MS1s", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")

			val, err = getTag(reflect.TypeOf(St9{}), []string{"PMS1s", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")
		})
		Convey("getTag: 2层递归结构体", func() {
			val, err := getTag(reflect.TypeOf(St10{}), []string{"S2", "S1", "A"}, "key")
			So(err, ShouldBeNil)
			So(val, ShouldEqual, "tag")
		})
	})
}
