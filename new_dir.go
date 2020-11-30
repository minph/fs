package fs

import (
	"os"
)

// Dir 目录
type Dir struct {
	Path string
}

// NewDir  创建一个 Dir 对象
func NewDir(path string) *Dir {
	return &Dir{path}
}

// Create 创建文件夹
func (d *Dir) Create() error {
	return CreateDir(d.Path)
}

// Info 原生方法读取目录信息
func (d *Dir) Info() ([]os.FileInfo, error) {

	return DirInfo(d.Path)
}

// Sub 读取目录下所有文件和目录信息
func (d *Dir) Sub() ([]string, error) {
	return DirSub(d.Path)
}

// Subfolder 读取目录下的子文件夹
// 即不包含子文件
func (d *Dir) Subfolder() ([]string, error) {
	return DirSubfolder(d.Path)
}

// Subfile 读取目录下的所有子文件
// 即不包含子目录
func (d *Dir) Subfile() ([]string, error) {
	return DirSubfile(d.Path)
}

// All 读取目录和子目录下的所有文件名信息
//
// 全部为文件名，不再包含目录名
func (d *Dir) All() ([]string, error) {
	return DirAll(d.Path)
}

// Copy 复制文件夹
func (d *Dir) Copy(target string) error {
	return CopyDir(d.Path, target)
}

// Remove 删除文件夹
func (d *Dir) Remove() error {
	return RemoveAll(d.Path)
}

// RemoveContent 删除目录全部内容，但不包含目录自身
func (d *Dir) RemoveContent() error {
	return RemoveDirContent(d.Path)
}

// RemoveExt 删除目录下指定后缀名的文件
func (d *Dir) RemoveExt(ext string) error {
	return RemoveExt(d.Path, ext)
}

// RemoveNamesRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件
func (d *Dir) RemoveNamesRegexp(pattern string) error {
	return RemoveNamesRegexp(d.Path, pattern)
}

// RemoveNamesContain 目录下所有文件中，删除文件名包含给定字符的文件
func (d *Dir) RemoveNamesContain(containStr string) error {
	return RemoveNamesContain(d.Path, containStr)
}

// Move 移动文件夹
func (d *Dir) Move(target string) error {
	return MoveSafe(d.Path, target)
}

// Rename 重命名文件夹
func (d *Dir) Rename(target string) error {
	return MoveSafe(d.Path, target)
}
