package znet

import (
	"zinx/iznet"
)

// 让自定义的router 去继承Router
// 让自定义的router 对基类的Handle去重写
type Router struct {}

func (r *Router) PreHandle(req iznet.IRequest) {}

func (r *Router) Handel(req iznet.IRequest) {}

func (r *Router) PostHandel(req iznet.IRequest) {}


