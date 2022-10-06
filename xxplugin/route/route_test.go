package router

import (
	"context"
	"encoding/json"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

type Handler1 struct {
}

type Arg1 struct {
	A int
}

type Resp1 struct {
	B int
}

func (h *Handler1) Func1(ctx context.Context, arg *Arg1, resp *Resp1) error {
	logx.Infof("[Func1] arg:%v", arg)
	defer func() {
		logx.Infof("[Func1] resp:%v", resp)
	}()
	resp.B = arg.A + 500
	return nil
}

func (h *Handler1) Func2(ctx context.Context, arg *Arg1, resp *Resp1) error {
	logx.Infof("[Func2] arg:%v", arg)
	defer func() {
		logx.Infof("[Func2] resp:%v", resp)
	}()
	resp.B = arg.A
	return nil
}

func (h *Handler1) Func3(ctx context.Context, arg Arg1, resp *[]string) error {
	defer func() {
		logx.Infof("[Func3] resp:%v", resp)
	}()
	*resp = append(*resp, "xxx")
	*resp = append(*resp, "yyy")
	*resp = append(*resp, "zzz")
	return nil
}

func (h *Handler1) Func4(ctx context.Context, arg Arg1, resp *map[string]string) error {
	defer func() {
		logx.Infof("[Func4] resp:%v", resp)
	}()
	(*resp)["xxx"] = "xxx"
	(*resp)["yyy"] = "yyy"
	(*resp)["zzz"] = "zzz"
	return nil
}

func TestRouter_Call1(t *testing.T) {
	r := New()
	//注册
	err := r.RegisterName("Handler1", &Handler1{})
	if err != nil {
		t.Error(err)
		return
	}
	arg, err := json.Marshal(&Arg1{A: 10000})
	if err != nil {
		t.Error(err)
		return
	}
	// 参数解析函数
	dec := func(v interface{}) error {
		return json.Unmarshal(arg, v)
	}
	logx.Infof("arg:%v", arg)
	res, err := r.Call(context.TODO(), "Handler1", "Func1", dec)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}

func TestRouter_Call(t *testing.T) {
	r := New()
	Convey("注册函数到Router中", t, func() {
		err := r.RegisterName("Handler1", &Handler1{})
		So(err, ShouldBeNil)
		arg, err := json.Marshal(&Arg1{A: 10000})
		So(err, ShouldBeNil)
		dec := func(v interface{}) error {
			return json.Unmarshal(arg, v)
		}
		Convey("注册后调用", func() {
			Convey("调用入参为指针的情况\n", func() {
				res, err1 := r.Call(context.TODO(), "Handler1", "Func1", dec)
				So(err1, ShouldBeNil)
				t.Log("res:", res)
				resp, ok := res.(*Resp1)
				So(ok, ShouldBeTrue)
				So(resp.B, ShouldEqual, 10500)
			})
			Convey("调用入参不为指针的情况\n", func() {
				res, err := r.Call(context.TODO(), "Handler1", "Func2", dec)
				So(err, ShouldBeNil)

				resp, ok := res.(*Resp1)
				So(ok, ShouldBeTrue)
				So(resp.B, ShouldEqual, 10000)
			})
			Convey("参数非法的情况\n", func() {
				_, err = r.Call(context.TODO(), "Handler1", "Func2", func(interface{}) error {
					return fmt.Errorf("param is error!")
				})
				t.Log(err)
				So(err, ShouldBeError)
			})

			Convey("调用未注册的函数\n", func() {
				_, err = r.Call(context.TODO(), "Handler1", "Func22", dec)
				t.Log(err)
				So(err, ShouldBeError)
			})
			Convey("注销后调用\n", func() {
				r.UnRegister("Handler1")
				_, err = r.Call(context.TODO(), "Handler1", "Func1", dec)
				t.Log(err)
				So(err, ShouldBeError)
			})
			Convey("出参为切片\n", func() {
				res, err := r.Call(context.TODO(), "Handler1", "Func3", dec)
				So(err, ShouldBeNil)
				t.Log(res)

				resp, ok := res.(*[]string)
				So(ok, ShouldBeTrue)

				So((*resp)[0], ShouldEqual, "xxx")
				So((*resp)[1], ShouldEqual, "yyy")
				So((*resp)[2], ShouldEqual, "zzz")
			})
			Convey("出参为map\n", func() {
				res, err := r.Call(context.TODO(), "Handler1", "Func4", dec)
				So(err, ShouldBeNil)
				t.Log(res)

				resp, ok := res.(*map[string]string)
				So(ok, ShouldBeTrue)

				So((*resp)["xxx"], ShouldEqual, "xxx")
				So((*resp)["yyy"], ShouldEqual, "yyy")
				So((*resp)["zzz"], ShouldEqual, "zzz")
			})
		})
	})
}

type Handler2 struct {
}

type handler2 struct {
}

type Arg2 struct {
	A int
}
type Resp2 struct {
	A int
}

// Func1 ...
func (h *handler2) Func1(ctx context.Context, arg *Arg2, resp *Resp2) error {
	resp.A = arg.A
	return nil
}

type Handler3 struct {
}

type Arg3 struct {
	A int
}

type arg3 struct {
	A int
}

type Resp3 struct {
	A int
}
type resp3 struct {
	A int
}

// Func1 ...
func (h *Handler3) Func1(ctx context.Context, arg *Arg3) error {
	return nil
}

// Func2 ...
func (h *Handler3) Func2(a int, arg *Arg3, resp *Resp3) error {
	return nil
}

// Func3 ...
func (h *Handler3) Func3(ctx context.Context, arg *arg3, resp *Resp3) error {
	return nil
}

// Func4 ...
func (h *Handler3) Func4(ctx context.Context, arg *Arg3, resp Resp3) error {
	return nil
}

// Func5 ...
func (h *Handler3) Func5(ctx context.Context, arg *Arg3, resp *resp3) error {
	return nil
}

// Func6 ...
func (h *Handler3) Func6(ctx context.Context, arg *Arg3, resp *Resp3) (error, int) {
	return nil, 0
}

// Func7 ...
func (h *Handler3) Func7(ctx context.Context, arg *Arg3, resp *Resp3) int {
	return 1
}

// Func ...
func (h *Handler3) Func(ctx context.Context, arg *Arg3, resp *Resp3) error {
	resp.A = arg.A
	return nil
}

// TestRouter_RegisterName ...
func TestRouter_RegisterName(t *testing.T) {
	r := New()

	Convey("注册无成员函数的结构体", t, func() {
		err := r.RegisterName("Handler1", &Handler2{})
		So(err, ShouldBeError)
	})
	Convey("注册内部结构体", t, func() {
		err := r.RegisterName("Handler1", &handler2{})
		So(err, ShouldBeError)
	})

	Convey("参数异常情况", t, func() {
		err := r.RegisterName("Handler3", &Handler3{})
		So(err, ShouldBeNil)

		arg, err := json.Marshal(&Arg1{A: 1})
		dec := func(v interface{}) error {
			return json.Unmarshal(arg, v)
		}
		So(err, ShouldBeNil)

		_, err = r.Call(context.TODO(), "Handler3", "Func1", dec)
		So(err, ShouldBeError)

		_, err = r.Call(context.TODO(), "Handler3", "Func2", dec)
		So(err, ShouldBeError)

		_, err = r.Call(context.TODO(), "Handler3", "Func3", dec)
		So(err, ShouldBeError)

		_, err = r.Call(context.TODO(), "Handler3", "Func4", dec)
		So(err, ShouldBeError)

		_, err = r.Call(context.TODO(), "Handler3", "Func5", dec)
		So(err, ShouldBeError)

		_, err = r.Call(context.TODO(), "Handler3", "Func6", dec)
		So(err, ShouldBeError)

		_, err = r.Call(context.TODO(), "Handler3", "Func7", dec)
		So(err, ShouldBeError)

		res, err := r.Call(context.TODO(), "Handler3", "Func", dec)
		So(err, ShouldBeNil)

		resp, ok := res.(*Resp3)
		So(ok, ShouldBeTrue)
		So(resp.A, ShouldEqual, 1)
	})
}
