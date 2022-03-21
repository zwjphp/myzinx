package znet

import (
	"fmt"
	"myzinx/zinx/utils"
	"myzinx/zinx/ziface"
	"net"
)

type Connection struct {
	// 当前链接的socket TCP套接字
	Conn *net.TCPConn
	// 当前链接的ID 也可以称作为SessionID, ID 全局唯一
	ConnID uint32
	// 当前链接的关闭状态
	isClosed bool
    // 该连接的处理方法router
    Router ziface.IRouter

	// 告知该链接已经退出/停止的channel
	ExitBuffChan chan bool
}

// 创建链接的方法
func NewConntion(conn *net.TCPConn, connID uint32, router ziface.IRouter) *Connection{
	c := &Connection{
		Conn:      conn,
		ConnID:    connID,
		isClosed:  false,
		Router: router,
	}

	return c
}

// 处理conn读取数据的Goroutine
func (c *Connection) StartReader() {
	fmt.Println("Reader Goroutine is running")
	defer fmt.Println(c.RemoteAddr().String(), " conn reader exit!")
	defer c.Stop()

	for {
		// 读取最大的数据到buf中
		buf := make([]byte, utils.GlobalObject.MaxPacketSize)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv buf err ", err)
			c.ExitBuffChan <- true
			continue
		}

		// 得到当前客户端请求的Request数据
		req := Request{
			conn:c,
			data:buf,
		}

		// 从路由Routers 中找到注册绑定Conn的对应Handle
		go func (request ziface.IRequest) {
			// 执行注册的路由方法
			c.Router.PostHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
	}
}


// 启动链接，让当前链接开始工作
func (c *Connection) Start() {
	// 开启处理该链接读取到客户端数据之后的请求业务
	go c.StartReader()

	for {
		select {
		case <- c.ExitBuffChan:
			// 得到推出消息， 不在阻塞
			return
		}
	}
}
// 停止链接，结束当前链接状态
func (c *Connection) Stop() {
	// 1. 如果当前链接已经关闭
	if c.isClosed == true {
		return
	}
	c.isClosed = true

	//TODO Connection Stop() 如果用户注册了该链接的关闭回调业务，那么在此刻应该显示调用

	// 关闭socket链接
	c.Conn.Close()

	// 通知从缓冲队列读取数据的业务， 该链接已经关闭
	c.ExitBuffChan <- true

	// 关闭该链接全部管道
	close(c.ExitBuffChan)
}
// 从当前连接获取原始的socket TCPConn
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}
// 获取当前连接ID
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}
// 获取远程客户端地址信息
func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}


