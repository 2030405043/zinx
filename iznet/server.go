package iznet

type IServer interface {
	//启动服务
	Start()
	//关闭服务
	Stop()
	//运行服务
	Serve()

	// 添加路由
	AddRouter(router IRouter)
}
