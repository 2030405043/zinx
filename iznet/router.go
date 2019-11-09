package iznet


// 路由的抽象接口
// 路由里的数据都是IRequest

type IRouter interface {
	// 处理业务的方法以及钩子方法 之前 - 主 - 之后
	PreHandle(req IRequest)
	Handel(req IRequest)
	PostHandel(req IRequest)

}

