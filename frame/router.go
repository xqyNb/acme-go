package frame

import (
	"acme/library/libprint"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RouterManager 路由管理器
type RouterManager struct {
	getControllerList  map[string]*Controller
	postControllerList map[string]*Controller
	notFoundController *Controller
}

// 设置默认首页控制器
func (r *RouterManager) setDefaultIndexController(ginEngine *gin.Engine) {
	indexPath := "/"
	// 判断是否设置在了 getControllerList 中
	if controller, ok := r.getControllerList[indexPath]; ok {
		// 设置指定的首页
		router := ginEngine.GET(indexPath, controller.Action)
		if len(controller.middleware) != 0 {
			router.Use(controller.middleware...)
		}
		// 从Get控制器列表中删除
		delete(r.getControllerList, indexPath)
	} else { // 没有 - 设置默认首页并提醒
		// 设置默认首页
		ginEngine.GET(indexPath, func(c *gin.Context) {
			welcome := fmt.Sprintf("欢迎使用框架[ %s - %s ], Author : %s 😄",
				acme.Name, acme.Verstion, acme.Author)
			c.String(http.StatusOK, welcome)
		})
		// 提醒一下
		libprint.PrintHint(libprint.HintWarning, "您尚未设置默认的首页控制器!首页将展示框架的默认首页")
	}
}

// 设置默认的404控制器
func (r *RouterManager) setDefaultNotFound(ginEngine *gin.Engine) {
	// 判断有没有设置404路由
	if r.notFoundController != nil {
		var actions []gin.HandlerFunc
		actions = append(actions, r.notFoundController.Action)
		if len(r.notFoundController.middleware) > 0 {
			actions = append(actions, r.notFoundController.middleware...)
		}
		// 添加action和中间件
		ginEngine.NoRoute(actions...)
	} else {
		// 设置默认路由，并提示开发者!
		ginEngine.NoRoute(func(c *gin.Context) {
			c.String(http.StatusNotFound, fmt.Sprintf("框架[ %s - %s ],404 Not Found！您访问的页面没有找到!",
				acme.Name, acme.Verstion))
		})
		// 提醒一下
		libprint.PrintHint(libprint.HintWarning, "404 页面为框架默认！建议您设置一下自己的404页面.")
	}
}

// SetNotFoundController 设置404 控制器
func (r *RouterManager) SetNotFoundController(controller *Controller) {
	r.notFoundController = controller
}

// 注册路由控制器
func (r *RouterManager) registeRouter(ginEngine *gin.Engine) {
	// 设置默认的首页控制器
	r.setDefaultIndexController(ginEngine)
	// 设置默认的404控制器
	r.setDefaultNotFound(ginEngine)
	// 设置Get控制器
	for path, getController := range r.getControllerList {
		router := ginEngine.GET(path, getController.Action)
		if len(getController.middleware) > 0 {
			router.Use(getController.middleware...)
		}
	}
	// 设置post路由
	for path, postController := range r.postControllerList {
		router := ginEngine.POST(path, postController.Action)
		if len(postController.middleware) > 0 {
			router.Use(postController.middleware...)
		}
	}
}

// AddGet 添加Get控制器 - 注意: 不能注册相同路径的控制器！否则gin框架报错！
func (r *RouterManager) AddGet(controller *Controller) {
	// 初始化map
	if r.getControllerList == nil {
		r.getControllerList = make(map[string]*Controller)
	}
	r.getControllerList[controller.Path] = controller
}

// AddPost 添加Post控制器 - 注意: 不能注册相同的控制器！否则gin框架控制器！
func (r *RouterManager) AddPost(controller *Controller) {
	// 初始化map
	if r.postControllerList == nil {
		r.postControllerList = make(map[string]*Controller)
	}
	r.postControllerList[controller.Path] = controller
}
