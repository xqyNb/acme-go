package validate

import "acme/frame"

// UserValidate 用户验证器
type UserValidate struct {
	BaseValidate
}

// Add 添加校验
func (v *UserValidate) Add(path string, validates []frame.Validate) {
	if v.Validates == nil {
		v.Validates = make(map[string][]frame.Validate)
	}
	v.Validates[path] = validates
}

const (
	SceneUserlogin = "UserLogin"
)

// GetUserValidate 获取验证器
func GetUserValidate() *UserValidate {
	userValidate := &UserValidate{}
	// 验证变量
	// 用户名
	username := frame.Validate{
		Require: true,
		Name:    "username",
		Type:    "string",
		Msg:     "用户名必须",
	}
	// 密码
	password := frame.Validate{
		Require: true,
		Name:    "password",
		Type:    "string",
		Msg:     "密码必须",
	}

	// 用户登录
	var userLogin []frame.Validate
	userLogin = append(userLogin, username, password)

	// 添加场景
	userValidate.Add(SceneUserlogin, userLogin)

	return userValidate
}
