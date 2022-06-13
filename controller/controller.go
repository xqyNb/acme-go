package controller

import (
	"acme/common/business"
	"acme/frame"
	"github.com/gin-gonic/gin"
)

// BaseController 基础控制器
type BaseController struct {
	// 控制器需要的模版列表
	Templates []string
}

// ApiParameter 参数错误响应
func (b *BaseController) ApiParameter(c *gin.Context, va *frame.Validate) {
	msg := va.Msg
	acmeFrame.GetApi().Fail(c, business.CodeParameterError, msg, gin.H{va.Name: va.Value})
}

// acme框架
var acmeFrame *frame.Acme

// RegisteControllers 注册控制器
func RegisteControllers(acme *frame.Acme) {
	// 初始化框架
	acmeFrame = acme

	// 注册路由
	routerManager := acme.GetRouteManager()
	// 首页
	index := GetIndex()
	routerManager.AddGet(&frame.Controller{
		Path:      "/",
		Action:    index.index,
		Templates: index.Templates,
	})

	// 添加控制器组
	userGroup := frame.ControllerGroup{Path: "/user"}
	{
		user := GetUser()
		userGroup.AddPost(&frame.Controller{Path: "/login", Action: user.login, Templates: user.Templates})
		userGroup.AddPost(&frame.Controller{Path: "/registe", Action: user.registe, Templates: user.Templates})
	}

	// 404
	errorController := GetErrorController()
	routerManager.SetNotFoundController(&frame.Controller{
		Action:    errorController.NotFound,
		Templates: errorController.Templates,
	})
}
