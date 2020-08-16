package handlers

import "github.com/kataras/iris/v12"

//NotFound -handler 前端请求地址错误，调用此handler处理
func NotFound(ctx iris.Context) {
	println("404-找不到此路由/路径:", ctx.RequestPath(true))
	ctx.WriteString("路由/请求地址错误")
}
