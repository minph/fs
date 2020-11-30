package fs

import (
	"os"
)

// Truncate 截短文件内容，使文件为指定长度
// 传入 0 则清空文件
// 传入负数则截短该数值长度
func Truncate(src string, length int64) error {

	// 传入负数
	if length < 0 {
		content, err := ReadByte(src)
		if err != nil {
			return err
		}

		// 改为截短该长度
		// 切记 length 是负数，用加号
		length = int64(len(content)) + length
	}
	return os.Truncate(src, length)
}
