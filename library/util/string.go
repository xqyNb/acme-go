package util

import "fmt"

// RandomUniqueNumber 生成一个唯一的数字字符串
func RandomUniqueNumber() string {
	start := RandomInt(10000, 99999)
	middle := RandomNumber()
	end := RandomInt(10000, 99999)
	return fmt.Sprintf("%d.%d.%d", start, middle, end)
}

// RandomUniqueString 生成一个唯一的字符串
func RandomUniqueString(bit uint) string {
	base := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var str = ""
	for i := 0; i < int(bit); i++ {
		randomBit := RandomInt(1, len(base))
		str += base[randomBit-1 : randomBit]
	}

	return str
}
