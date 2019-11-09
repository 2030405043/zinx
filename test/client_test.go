package test

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestClient(t *testing.T) {
	// TCP 的客户端
	// 1.
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("Dial  err:", err)
		return
	}

	for {
		_, err = conn.Write([]byte("Hello..Server.."))
		if err != nil {
			fmt.Println("Write  err:", err)
			return
		}
		buf := make([]byte, 1024)
		_, err = conn.Read(buf)
		if err != nil {
			fmt.Println("Write  err:", err)
			return
		}
		fmt.Println("client 接收到的数据:", string(buf), "")
		time.Sleep(time.Second * 1)
	}
}
