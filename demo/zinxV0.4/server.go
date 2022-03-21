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
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
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
