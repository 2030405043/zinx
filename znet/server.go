package znet

import (
	"errors"
	"fmt"
	"net"
	"zinx/iznet"
)

/*
	基础的Server模块
*/

type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
	Router    iznet.IRouter
}

// 简单 初始化Server模块
func NewServer(name string) iznet.IServer {
	server := &Server{
		Name:      name,
		IP:        "0.0.0.0",
		Port:      8088,
		IPVersion: "tcp4",
		Router:nil,
	}
	return server
}
func (s *Server) Start() {

	// 常规Go TCP服务端开发
	// 1. 获取一个tcp Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("ResolveTCPAddr err:", err)
		return
	}

	// 2. 监听
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("ListenTCP err:", err)
		return
	}

	var connId uint = 0
	// 3.
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("AcceptTCP err:", err)
			continue
		}
		connId++
		//connection := NewConnection2(conn, connId, CallBackToClient)
		connection := NewConnection(conn,connId,s.Router)
		connection.Start()
	}

}

// 写死的处理客户端请求的api
// 往 客户端 写数据
func CallBackToClient(conn *net.TCPConn, buf []byte, len int) error {
	_, err := conn.Write(buf[:len])
	if err != nil {
		fmt.Println("write eee", err.Error())
		return errors.New("CallBackToClient err:" + err.Error())
	}
	return nil
}

func (s *Server) Start2() {
	// 常规Go TCP服务端开发
	// 1. 获取一个tcp Addr
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("ResolveTCPAddr err:", err)
		return
	}
	// 2. 监听
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("ListenTCP err:", err)
		return
	}
	//异步
	go func() {
		// 3. 阻塞等待客户端请求 处理请求
		for {
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("ListenTCP err:", err)
				continue
			} //
			go func() {
				for {
					// 读取 客户端请求数据
					buf := make([]byte, 1024)
					l, err := conn.Read(buf)
					if err != nil {
						fmt.Println("Read buf err:", err)
						continue
					}

					fmt.Println("server 收到的数据: ", string(buf))
					// 回写客户端请求的数据
					if _, err := conn.Write(buf[:l]); err != nil {
						fmt.Println("Write buf err:", err)
						continue
					}
				}
			}()
		}
	}()

}

func (s *Server) Stop() {
	//关闭连接
}

func (s *Server) Serve() {
	s.Start()
	// TODO 启动服务器之后的其他操作
	// 阻塞
	select {}
}

func (s *Server) AddRouter(router iznet.IRouter) {
	s.Router = router
	fmt.Println("success AddRouter..")
}
