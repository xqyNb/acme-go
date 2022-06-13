package controller

import (
	"github.com/gin-gonic/gin"
)

// Index 首页相关控制器
type Index struct {
	BaseController
}

// GetIndex GetUser 获取获取首页控制器
func GetIndex() *Index {
	var templates []string
	// 初始化模版
	templates = append(templates, "/index/home.html")

	return &Index{
		BaseController{
			Templates: templates,
		},
	}
}

// 首页页面
func (i Index) index(c *gin.Context) {
	//c.HTML(http.StatusOK, "home.html", gin.H{
	//	"master": "billion",
	//})
	acmeFrame.HtmlTemplate(c, "home.html", gin.H{"master": "首页"})
}
