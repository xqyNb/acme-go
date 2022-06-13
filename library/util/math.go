package util

import (
	"math/rand"
	"time"
)

// RandomInt 生成一个指定区间的随机整数
func RandomInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max) + min
}

// RandomNumber 获得一个随机数字
func RandomNumber() int64 {
	return time.Now().UnixNano()
}
