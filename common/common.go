package common

// 交换方式 避免了声明  temp  = a[i]   a[i] = a[j]   a[j] = temp 的这种形式
func Swap(slice []int, i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
