package frame

import "github.com/gin-gonic/gin"

// Rule 规则类型
type Rule struct {
}

// Validate 数据验证器
type Validate struct {
	Name    string
	Type    string
	Require bool
	Msg     string
	Rule    Rule
	Value   string
	IsOk    bool
}

// Check 校验数据
func (v *Validate) Check(c *gin.Context) bool {
	value, exist := c.GetQuery(v.Name)
	// 是否必须
	if v.Require && !exist {
		v.IsOk = false
		return v.IsOk
	}
	// 判断是否存在
	if exist {
		v.Value = value
	} else {
		v.Value = ""
	}
	// 校验类型与规则 TODO:

	v.IsOk = true
	return v.IsOk
}
