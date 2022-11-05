//@Title:		rule.go
//@Description:		给validator提供的扩展校验规则
package validator

import (
	"reflect"
	"regexp"
	"strconv"
	"time"
)

const (
	//正则规则
	rule_regex int = iota
	//函数规则
	rule_func

	//整数最大值和最小值
	MaxValidFloat = 9007199254740992  //2^53
	MinValidFloat = -9007199254740992 //-(2^53)
)

type rule struct {
	typ   int
	regex *regexp.Regexp
	fun   func(val interface{}, param string, top interface{}) bool //第一个参数是字段值，第二个参数是校验tag中指定的value，第三个参数是整个结构体ide值
}

var (
	emailRegex       = regexp.MustCompile(`^[\w_.-]+@[\w-]+(\.[\w-]+)*\.[\w]{2,6}$`)
	adAttributeRegex = regexp.MustCompile(`^[\w-]{1,95}$`)
)
var buildin_rules = map[string]rule{
	"validFloat": {
		typ: rule_func,
		fun: func(val interface{}, param string, top interface{}) bool {
			switch v := val.(type) {
			case uint64:
				return v <= uint64(MaxValidFloat)
			case int64:
				return v >= MinValidFloat && v <= MaxValidFloat
			case float64:
				return v >= MinValidFloat && v <= MaxValidFloat
			case int, uint, int8, int16, int32, uint8, uint16, uint32:
				return true
			default:
				return false
			}
		},
	},
	"orgname": {
		typ:   rule_regex,
		regex: regexp.MustCompile(`^[^/\r\n]{1,95}$`),
	},
	"ud-email": {
		typ: rule_func,
		fun: func(val interface{}, param string, top interface{}) bool {
			email, ok := toString(val)
			if !ok {
				return false
			}

			if email == "" {
				return true
			}

			if len(email) < 6 || len(email) > 95 {
				return false
			}

			return emailRegex.MatchString(email)
		},
	},
	"bmax": {
		typ: rule_func,
		fun: func(val interface{}, param string, top interface{}) bool {
			s, ok := toString(val)
			if !ok {
				return false
			}

			n, err := strconv.Atoi(param)
			if err != nil {
				return false
			}

			return len(s) <= n
		},
	},
	"bmin": {
		typ: rule_func,
		fun: func(val interface{}, param string, top interface{}) bool {
			s, ok := toString(val)
			if !ok {
				return false
			}

			n, err := strconv.Atoi(param)
			if err != nil {
				return false
			}

			return len(s) >= n
		},
	},
	"timestr": {
		typ: rule_func,
		fun: func(val interface{}, param string, top interface{}) bool {
			s, ok := toString(val)
			if !ok {
				return false
			}

			_, err := time.ParseDuration(s)
			return err == nil
		},
	},
	"adAttribute": {
		typ:   rule_regex,
		regex: adAttributeRegex,
	},
}

// 将string或者string指针的接口转换成string类型
func toString(val interface{}) (string, bool) {
	rval := reflect.ValueOf(val)
	if rval.Kind() == reflect.Ptr {
		rval = rval.Elem()
	}

	if rval.Kind() == reflect.String {
		return rval.String(), true
	} else {
		return "", false
	}
}
