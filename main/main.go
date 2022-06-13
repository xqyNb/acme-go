package main

import (
	"acme/controller"
	"acme/frame"
	framedata "acme/frame/frame-data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	acme := frame.GetAcme()
	//注册中间件
	middleware := acme.GetMiddleware()
	// 设置错误处理
	middleware.SetPanicFunc(func(ctx *gin.Context, data framedata.PanicData) {
		ctx.String(http.StatusInternalServerError, "服务器繁忙！请稍后再试!")
	})

	// 注册控制器
	controller.RegisteControllers(acme)

	acme.Run()
}
