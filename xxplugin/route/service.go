//@Title:		service.go
//@Description:	这部分代码借鉴其他RPC框架实现，将一个interface上的函数注册进来并供后续回调
package router

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"reflect"
	"unicode"
	"unicode/utf8"
)

// Precompute the reflect type for error. Can't use error directly
// because Typeof takes an empty interface value. This is annoying.
var typeOfError = reflect.TypeOf((*error)(nil)).Elem()

// Precompute the reflect type for context.
var typeOfContext = reflect.TypeOf((*context.Context)(nil)).Elem()

type methodType struct {
	Method    reflect.Method
	ArgType   reflect.Type
	ReplyType reflect.Type
}

type service struct {
	name   string                 // name of srv
	rcvr   reflect.Value          // receiver of methods for the srv
	typ    reflect.Type           // type of the receiver
	method map[string]*methodType // registered methods
}

// Val ...
func (s *service) Val() interface{} {
	return s.rcvr.Interface()
}

// preProc ...
func (s *service) preProc(methodName string, dec func(interface{}) error) (argv, replyv, function reflect.Value, err error) {
	mtype := s.method[methodName]
	if mtype == nil {
		err = ErrNoSuchFunction
		return
	}

	// Decode the argument value.
	argIsValue := false
	if mtype.ArgType.Kind() == reflect.Ptr {
		argv = reflect.New(mtype.ArgType.Elem())
	} else {
		argv = reflect.New(mtype.ArgType)
		argIsValue = true
	}

	if err = dec(argv.Interface()); err != nil {
		err = ErrParamUnmarshalFail
		return
	}
	if argIsValue {
		argv = argv.Elem()
	}

	//careate return type value
	replyv = reflect.New(mtype.ReplyType.Elem())
	switch mtype.ReplyType.Elem().Kind() {
	case reflect.Map:
		replyv.Elem().Set(reflect.MakeMap(mtype.ReplyType.Elem()))
	case reflect.Slice:
		replyv.Elem().Set(reflect.MakeSlice(mtype.ReplyType.Elem(), 0, 0))
	}

	//func
	function = mtype.Method.Func

	return
}

// Call ...
func (s *service) Call(ctx context.Context, methodName string, dec func(interface{}) error) (res interface{}, err error) {
	argv, replyv, function, err := s.preProc(methodName, dec)
	if err != nil {
		logx.Error(err)
		return
	}

	// Invoke the Method, providing a new value for the reply.
	returnValues := function.Call([]reflect.Value{s.rcvr, reflect.ValueOf(ctx), argv, replyv})
	// The return value for the Method is an error.
	errInter := returnValues[0].Interface()
	if errInter != nil {
		err = errInter.(error)
		logx.Error(err)
		return
	}

	res = replyv.Interface()
	return
}

// newService ...
func newService(rcvr interface{}) (name string, srv *service, err error) {
	srv = new(service)

	srv.typ = reflect.TypeOf(rcvr)
	srv.rcvr = reflect.ValueOf(rcvr)

	name = reflect.Indirect(srv.rcvr).Type().Name()
	if !isExported(name) {
		err = ErrUnexportedType
		return
	}
	srv.name = name

	srv.method = suitableMethods(srv.typ, true)

	if len(srv.method) == 0 {
		str := ""

		// To help the user, see if a pointer receiver would work.
		method := suitableMethods(reflect.PtrTo(srv.typ), false)
		if len(method) != 0 {
			str = "rpc.Register: type " + name + " has no exported methods of suitable type (hint: pass a pointer to value of that type)"
		} else {
			str = "rpc.Register: type " + name + " has no exported methods of suitable type"
		}

		logx.Info(str)
		err = ErrNoExportedMethod

		return
	}

	return
}

// isExported ...
func isExported(name string) bool {
	rune, _ := utf8.DecodeRuneInString(name)
	return unicode.IsUpper(rune)
}

// isExportedOrBuiltinType ...
func isExportedOrBuiltinType(t reflect.Type) bool {
	for t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	// PkgPath will be non-empty even for an exported type,
	// so we need to check the type name as well.
	return isExported(t.Name()) || t.PkgPath() == ""
}

// suitableMethods returns suitable Rpc methods of typ, it will report
// error using log if reportErr is true.
func suitableMethods(typ reflect.Type, reportErr bool) map[string]*methodType {
	methods := make(map[string]*methodType)
	for m := 0; m < typ.NumMethod(); m++ {
		method := typ.Method(m)
		mtype := method.Type
		mname := method.Name
		// Method must be exported.
		if method.PkgPath != "" {
			continue
		}
		// Method needs four ins: receiver, context.Context, *args, *reply.
		if mtype.NumIn() != 4 {
			if reportErr {
				logx.Info("Method ", mname, " has wrong number of ins:", mtype.NumIn())
			}
			continue
		}
		// First arg must be context.Context
		ctxType := mtype.In(1)
		if !ctxType.Implements(typeOfContext) {
			if reportErr {
				logx.Info("Method ", mname, " must use context.Context as the first parameter")
			}
			continue
		}

		// Second arg need not be a pointer.
		argType := mtype.In(2)
		if !isExportedOrBuiltinType(argType) {
			if reportErr {
				logx.Info(mname, " parameter type not exported: ", argType)
			}
			continue
		}
		// Third arg must be a pointer.
		replyType := mtype.In(3)
		if replyType.Kind() != reflect.Ptr {
			if reportErr {
				logx.Info("Method ", mname, " reply type not a pointer:", replyType)
			}
			continue
		}
		// Reply type must be exported.
		if !isExportedOrBuiltinType(replyType) {
			if reportErr {
				logx.Info("Method ", mname, " reply type not exported:", replyType)
			}
			continue
		}
		// Method needs one out.
		if mtype.NumOut() != 1 {
			if reportErr {
				logx.Info("Method ", mname, " has wrong number of outs:", mtype.NumOut())
			}
			continue
		}
		// The return type of the Method must be error.
		if returnType := mtype.Out(0); returnType != typeOfError {
			if reportErr {
				logx.Info("Method", mname, " returns ", returnType.String(), " not error")
			}
			continue
		}
		methods[mname] = &methodType{Method: method, ArgType: argType, ReplyType: replyType}
	}
	return methods
}
