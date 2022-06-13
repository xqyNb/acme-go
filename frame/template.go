package frame

import (
	"acme/library/libprint"
	"acme/library/util"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	// TemplateBasePath 模版的基础路径
	TemplateBasePath = "templates"
)

// TemplateManager 模版管理器
type TemplateManager struct {
	pathList map[string]string
}

// gin框架的模版，直接使用 文件夹路径/*,匹配下面的所有文件! /**/* 恶心的一B，需要在模版文件里面注入代码，强烈不建议使用!
// （ gin框架会实时的去读取模版文件并输出，没有缓存 ）

// 注册模板
func (t *TemplateManager) registe(ginEngine *gin.Engine) {
	// 初始化list
	if t.pathList == nil {
		t.pathList = make(map[string]string)
	}
	// 模版只存在于 get 与 notFond中 其他就不用处理了

	// 加载所有的html文件！ 注意 : LoadHTMLFiles和LoadHTMLGlob只能使用一次，再次使用将覆盖前一次
	// 添加路由的 getControllerList
	routerManager := acme.GetRouteManager()
	for _, controller := range routerManager.getControllerList {
		for _, path := range controller.Templates {
			files := strings.Split(path, "/")
			t.pathList[path] = files[len(files)-1]
		}
	}
	// Not Found
	if routerManager.notFoundController != nil {
		for _, path := range routerManager.notFoundController.Templates {
			files := strings.Split(path, "/")
			t.pathList[path] = files[len(files)-1]
		}
	}

	// debug
	fmt.Println(t.pathList)
	// 需要一个泛型 等待1.8版本 TODO:
	var pathList []string
	var fileList []string
	// 初始化路径列表 - 检测重名
	for path, file := range t.pathList {
		if util.InSlice(file, fileList) {
			panic(errors.New(fmt.Sprintf("模版文件名重名了![ %s ] -> [ %s ]", path, file)))
		} else {
			fileList = append(fileList, file)
			pathList = append(pathList, TemplateBasePath+path)
		}
	}
	// 没有任何重名的文件则加载
	ginEngine.LoadHTMLFiles(pathList...)

	// 显示提示模版注册
	for _, path := range pathList {

		// 判断是否是debug模式
		if !acme.config.App.Production {
			libprint.PrintColorln(fmt.Sprintf("模版注册 -> %s", path), libprint.TextGreen)
		}
	}
}
