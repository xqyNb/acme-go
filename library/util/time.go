package util

import (
	"fmt"
	"time"
)

const (
	Year   = "2006"
	Month  = "01"
	Day    = "02"
	Hour   = "15"
	Minute = "04"
	Second = "05"
)

// GetDateFormat 获取默认的日期格式
func GetDateFormat() string {
	return fmt.Sprintf("%s-%s-%s %s:%s:%s", Year, Month, Day, Hour, Minute, Second)
}

// GetCurrentTime 获取当前的时间
func GetCurrentTime() string {
	return time.Now().Format(GetDateFormat())
}
