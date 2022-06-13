package frame

import (
	"acme/frame/frame-data"
	"acme/library/libprint"
	"acme/library/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// PanicFuc 定义异常处理函数
type PanicFuc func(ctx *gin.Context, data frame_data.PanicData)

// GlobalMiddleware 此中间件指的是引擎（全局）中间件
// Controller（控制器）的中间件直接写给controller即可
type GlobalMiddleware struct {
	middleWareList []gin.HandlerFunc
	isRegiste      bool
	panicFuc       PanicFuc
}

// 初始化默认中间件
func (m *GlobalMiddleware) ini() {
	// 添加日志中间件
	m.Add(func(c *gin.Context) {
		t := time.Now()
		// 请求之前
		requestTime := util.GetCurrentTime()
		requestKey := util.RandomUniqueString(32)
		c.Set(frame_data.KeyRequestKey, requestKey)
		// 请求数据
		requestData := frame_data.RequestData{
			RequestKey: requestKey,
			ClientIp:   c.ClientIP(),
			Time:       requestTime,
			Method:     c.Request.Method,
			Path:       c.FullPath(),
			Protocal:   c.Request.Proto,
			UserAgent:  c.Request.UserAgent(),
			Uri:        c.Request.RequestURI,
		}
		// 记录请求日志数据
		go WriteRequestLog(acme.config.App.LogPath, requestData)
		// 开始请求
		c.Next()
		// 请求结束
		latency := time.Since(t)
		// 响应数据
		responseData := frame_data.ResponseData{
			RequestKey: requestKey,
			Time:       requestTime,
			Method:     c.Request.Method,
			Latency:    latency.String(),
			StatusCode: c.Writer.Status(),
			ApiData:    c.GetString(KeyApiData),
		}
		// 记录响应数据
		go WriteResponseLog(acme.config.App.LogPath, responseData)

		fmt.Println("requsetData : ", requestData)
		fmt.Println("responseData : ", responseData)

	})

	// 添加错误异常中间件
	m.Add(func(c *gin.Context) {
		defer func() {
			// 捕捉异常
			if err := recover(); err != nil {
				// 请求之前
				requestTime := util.GetCurrentTime()
				requestKey := c.GetString(frame_data.KeyRequestKey)

				var detail string
				// 判断recover的是不是一个error
				if e, ok := err.(error); ok {
					detail = e.Error()
				} else {
					detail = fmt.Sprintf("%s", err)
				}

				// panic数据
				panicData := frame_data.PanicData{
					ReqData: frame_data.RequestData{
						RequestKey: requestKey,
						ClientIp:   c.ClientIP(),
						Time:       requestTime,
						Method:     c.Request.Method,
						Path:       c.FullPath(),
						Protocal:   c.Request.Proto,
						UserAgent:  c.Request.UserAgent(),
						Uri:        c.Request.RequestURI,
					},
					Detail: detail,
				}

				// 异常警告
				libprint.PrintErrorHint(libprint.HintPanic, "发生一个异常!详情请看日志!")

				// 记录panic日志
				WritePanicLog(acme.config.App.LogPath, panicData)

				// 回调错误处理
				if m.panicFuc != nil {
					m.panicFuc(c, panicData)
				} else {
					// 自定义错误响应
					c.String(http.StatusInternalServerError, "服务器错误!请稍后再试!")
				}
			}

		}()
		// 运行控制器
		c.Next()
	})
}

// 注册中间件 - 注意！本操作需要紧接着引擎初始化后
func (m *GlobalMiddleware) regiseMiddleware(ginEngine *gin.Engine) {
	if m.isRegiste {
		return
	}
	m.isRegiste = true
	// 初始化
	m.ini()
	// 使用中间件
	ginEngine.Use(m.middleWareList...)

	// 判断有没有错误处理
	if m.panicFuc == nil {
		// 提醒一下
		libprint.PrintHint(libprint.HintWarning, "您尚未设置控制器错误响应处理！请调用 GlobalMiddleware.SetPanicFunc设置错误处理!")
	}

	//ginEngine.Use(gin.Recovery())
}

// SetPanicFunc 设置Panic处理函数
func (m *GlobalMiddleware) SetPanicFunc(panicFuc PanicFuc) {
	m.panicFuc = panicFuc
}

// Add 添加中间件
func (m *GlobalMiddleware) Add(middleWare gin.HandlerFunc) {
	m.middleWareList = append(m.middleWareList, middleWare)
}
