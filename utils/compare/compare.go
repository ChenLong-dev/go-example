package compare

import (
	"github.com/bytedance/sonic"
)

// ToStructBySonic 通过Sonic 源库解析
func ToStructBySonic(src interface{}, dest any) (err error) {
	// sonic
	byteSrc, err := sonic.Marshal(src)
	if err != nil {
		return
	}
	return sonic.Unmarshal(byteSrc, dest)
}

// ToStringBySonic 通过Sonic 源库struct转string
func ToStringBySonic(src interface{}) (str string, err error) {
	// sonic
	var byteSrc []byte
	byteSrc, err = sonic.Marshal(src)
	if err != nil {
		return
	}
	return string(byteSrc), nil
}

// ToMapBySonic 通过Sonic 源库struct转struct
func ToMapBySonic(src []byte, dest any) (err error) {
	// sonic
	return sonic.Unmarshal(src, dest)
}

func CompareTwoInterface(data1, data2 interface{}) bool {
	f := func(data interface{}) string {
		switch data1.(type) {
		case map[string]interface{}:
			return "map"
		case []interface{}:
			return "list"
		case string:
			return "string"
		case uint8, uint16, uint32, uint64, int8, int16, int32, int64, int, uint, uintptr, float32, float64,
			complex64, complex128:
			return "int"
		default:
			return "default"
		}
	}
	vt1 := f(data1)
	vt2 := f(data2)
	if vt1 != vt2 {
		return false
	}
	if vt1 == "map" {
		valueMap1, ok1 := data1.(map[string]interface{})
		if !ok1 {
			return false
		}
		valueMap2, ok2 := data2.(map[string]interface{})
		if !ok2 {
			return false
		}
		if ok := CompareTwoMapInterface(valueMap1, valueMap2); !ok {
			return false
		}
	} else if vt1 == "list" {
		valueList1, ok1 := data1.([]interface{})
		if !ok1 {
			return false
		}
		valueList2, ok2 := data2.([]interface{})
		if !ok2 {
			return false
		}
		if len(valueList1) != len(valueList2) {
			return false
		}
		for i := 0; i < len(valueList1); i++ {
			if ok := CompareTwoInterface(valueList1[i], valueList2[i]); !ok {
				return false
			}
		}
	} else if vt1 == "string" || vt1 == "int" {
		return data1 == data2
	} else {
		dataStr1, err1 := ToStringBySonic(data1)
		if err1 != nil {
			return false
		}
		dataStr2, err2 := ToStringBySonic(data2)
		if err2 != nil {
			return false
		}
		return dataStr1 == dataStr2
	}
	return true
}

func CompareTwoMapInterface(data1, data2 map[string]interface{}) bool {
	if len(data1) != len(data2) {
		return false
	}
	for key1, value1 := range data1 {
		value2, ok := data2[key1]
		if !ok {
			return false
		}
		if ok = CompareTwoInterface(value1, value2); !ok {
			return false
		}
	}
	return true
}
