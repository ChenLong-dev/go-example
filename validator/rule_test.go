//单测文件
package validator

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

//TestExcludeSelf ...
func TestBmaxBmin(t *testing.T) {
	type A struct {
		Attr string `validate:"bmin=1,bmax=3"`
	}
	Convey("bmin-1", t, func() {
		a := &A{Attr: "1"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeNil)
	})

	Convey("bmin-nil", t, func() {
		a := &A{Attr: ""}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeError)
	})
	Convey("bmin-张long", t, func() {
		a := &A{Attr: "张long"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeError)
	})
	Convey("bmin-张", t, func() {
		a := &A{Attr: "张"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeNil)
	})

	Convey("bmax-22", t, func() {
		a := &A{Attr: "22"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeNil)
	})

	Convey("bmax-333", t, func() {
		a := &A{Attr: "333"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeNil)
	})
	Convey("bmax-4444", t, func() {
		a := &A{Attr: "4444"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeError)
		So(err.Error(), ShouldEqual, "Key: 'A.Attr' Error:Field validation for 'Attr' failed on the 'bmax' tag")
	})

	Convey("bmin-0", t, func() {
		a := &A{Attr: "0"}
		err := Validate(a)
		//t.Logf("err:%v", err)
		So(err, ShouldBeNil)
	})

}
