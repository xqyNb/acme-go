package controller

import (
	"acme/validate"
	"github.com/gin-gonic/gin"
)

// User 用户相关控制器
type User struct {
	BaseController
}

// GetUser 获取user控制器
func GetUser() *User {
	var templates []string
	// 初始化模版
	templates = append(templates, "/user/userLogin.html")
	templates = append(templates, "/user/userRegiste.html")

	return &User{
		BaseController{
			Templates: templates,
		},
	}
}

// 登陆页面
func (u User) login(c *gin.Context) {
	// 检测数据
	userValidate := validate.GetUserValidate()
	data, ok := userValidate.CheckData(c, validate.SceneUserlogin)
	if ok {
		//acmeFrame.Html(c, fmt.Sprintf("username = %s,password = %s", data["username"], data["password"]))
		acmeFrame.GetApi().Success(c, "ok", data)
	} else {
		u.ApiParameter(c, userValidate.ErrorValidate)
	}

	//userName := c.Query("userName")
	//password := c.Query("password")
	//acmeFrame.Html(c, fmt.Sprintf("username = %s,password = %s", userName, password))
	//acmeFrame.HtmlTemplate(c, "userLogin.html", gin.H{"master": "billion"})
}

// 注册用户
func (u *User) registe(c *gin.Context) {
	acmeFrame.HtmlTemplate(c, "userRegiste.html", gin.H{"master": "billion"})

}
