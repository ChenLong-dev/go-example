package example

import (
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

func FieldExample() {
	tests := []struct {
		name string
		f    logx.LogField
		want map[string]interface{}
	}{
		{
			name: "error",
			f:    logx.Field("foo", errors.New("bar")),
			want: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "errors",
			f:    logx.Field("foo", []error{errors.New("bar"), errors.New("baz")}),
			want: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
		},
		{
			name: "strings",
			f:    logx.Field("foo", []string{"bar", "baz"}),
			want: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
		},
		{
			name: "duration",
			f:    logx.Field("foo", time.Second),
			want: map[string]interface{}{
				"foo": "1s",
			},
		},
		{
			name: "durations",
			f:    logx.Field("foo", []time.Duration{time.Second, 2 * time.Second}),
			want: map[string]interface{}{
				"foo": []interface{}{"1s", "2s"},
			},
		},
		{
			name: "times",
			f: logx.Field("foo", []time.Time{
				time.Date(2020, time.January, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2020, time.January, 2, 0, 0, 0, 0, time.UTC),
			}),
			want: map[string]interface{}{
				"foo": []interface{}{"2020-01-01 00:00:00 +0000 UTC", "2020-01-02 00:00:00 +0000 UTC"},
			},
		},
		{
			name: "stringer",
			f:    logx.Field("foo", ValStringer{val: "bar"}),
			want: map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			name: "stringers",
			f:    logx.Field("foo", []fmt.Stringer{ValStringer{val: "bar"}, ValStringer{val: "baz"}}),
			want: map[string]interface{}{
				"foo": []interface{}{"bar", "baz"},
			},
		},
	}

	for _, test := range tests {
		key := "FieldExample:" + test.name
		logx.Infof("%v", test.f)
		logx.Infov(test.f)
		logx.Infow(key, test.f)
		logx.Info("==============================")
	}
}
