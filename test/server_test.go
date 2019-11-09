package test

import (
	"fmt"
	"testing"
	"zinx/iznet"
	"zinx/znet"
)

func TestServer(t *testing.T) {
	// 创建一个server的句柄
	s := znet.NewServer("zinx")
	r := &ZRouter{}
	s.AddRouter(r)
	//运行server
	s.Serve()
}

type ZRouter struct {
	znet.Router
}

func (z *ZRouter) PreHandle(req iznet.IRequest) {
	data := req.GetData()
	fmt.Println("server pre 接收到来自client的数据：",string(data))
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("pre ping...\n"))
	if err!=nil {
		fmt.Println("PreHandle Write err:",err)
	}
}
func (z *ZRouter) PostHandel(req iznet.IRequest) {
	data := req.GetData()
	fmt.Println("server pre 接收到来自client的数据：",string(data))
	_, err := req.GetConnection().GetTCPConnection().Write([]byte("post ping...\n"))
	if err!=nil {
		fmt.Println("PostHandel Write err:",err)
	}
}
func (z *ZRouter) Handel(req iznet.IRequest) {
	data := req.GetData()
	fmt.Println("server 接收到来自client的数据：",string(data))
	_, err:= req.GetConnection().GetTCPConnection().Write([]byte("ping..\n"))
	if err!=nil {
		fmt.Println("Handel Write err:",err)
	}
}