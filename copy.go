package fs

import (
	"io"
	"os"
	"path"

	"github.com/minph/fs/utils"
)

// Create 跨目录创建文件
func Create(src string) error {

	// 创建目录
	folder := path.Dir(src)
	if !utils.Exist(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 创建文件
	if !utils.Exist(src) {
		file, err := os.Create(src)
		defer file.Close()
		return err
	}
	return nil
}

//Copy 复制文件
func Copy(src, target string) error {

	// 打开待处理文件
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// 打开目标地址
	destFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// 复制文件
	_, err = io.Copy(destFile, srcFile)
	return err
}

// CopySafe 安全地复制文件
//
// 保证复制位置文件存在
func CopySafe(src, target string) error {

	// 不存在则创建
	if !utils.Exist(src) {

		err := Create(src)

		if err != nil {
			return err
		}

	}

	// 不存在则创建
	if !utils.Exist(target) {

		err := Create(target)

		if err != nil {
			return err
		}

	}

	return Copy(src, target)
}
