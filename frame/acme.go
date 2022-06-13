package frame

import (
	"acme/library/libprint"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
)

// Acme 框架
type Acme struct {
	Name     string
	Email    string
	Author   string
	Verstion string
	// 配置信息
	config *Config
	// 中间件
	middleware *GlobalMiddleware
	// 路由管理器
	routerManager *RouterManager
	// 模版管理器
	templateManager *TemplateManager
}

// GetMiddleware 获取中间件
func (a *Acme) GetMiddleware() *GlobalMiddleware {
	if a.middleware == nil {
		a.middleware = &GlobalMiddleware{}
	}
	return a.middleware
}

// GetRouteManager 获取路由管理器
func (a *Acme) GetRouteManager() *RouterManager {
	if a.routerManager == nil {
		a.routerManager = &RouterManager{}
	}
	return a.routerManager
}

// GetApi 获取Api
func (a *Acme) GetApi() *Api {
	return &Api{}
}

// GetHtml 获取Html
func (a *Acme) GetHtml() *Html {
	return &Html{}
}

// Html 响应纯html
func (a *Acme) Html(c *gin.Context, body string) {
	a.HtmlWithStatus(c, http.StatusOK, body)
}

// HtmlWithStatus 响应HTML并给个自定义状态
func (a *Acme) HtmlWithStatus(c *gin.Context, status int, body string) {
	a.GetHtml().Response(c, status, body)
}

// HtmlTemplate 响应一个html模版
func (a *Acme) HtmlTemplate(c *gin.Context, name string, assign any) {
	a.HtmlTemplateWithStatus(c, http.StatusOK, name, assign)
}

// HtmlTemplateWithStatus 响应一个html模版并给个自定义状态
func (a *Acme) HtmlTemplateWithStatus(c *gin.Context, status int, name string, assign any) {
	// 将路径注入到
	SetHeaderServer(c)
	c.HTML(status, name, assign)
}

// GetTemplateManager 获取模版管理器
func (a *Acme) GetTemplateManager() *TemplateManager {
	if a.templateManager == nil {
		a.templateManager = &TemplateManager{}
	}
	return a.templateManager
}

// Run 运行
func (a *Acme) Run() {
	// 输出版本
	a.printVersion()
	// 初始化配置
	a.iniConfig()
	// 启动服务
	a.startServer()
}

// 启动服务
func (a *Acme) startServer() {
	// 初始化
	ginEngine := gin.New()
	// 注册模板
	a.GetTemplateManager().registe(ginEngine)

	// 注册中间件，SB gin框架 - 中间件的调用位置应该在其他方法之前调用
	a.GetMiddleware().regiseMiddleware(ginEngine)
	// 注册路由
	a.GetRouteManager().registeRouter(ginEngine)

	// 控制器参数绑定 TODO:
	// 优雅重启 TODO：

	// 设置writer
	gin.DefaultWriter = io.MultiWriter(os.Stdout)

	// 设置信任的代理
	err := ginEngine.SetTrustedProxies(a.config.App.TrustedProxies)
	if err != nil {
		panic(err.Error())
	}

	// 设置地址
	bindAddr := fmt.Sprintf("%s:%d", a.config.App.Host, a.config.App.Port)
	// 输出信息
	libprint.PrintColorln(fmt.Sprintf("Acme 已启动！绑定地址 : %s", bindAddr), libprint.TextGreen)
	// 判断是否是开发者模式
	if a.config.App.Production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		libprint.PrintHint(libprint.HintWarning, "您现在处于【 开发者 】模式下！正式环境中请配置production = true.")
	}
	// 启动
	err = ginEngine.Run(bindAddr)
	if err != nil {
		panic(err.Error())
	}

}

// 初始化配置
func (a *Acme) iniConfig() {
	// 读取配置文件
	err, conifg := iniConfig()
	if err != nil {
		libprint.PrintColorln(err.Error(), libprint.TextRed)
		os.Exit(1)
	}
	a.config = conifg

	//fmt.Println(a.config)
}

// 输出版本信息
func (a *Acme) printVersion() {
	a.Name = "Acme"
	a.Verstion = "v1.0.1"
	a.Author = "Billion"
	a.Email = "billionzx@qq.com"

	fmt.Printf("Acme Frame [ %s ] 欢迎使用%s框架,本框架基于gin[ %s ],作者: %s 😄,邮箱: %s \n",
		a.Verstion,
		a.Name,
		gin.Version,
		a.Author,
		a.Email)
	//fmt.("欢迎使用%s")
}

// acme实例
var acme *Acme

// GetAcme 获取acme
func GetAcme() *Acme {
	if acme == nil {
		acme = &Acme{}
	}
	return acme
}
