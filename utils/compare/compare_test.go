package compare

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestCompareTwoMapInterface(t *testing.T) {
	Convey("测试 CompareTwoMapInterface 功能\n", t, func() {
		Convey("测试 CompareTwoMapInterface 判断两个map是否相等\n", func() {
			a := map[string]interface{}{"a": "b", "c": "d"}
			b := map[string]interface{}{"c": "d", "a": "b"}
			ok := CompareTwoMapInterface(a, b)
			So(ok, ShouldBeTrue)
			c := map[string]interface{}{"a": "b", "c": "d", "e": "f"}
			ok = CompareTwoMapInterface(a, c)
			So(ok, ShouldBeFalse)
		})
		Convey("测试 CompareTwoMapInterface 判断两个多级嵌套map是否相等\n", func() {
			a := map[string]interface{}{
				"key1": "value1",
				"key2": map[string]interface{}{
					"key3": "value3",
					"key4": map[string]interface{}{
						"key5": "value5",
						"key6": map[string]interface{}{
							"key7": "value7",
							"key8": []interface{}{1, "hello", map[string]interface{}{"key9": "value9"}},
						},
					},
				},
			}
			b := map[string]interface{}{
				"key1": "value1",
				"key2": map[string]interface{}{
					"key3": "value3",
					"key4": map[string]interface{}{
						"key5": "value5",
						"key6": map[string]interface{}{
							"key7": "value7",
							"key8": []interface{}{1, "hello", map[string]interface{}{"key9": "value9"}},
						},
					},
				},
			}
			ok := CompareTwoMapInterface(a, b)
			So(ok, ShouldBeTrue)
			c := map[string]interface{}{
				"key1": "value1",
				"key2": map[string]interface{}{
					"key3": "value3",
					"key4": map[string]interface{}{
						"key5": "value5",
						"key6": map[string]interface{}{
							"key7": "value7",
							"key8": []interface{}{1, "Hello", map[string]interface{}{"key9": "value9"}},
						},
					},
				},
			}
			ok = CompareTwoMapInterface(a, c)
			So(ok, ShouldBeFalse)
		})
	})
}
