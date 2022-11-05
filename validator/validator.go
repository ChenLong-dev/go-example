//@Title:		validator.go
//@Description:		此模块基于validator的v10版本实现，额外扩展的功能是允许设置errcodetag，在错误时返回该tag的错误信息
package validator

import (
	"fmt"
	rawValidator "github.com/go-playground/validator/v10"
	"reflect"
	"regexp"
	"strings"
	"sync"
)

const (
	VALIDATE_TAG  = "validate"
	ERRORCODE_TAG = "errcode"
)

type Validator interface {
	/*
	 * 校验结构体参数有效性
	 * @val 		结构体或结构体指针
	 * 返回			错误信息，包含两种：参数校验的错误信息（可以直接回复给前端的）、内部错误（代码写的有错误之类的）
	 */
	Validate(val interface{}) error
}

type validator struct {
	validator *rawValidator.Validate
	cache     *cache
}

// parse ...
func (v *validator) parse(typ reflect.Type, fields []string) (m errCodeMap, ok bool) {
	errCodeStr, err := getTag(typ, fields, ERRORCODE_TAG)
	if err != nil || len(errCodeStr) == 0 {
		return
	}
	validateStr, err := getTag(typ, fields, VALIDATE_TAG)
	if err != nil || len(validateStr) == 0 {
		return
	}

	errCodes := strings.Split(errCodeStr, ",")
	validates := strings.Split(validateStr, ",")

	if len(errCodes) != 1 && len(errCodes) != len(validates) {
		return
	}

	//支持 [,,ErrA]对应[ErrA,Erra,Erra]
	for i := len(errCodes) - 1; i >= 0; i-- {
		if errCodes[i] == "" && i < len(errCodes)-1 {
			errCodes[i] = errCodes[i+1]
		}
	}
	m = make(errCodeMap)
	if len(errCodes) == 1 {
		for _, v := range validates {
			m[procValidateTag(v)] = errCodes[0]
		}
	} else {
		for i, v := range validates {
			m[procValidateTag(v)] = errCodes[i]
		}
	}

	return m, true
}

// val应该是一个结构体或者结构体指针，将会校验结构体中字段的限制
func (v *validator) Validate(val interface{}) error {
	err := v.validator.Struct(val)
	if err == nil {
		return nil
	}

	if _, ok := err.(*rawValidator.InvalidValidationError); ok {
		return err
	}

	//经过validator库，这里一定会是结构体类型
	typ := reflect.TypeOf(val)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	//只取第一个错误信息
	fieldError := err.(rawValidator.ValidationErrors)[0]
	fieldStr := fieldError.StructNamespace()
	fields := parseFieldFromValidator(fieldStr)

	//首先从缓存中取值
	if v, ok := v.cache.get(typ, fields, fieldError.ActualTag()); ok {
		if v != "" {
			return fmt.Errorf("%s", v)
		} else {
			return err
		}
	}

	m, ok := v.parse(typ, fields)
	if !ok {
		//如果这里解析错误以后也不可能成功，直接存一个空map
		v.cache.put(typ, fields, make(errCodeMap))
		return err
	}

	v.cache.put(typ, fields, m)
	code, _ := v.cache.get(typ, fields, fieldError.ActualTag())
	if code == "" {
		return err
	}

	return fmt.Errorf("%s", code)
}

type errCode = string
type errCodeMap = map[string]errCode

//缓存结构体，以errCodeMap为单位
type cache struct {
	mutex sync.RWMutex
	data  map[reflect.Type]map[string]errCodeMap //{typ:{"field":{"validate1":"errcode1"...}...}...}
}

/*
 * 从缓存中获取错误码
 * @typ 	结构体类型
 * @fields 	结构体字段路径
 * @key		key
 */
func (c *cache) get(typ reflect.Type, fields []string, key string) (string, bool) {
	field := strings.Join(fields, ",")

	c.mutex.RLock()
	defer c.mutex.RUnlock()

	stMap, ok := c.data[typ]
	if !ok {
		return "", false
	}

	v, ok := stMap[field]
	if !ok {
		return "", false
	}

	//以errCodeMap为单位，errCodeMap中不存在则返回默认错误信息
	code, ok := v[key]
	if !ok {
		return "", true
	}
	return code, true
}

// put ...
func (c *cache) put(typ reflect.Type, fields []string, errs errCodeMap) {
	field := strings.Join(fields, ",")

	c.mutex.Lock()
	defer c.mutex.Unlock()

	_, ok := c.data[typ]
	if !ok {
		c.data[typ] = make(map[string]errCodeMap)
	}

	c.data[typ][field] = errs
}

// newCache ...
func newCache() *cache {
	return &cache{
		data: make(map[reflect.Type]map[string]errCodeMap),
	}
}

// New函数，将会注册rule中定义的规则
func New() Validator {
	ans := &validator{
		validator: rawValidator.New(),
		cache:     newCache(),
	}

	for k := range buildin_rules {
		_ = ans.validator.RegisterValidation(k, func(fl rawValidator.FieldLevel) bool {
			v := buildin_rules[fl.GetTag()]
			if v.typ == rule_regex {
				s, ok := toString(fl.Field().Interface())
				if !ok {
					return false
				}

				return v.regex.MatchString(s)
			}

			return v.fun(fl.Field().Interface(), fl.Param(), fl.Top().Interface())
		})
	}

	return ans
}

var (
	arrRegexp   = regexp.MustCompile(`^([^[]+)\[.+\]$`)
	equalRegexp = regexp.MustCompile(`^([^=]+)=[^=]+$`)
)

// 将validator中的返回值进行解析
func parseFieldFromValidator(fieldStr string) []string {
	fields := strings.Split(fieldStr, ".")

	//排除结构体自身的名字，即第一个元素
	fields = fields[1:]

	//排除切片、map、数组字段后面的[x]部分
	for i, v := range fields {
		res := arrRegexp.FindStringSubmatch(v)
		if len(res) > 0 {
			fields[i] = res[1]
			continue
		}
	}

	return fields
}

// procValidateTag ...
func procValidateTag(vTag string) string {
	//对于a|b这种校验tag，校验结果会返回整个
	if strings.ContainsAny(vTag, "|") {
		return vTag
	}

	//对于a=b这种携带参数的tag，校验结果只包含前面
	v := equalRegexp.FindStringSubmatch(vTag)
	if len(v) > 1 {
		return v[1]
	}

	return vTag
}

var defaultValidator = New()

// Validate ...
func Validate(val interface{}) error {
	return defaultValidator.Validate(val)
}
