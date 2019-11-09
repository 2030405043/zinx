package iznet

import "net"

type IConnection interface {
	// 启动
	Start()
	// 停止
	Stop()

	GetTCPConnection() *net.TCPConn

	GetConnId() uint

	RemoteAddr() net.Addr

	Send([]byte) error
}

//
type HandleFunc func(*net.TCPConn, []byte, int) error
