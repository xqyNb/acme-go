package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ErrorController 404控制器
type ErrorController struct {
	BaseController
}

// GetErrorController 获取ErrorController
func GetErrorController() *ErrorController {
	var templates []string
	// 初始化模版
	templates = append(templates, "/error/404.html")

	return &ErrorController{
		BaseController{
			Templates: templates,
		},
	}
}

// NotFound 404页面
func (e *ErrorController) NotFound(c *gin.Context) {
	acmeFrame.HtmlTemplateWithStatus(c, http.StatusNotFound, "404.html", gin.H{})
}
