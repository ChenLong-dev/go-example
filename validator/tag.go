//@Title:		tag.go
//@Description:	解析tag的函数实现
package validator

import (
	"fmt"
	"reflect"

	"time"
)

var (
	timeType = reflect.TypeOf(time.Time{})
)

// toStructType ...
func toStructType(typ reflect.Type) (reflect.Type, bool) {
	if typ == nil {
		return nil, false
	}

	//如果是切片/数组则判断其的元素、如果是map则判断其value
	switch typ.Kind() {
	case reflect.Slice, reflect.Array, reflect.Map:
		typ = typ.Elem()
	}

	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct || typ == timeType {
		return nil, false
	}

	return typ, true
}

// getTag ...
func getTag(typ reflect.Type, fields []string, key string) (tag string, err error) {
	typ, ok := toStructType(typ)
	if !ok || len(fields) == 0 {
		err = fmt.Errorf("param err, tag: [%v], fields: [%v]", typ, fields)
		return
	}

	field := fields[0]
	fields = fields[1:]
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		//匿名字段、unexport字段均不处理
		if !f.Anonymous && len(f.PkgPath) > 0 {
			continue
		}

		if f.Name != field {
			continue
		}

		if len(fields) > 0 {
			return getTag(f.Type, fields, key)
		}

		return f.Tag.Get(key), nil
	}

	err = fmt.Errorf("param err, tag: [%v], fields: [%v]", typ, fields)
	return
}
