package fs

import "os"

// Rename 重命名文件
func Rename(src, target string) error {
	return os.Rename(src, target)
}

// Move 移动文件或文件夹
//
// 和 Rename 同义，函数名不一致是为了语义的区别
func Move(src, target string) error {
	return Rename(src, target)
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
