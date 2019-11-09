package znet

import (
	"fmt"
	"net"
	"zinx/iznet"
)

type Connection struct {
	// 连接ID
	ConnId uint
	// 套接字
	Conn *net.TCPConn

	IsClosed bool

	HandleAPI iznet.HandleFunc

	ExitChan chan bool

	Router iznet.IRouter
}

func (c *Connection) Read() {
	fmt.Println("read running...")
	defer fmt.Println("read is exit.. connID : ", c.ConnId)
	defer c.Stop()

	for {
		buf := make([]byte, 1024)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err.Error())
			continue
		}
		fmt.Println("server 接受到：",string(buf))

		req := &Request{
			Conn: c,
			Data: buf,
		}
		// 注册路由方法
		go func(r iznet.IRequest) {
			// 调用路由
			c.Router.PreHandle(r)
			c.Router.Handel(r)
			c.Router.PostHandel(r)
		}(req)
	}

}

func (c *Connection) Start() {
	fmt.Println("start conn.. connID ： ", c.ConnId)
	// 读数据
	// 写数据
	go c.Read()
}

func (c *Connection) Stop() {
	fmt.Println("stop connection.. connID :", c.ConnId)
	//判断是否关闭
	if c.IsClosed == true {
		return
	}
	// 设置状态
	c.IsClosed = true
	//  关闭链接
	c.Conn.Close()
	// 关闭 channel 回收资源
	close(c.ExitChan)
}

func (c *Connection) GetConnId() uint {
	return c.ConnId
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) Send([]byte) error {
	return nil
}

// 初始化
func NewConnection2(conn *net.TCPConn, connId uint, callBackAPI iznet.HandleFunc) iznet.IConnection {
	connection := &Connection{
		ConnId:    connId,
		Conn:      conn,
		IsClosed:  false,
		HandleAPI: callBackAPI,
		ExitChan:  nil,
	}
	return connection
}

func NewConnection(conn *net.TCPConn, connId uint, router iznet.IRouter) iznet.IConnection {
	connection := &Connection{
		ConnId:    connId,
		Conn:      conn,
		IsClosed:  false,
		ExitChan:  nil,
		Router:router,
	}
	return connection
}
