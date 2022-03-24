package main

import (
	"fmt"
	"myzinx/zinx/ziface"
	"myzinx/zinx/znet"
)

//ping test 自定义路由
type PingRouter struct {
	znet.BaseRouter //一定要先基础BaseRouter
}

//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	//先读取客户端的数据，再回写ping...ping...ping
	fmt.Println("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	//回写数据
	err := request.GetConnection().SendMsg(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}


// server 模块的测试函数
func main() {
	// 1 创建一个server 句柄 s
	s := znet.NewServer()

	s.AddRouter(&PingRouter{})

	// 2 开启服务
	s.Serve()
}
