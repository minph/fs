package fs

import (
	"io"
	"os"
)

//CopyFile 复制文件
func CopyFile(src, target string) error {
	// 获取文件状态
	_, err := os.Stat(src)
	if err != nil {
		return err
	}

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

// CopyFileSafe 安全地移动文件
func CopyFileSafe(src, target string) error {
	err := CreateFile(src)

	if err != nil {
		return err
	}
	err = CreateFile(target)

	if err != nil {
		return err
	}

	return CopyFile(src, target)
}
