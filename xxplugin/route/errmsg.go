//@Title:		errmsg.go
//@Description:		函数路由模块的错误码定义
package router

import (
	"fmt"
)

var (
	//这里不用国际化包的错误，因为这个只是内部使用而已
	ErrUnexportedType     = fmt.Errorf("router: the type is unexported")
	ErrNoExportedMethod   = fmt.Errorf("router: the type has no exported methods of suitable type")
	ErrNoSuchFunction     = fmt.Errorf("E_AGENT_UNSUPPORTED_METHOD")
	ErrParamUnmarshalFail = fmt.Errorf("router: the param unmarshal fail")
	ErrNoSuchService      = fmt.Errorf("router: can't find service  for func raw function")
)
