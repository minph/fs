package utils

import (
	"os"
	"path"
	"strings"
)

// Format 统一路径格式
func Format(src string) string {
	res := path.Clean(src)
	res = strings.ReplaceAll(res, `\\`, "/")
	res = strings.ReplaceAll(res, "\\", "/")
	return res
}



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
