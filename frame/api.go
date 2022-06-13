package frame

import (
	frame_data "acme/frame/frame-data"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ApiData 响应数据
type ApiData struct {
	Code int         `json:"code"`
	Key  string      `json:"key"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// KeyApiData 设置api响应数据key
const (
	KeyApiData = "KEY_API_DATA"
	// ApiCodeOk api响应的业务码 响应成功
	ApiCodeOk = 200
)

// Api 接口响应
type Api struct {
}

// NoData 空数据
func (a *Api) NoData() []string {
	return []string{}
}

// 响应
func (a *Api) response(c *gin.Context, httpCode int, code int, msg string, data interface{}) {
	apiData := ApiData{
		Code: code,
		Key:  c.GetString(frame_data.KeyRequestKey),
		Msg:  msg,
		Data: data,
	}
	// 返回json
	response, err := json.Marshal(apiData)
	if err != nil {
		panic(err.Error())
	}
	// 设置api响应数据
	c.Set(KeyApiData, string(response))

	// 设置响应头
	c.Header("Content-Type", "application/json")
	// 设置响应头Server为框架名称
	SetHeaderServer(c)
	// 响应给客户端
	c.String(httpCode, string(response))
}

// SetHeaderServer 设置响应头Server为框架名称
func SetHeaderServer(c *gin.Context) {
	c.Header("Serever", fmt.Sprintf("%s %s", acme.Name, acme.Verstion))
}

// Success 响应成功
func (a *Api) Success(c *gin.Context, msg string, data interface{}) {
	a.response(c, http.StatusOK, ApiCodeOk, msg, data)
}

// Fail 失败响应
func (a *Api) Fail(c *gin.Context, code int, msg string, data interface{}) {
	a.response(c, http.StatusBadRequest, code, msg, data)
}
