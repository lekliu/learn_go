package rpc

import (
	"net"
	"reflect"
)

// 声明服务端
type Client struct {
	conn net.Conn
}

// 构造方法
func NewClient(conn net.Conn) *Client {
	return &Client{conn: conn}
}

// 实现通用的 RPC 客户端
// 传入访问的函数名
// fPtr指向的是函数原型
// var select fun  xx(User)
// cli.callRPC("selectUser", &select)
func (c *Client) callRPC(rpcName string, fPtr interface{}) {
	// 通过反射，获取fPtr未初始化的函数原弄
	fn := reflect.ValueOf(fPtr).Elem()
	// 需要另一个函数，作用是对第一个函数参数操作
	f := func(args []reflect.Value) []reflect.Value {
		// 处理参数
		inArgs := make([]interface{}, len(args))
		for _, arg := range args {
			inArgs = append(inArgs, arg.Interface())
		}
		//连接

		//编码数据
		reqRPC := RPCData{Name: rpcName, Args: inArgs}
		b, err := encode(reqRPC)
		if err != nil {
			panic(err)
		}
		// 写数据
		err = cliSession.Write(b)
		if err != nil {
			panic(err)
		}
		//服务端发过来返回值，此时应该读取和解拆
		respBytes, err := cliSession.Read()
		if err != nil {
			panic(err)
		}
		//解码
		respRPC, err := decode(respBytes)
		if err != nil {
			panic(err)
		}
		//处理服务端返回的数据
		outArgs := make([]reflect.Value, 0, len(respRPC.Args))
		for i, arg := range respRPC.Args {
			// 必须进行 nil 转换
			if arg == nil {
				//reflect. Zero()会返回类型的零值的 value
				//.out()会返回函数输出的参数类型
				outArgs = append(outArgs, reflect.Zero(fn.Type().Out(i)))
				continue
			}
			outArgs = append(outArgs, reflect.ValueOf(arg))
		}
		return outArgs
	}

	//完成原型到函数调用的内部转换
	//参数 1 是 reflect.Type
	//参数 2 f 是函数类型，是对于参数 1 fn 函数的操作
	//fn 是定义，f 是具体操作
	v := reflect.MakeFunc(fn.Type(), f)
	// 为函数 fPtr赋值，过程
	fn.Set(v)
}
