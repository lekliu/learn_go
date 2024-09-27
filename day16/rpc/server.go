package rpc

import "reflect"

// 声明服务端
type Server struct {
	//地址
	addr string
	// map 用于维护关系的
	funcs map[string]reflect.Value

	//TODO
}
