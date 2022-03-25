package ziface

import "net"

// 定义链接接口
type IConnection interface {
	// 启动链接，让当前链接开始工作
	Start()
	// 停止链接，结束当前链接状态
	Stop()
	// 从当前连接获取原始的socket TCPConn
	GetTCPConnection() *net.TCPConn
	// 获取当前连接ID
	GetConnID() uint32
	// 获取远程客户端地址信息
	RemoteAddr() net.Addr
	//直接将Message数据发送数据给远程的TCP客户端
	SendMsg(msgId uint32, data []byte) error

	// 写消息Goroutine， 用户将数据发送给客户端
	StartWriter()
}

// 定义一个统一处理链接业务的接口
type HandFunc func(*net.TCPConn, []byte, int) error