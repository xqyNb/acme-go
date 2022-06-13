package util

// InSlice 判断元素是否在slice中
func InSlice[T ~int | ~string](element T, slice []T) bool {
	for _, el := range slice {
		if el == element {
			return true
		}
	}
	return false
}
