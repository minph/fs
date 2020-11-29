package fs

import "os"

// Exist 判断所给路径文件或文件夹是否存在
func Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
