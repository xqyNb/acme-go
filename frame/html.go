package frame

import "github.com/gin-gonic/gin"

// Html 管理器
type Html struct {
}

// Response 响应html内容
func (h *Html) Response(c *gin.Context, httpCode int, body string) {
	// 设置响应头Server为框架名称
	SetHeaderServer(c)
	// 输出body体
	c.String(httpCode, body)
}
