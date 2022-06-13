package util

import (
	"log"
)

// 工具类
type Tool struct {
}

func (f *Tool) FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

// tool工具
var tool *Tool

// 获取工具
func GetTool() *Tool {
	if tool != nil {
		return tool
	}
	tool = &Tool{}
	return tool
}
