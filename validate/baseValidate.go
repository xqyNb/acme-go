package validate

import (
	"acme/frame"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

// BaseValidate 基础验证器
type BaseValidate struct {
	Validates     map[string][]frame.Validate
	ErrorValidate *frame.Validate
}

// CheckData 校验数据
func (v *BaseValidate) CheckData(c *gin.Context, scene string) (map[string]string, bool) {
	data := make(map[string]string)
	// 检测场景
	validates, ok := v.Validates[scene]
	if !ok {
		panic(errors.New(fmt.Sprintf("验证场景[ %s ]不存在! 请在Validate中设置!", scene)))
	}
	// 循环检测
	for _, validate := range validates {
		// 检测是否符合
		if validate.Check(c) {
			data[validate.Name] = validate.Value
		} else { // 检测失败 - 返回api响应
			v.ErrorValidate = &validate
			return nil, false
		}
	}

	return data, true
}
