package fs

import (
	"os"
)

// ensureFile 检查文件是否存在
// 不存在则创建文件
func ensureFile(src string) error {
	_, err := os.Stat(src)

	// 文件不存在，则创建文件
	if os.IsNotExist(err) {
		return CreateFile(src)
	}
	return err

}
