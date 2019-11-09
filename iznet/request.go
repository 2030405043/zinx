package iznet

type IRequest interface {
	// 获取链接
	GetConnection() IConnection
	// 获取请求的消息数据
	GetData() []byte
}
