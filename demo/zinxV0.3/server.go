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

//Test PreHandle
func (this *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router PreHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping ....\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
	}
}
//Test Handle
func (this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping...ping...ping\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
	}
}

//Test PostHandle
func (this *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router PostHandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping .....\n"))
	if err !=nil {
		fmt.Println("call back ping ping ping error")
	}
}


// server 模块的测试函数
func main() {
	// 1 创建一个server 句柄 s
	s := znet.NewServer("[zinx V0.3]")

	s.AddRouter(&PingRouter{})

	// 2 开启服务
	s.Serve()
}
