package frame

import "github.com/gin-gonic/gin"

// Controller 控制器
type Controller struct {
	// 路径
	Path string
	// 动作
	Action gin.HandlerFunc
	// 控制器中间件
	middleware []gin.HandlerFunc
	// template模版
	Templates []string
}

// AddMiddleware 添加中间件
func (c *Controller) AddMiddleware(middleware gin.HandlerFunc) *Controller {
	c.middleware = append(c.middleware, middleware)
	return c
}

// ControllerGroup 控制器组
type ControllerGroup struct {
	Path string
	//getList  []*Controller
	//postList []*Controller
}

// Group 嵌套Group
func (c *ControllerGroup) Group(path string) *ControllerGroup {
	return &ControllerGroup{Path: c.Path + path}
}

// AddGet 添加Get控制器
func (c *ControllerGroup) AddGet(controller *Controller) *ControllerGroup {
	controller.Path = c.Path + controller.Path
	acme.routerManager.AddGet(controller)
	return c
}

// AddPost 添加Post控制器
func (c *ControllerGroup) AddPost(controller *Controller) *ControllerGroup {
	controller.Path = c.Path + controller.Path
	acme.routerManager.AddPost(controller)
	return c
}
