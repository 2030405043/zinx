package znet

import (
	"zinx/iznet"
)

// 将链接 和 数据绑定在一起
type Request struct {
	Conn iznet.IConnection
	Data []byte
}

func (r *Request) GetConnection() iznet.IConnection {
	return r.Conn
}

func (r *Request) GetData() []byte {
	return  r.Data
}

func NewRequest(connection iznet.IConnection,data []byte)iznet.IRequest  {
	request :=& Request{
		Conn: connection,
		Data: data,
	}
return  request
}