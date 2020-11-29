package fs

import "os"

// Rename 重命名文件
func Rename(src, target string) error {
	return os.Rename(src, target)
}

// MoveFile 移动文件
//
// 和 Rename 同义，函数名不一致是为了语义的区别
func MoveFile(src, target string) error {
	return Rename(src, target)
}

// MoveFileSafe 安全地移动文件
func MoveFileSafe(src, target string) error {
	err := CreateFile(src)

	if err != nil {
		return err
	}
	err = CreateFile(target)

	if err != nil {
		return err
	}

	return MoveFile(src, target)

}
