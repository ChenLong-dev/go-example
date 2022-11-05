package validator

import (
	validator2 "github.com/go-playground/validator/v10"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

// TestValidate ...
func TestValidate(t *testing.T) {
	Convey("结构体指针", t, func() {
		type St struct {
			V string `validate:"required" errcode:"E_UNKNOWN"`
		}

		v := New()
		err := v.Validate(&St{
			V: "",
		})

		So(err, ShouldBeError)
		So(err.Error(), ShouldEqual, "E_UNKNOWN")
	})

	Convey("缓存取值", t, func() {
		v := New()
		//通过覆盖率测试到缓存部分
		Convey("指针取缓存", func() {
			type St struct {
				V string `validate:"required" errcode:"E_UNKNOWN"`
			}

			err := v.Validate(St{
				V: "",
			})
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "E_UNKNOWN")

			//取缓存
			err = v.Validate(&St{
				V: "",
			})
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})
		//通过覆盖率测试到缓存部分
		Convey("取其他字段的缓存数据", func() {
			type St struct {
				V1 string `validate:"required" errcode:"E_UNKNOWN"`
				V2 string `validate:"required" errcode:"E_UNKNOWN"`
			}

			err := v.Validate(&St{
				V1: "",
				V2: "a",
			})
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "E_UNKNOWN")

			//取V2缓存失败
			err = v.Validate(&St{
				V1: "a",
				V2: "",
			})
			So(err, ShouldBeError)
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})
	})

	Convey("特殊validate标签", t, func() {
		type St struct {
			V1 string `validate:"gte=3" errcode:"E_UNKNOWN"`
			V2 string `validate:"eq=a|eq=b" errcode:"E_UNKNOWN"`
		}

		v := New()
		Convey("带参数标签", func() {
			err := v.Validate(St{
				V1: "xx",
				V2: "a",
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})
		Convey("多条件标签", func() {
			err := v.Validate(St{
				V1: "xxx",
				V2: "",
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})
	})
	Convey("多tag的情况", t, func() {
		type St struct {
			V1 string `validate:"required,lte=3" errcode:"E_A,E_B"`
			V2 string `validate:"required,lte=3" errcode:"E_A"`
			V3 string `validate:"omitempty,lte=3,gte=1" errcode:",E_A,E_B"`
		}

		v := New()

		//patches := ApplyFunc(errors.New, func(code string, args ...interface{}) *errors.Error {
		//	return &errors.Error{Code: code}
		//
		//})
		//defer patches.Reset()

		Convey("一对一的情况", func() {
			err := v.Validate(St{
				V1: "",
				V2: "a",
				V3: "a",
			})
			So(err.Error(), ShouldEqual, "E_A")
		})
		Convey("多对一的情况", func() {
			err := v.Validate(St{
				V1: "a",
				V2: "",
				V3: "a",
			})
			So(err.Error(), ShouldEqual, "E_A")
		})
		Convey("多对多的情况", func() {
			err := v.Validate(St{
				V1: "a",
				V2: "a",
				V3: "xxxx",
			})
			So(err.Error(), ShouldEqual, "E_A")
		})
		Convey("多标签错误的情况", func() {
			err := v.Validate(St{
				V1: "",
				V2: "",
				V3: "a",
			})
			So(err, ShouldBeError)
		})
	})
	Convey("特殊结构", t, func() {
		v := New()
		Convey("映射", func() {
			type St struct {
				V map[string]string `validate:"dive,required" errcode:",E_UNKNOWN"`
			}

			err := v.Validate(St{
				V: map[string]string{
					"a": "",
				},
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})

		Convey("数组", func() {
			type St struct {
				V [1]string `validate:"dive,required" errcode:",E_UNKNOWN"`
			}

			err := v.Validate(St{
				V: [1]string{
					"",
				},
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})

		Convey("切片", func() {
			type St struct {
				V []string `validate:"dive,required" errcode:",E_UNKNOWN"`
			}

			err := v.Validate(St{
				V: []string{
					"",
				},
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})
	})
}

// TestValidate_Error ...
func TestValidate_Error(t *testing.T) {
	type Test_ST1 struct {
		V string `validate:"required" errcode:"E_UNKNOWN"`
	}

	Convey("校验通过的情况", t, func() {
		v := New()
		Convey("校验通过的情况", func() {
			err := v.Validate(Test_ST1{
				V: "xx",
			})
			So(err, ShouldBeNil)
		})
	})

	Convey("校验参数不是结构体的情况", t, func() {
		v := New()
		Convey("数字", func() {
			err := v.Validate(1)
			So(err.Error(), ShouldEqual, "validator: (nil int)")
		})

		Convey("字符串", func() {
			err := v.Validate("ss")
			So(err.Error(), ShouldEqual, "validator: (nil string)")
		})
	})
	Convey("tag写错的情况", t, func() {
		type St1 struct {
			V string `validate:"required,gte=10" errcode:"E_UNKNOWN,,"`
		}

		type St2 struct {
			V1 string `validate:"required" errcode:""`
			V2 string `validate:"required,lte=3" errcode:"E_UNKNOWN,"`
		}

		v := New()
		Convey("标签个数不一致", func() {
			So(v.Validate(St1{
				V: "",
			}).Error(), ShouldEqual, validator2.New().Struct(St1{
				V: "",
			}).Error())
		})
		Convey("errcode标签为空字符串", func() {
			So(v.Validate(St2{
				V1: "",
				V2: "a",
			}).Error(), ShouldEqual, validator2.New().Struct(St2{
				V1: "",
				V2: "a",
			}).Error())
		})
		Convey("errcode标签个别为空", func() {
			err := v.Validate(St2{
				V1: "a",
				V2: "aaaa",
			})
			So(err.Error(), ShouldEqual, validator2.New().Struct(St2{
				V1: "a",
				V2: "aaaa",
			}).Error())
		})
	})
}

// TestValidate_CustomTag ...
func TestValidate_CustomTag(t *testing.T) {
	Convey("自定义正则匹配Tag", t, func() {
		type St1 struct {
			V string `validate:"orgname" errcode:"E_UNKNOWN"`
		}

		type St2 struct {
			V int `validate:"orgname" errcode:"E_UNKNOWN"`
		}

		v := New()

		Convey("校验失败的情况", func() {
			err := v.Validate(St1{
				V: "",
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})

		Convey("校验成功的情况", func() {
			err := v.Validate(St1{
				V: "a",
			})
			So(err, ShouldBeNil)
		})

		Convey("字段非字符串类型", func() {
			err := v.Validate(St2{
				V: 1,
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})

	})

	Convey("自定义函数匹配：不带参数", t, func() {
		type St struct {
			V string `validate:"f" errcode:"E_UNKNOWN"`
		}

		buildin_rules["f"] = rule{
			typ: rule_func,
			fun: func(val interface{}, param string, top interface{}) bool {
				v, ok := top.(St)
				if !ok {
					return false
				}

				if v.V != val {
					return false
				}

				return val == "v"
			},
		}

		v := New()
		Convey("校验通过的情况", func() {
			err := v.Validate(St{
				V: "v",
			})
			So(err, ShouldBeNil)
		})

		Convey("校验失败的情况", func() {
			err := v.Validate(St{
				V: "",
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})

		delete(buildin_rules, "f")
	})

	Convey("自定义函数匹配：带参数", t, func() {
		type St struct {
			V string `validate:"f=v" errcode:"E_UNKNOWN"`
		}

		buildin_rules["f"] = rule{
			typ: rule_func,
			fun: func(val interface{}, param string, top interface{}) bool {
				if val == param {
					return true
				}
				return false
			},
		}

		v := New()

		Convey("带参数：校验成功", func() {
			err := v.Validate(St{
				V: "v",
			})
			So(err, ShouldBeNil)
		})

		Convey("带参数：校验失败", func() {
			err := v.Validate(St{
				V: "",
			})
			So(err.Error(), ShouldEqual, "E_UNKNOWN")
		})

		delete(buildin_rules, "f")
	})
}
