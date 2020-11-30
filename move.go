package fs

import (
	"os"
	"path"
)

// Rename 重命名文件
// 不修改目录路径，只修改文件名
func Rename(src, newName string) error {
	newPath := path.Join(path.Dir(src), newName)
	return os.Rename(src, newPath)
}

// Move 移动文件或文件夹
func Move(src, target string) error {
	return os.Rename(src, target)
}

// MoveSafe 安全地移动文件或文件夹
// 保证跨目录的安全性
func MoveSafe(src, target string) error {
	err := CreateFile(src)

	if err != nil {
		return err
	}
	err = CreateFile(target)

	if err != nil {
		return err
	}

	return Move(src, target)

}
