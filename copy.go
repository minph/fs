package fs

import (
	"io"
	"os"
	"path"
	"path/filepath"
)

//CopyFile 复制文件
func CopyFile(src, target string) error {

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

// CopyFileSafe 安全地复制文件
//
// 保证复制位置文件存在
func CopyFileSafe(src, target string) error {

	// 不存在则创建
	if !Exist(src) {

		err := CreateFile(src)

		if err != nil {
			return err
		}

	}

	// 不存在则创建
	if !Exist(target) {

		err := CreateFile(target)

		if err != nil {
			return err
		}

	}

	return CopyFile(src, target)
}

// CopyDir 复制文件夹
func CopyDir(folder, target string) error {

	// 获取文件夹全部文件信息
	info, err := DirAll(folder)
	if err != nil {
		return err
	}

	for _, file := range info {

		// 获取相对路径
		relaPath, err := filepath.Rel(folder, file)

		if err != nil {
			return err
		}

		// 逐一复制文件
		// 使用 ToSlash 防止反斜杠
		copyFile := path.Join(target, filepath.ToSlash(relaPath))
		err = CopyFileSafe(file, copyFile)
		if err != nil {
			return err
		}
	}
	return nil
}
