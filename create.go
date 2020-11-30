package fs

import (
	"os"
	"path"
)

// CreateDir 跨目录创建文件夹
func CreateDir(folder string) error {
	if !Exist(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// CreateFile 跨目录创建文件
func CreateFile(src string) error {

	// 创建目录
	err := CreateDir(path.Dir(src))

	if err != nil {
		return nil
	}

	// 创建文件
	if !Exist(src) {
		file, err := os.Create(src)
		defer file.Close()
		return err
	}
	return nil
}
