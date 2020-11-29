package fs

// StringMapFunc 对字符串数组或切片的匿名函数
type StringMapFunc func(index int, file string)

// StringMap 对字符串数组或切片，进行匿名函数操作
func StringMap(stringArray []string, mapFunc StringMapFunc) {

	for index, filename := range stringArray {
		mapFunc(index, filename)
	}
}
