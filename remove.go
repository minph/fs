package fs

import (
	"os"
)

// Remove 同os.Remove
//
// 只用于删除单个文件和空目录
func Remove(src string) error {

	return os.Remove(src)
}

// RemoveAll 同os.RemoveAll
func RemoveAll(src string) error {

	return os.RemoveAll(src)
}

