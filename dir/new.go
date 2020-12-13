package dir

import (
	"github.com/minph/fs/utils"
)

// Dir 目录类型结构体
type Dir struct {
	utils.Wrapper
	Path string      //目录地址
	Res  interface{} // 函数返回值
}

// New 创建目录类型结构体
func New(src string) *Dir {
	return &Dir{
		Path: src,
	}
}

// Create 跨目录创建文件夹
func (d *Dir) Create() *Dir {
	err := Create(d.Path)
	d.AddError(err)
	return d
}

// Copy 复制文件夹
func (d *Dir) Copy(target string) *Dir {
	err := Copy(d.Path, target)
	d.AddError(err)
	return d
}

// Is 判断是不是文件夹
func (d *Dir) Is() *Dir {
	res := Is(d.Path)
	d.Res = res
	return d
}

// Info 原生方法读取目录信息
func (d *Dir) Info() *Dir {
	res, err := Info(d.Path)
	d.AddError(err)
	d.Res = res
	return d
}

// Content 读取目录下所有文件和目录信息
func (d *Dir) Content() *Dir {
	res, err := Content(d.Path)
	d.AddError(err)
	d.Res = res
	return d
}

// Folders 读取目录下的子文件夹
// 即不包含子文件
func (d *Dir) Folders() *Dir {
	res, err := Folders(d.Path)
	d.AddError(err)
	d.Res = res
	return d
}

// Files 读取目录下的所有子文件
// 即不包含子目录
func (d *Dir) Files() *Dir {
	res, err := Files(d.Path)
	d.AddError(err)
	d.Res = res
	return d
}

// All 读取目录和子目录下的所有文件名信息
//
// 全部为文件名，不再包含目录名
func (d *Dir) All() *Dir {
	res, err := All(d.Path)
	d.AddError(err)
	d.Res = res
	return d
}

// Map 读取目录和子目录下的所有文件名信息，然后进行操作
//
// 全部为文件名，不再包含目录名
func (d *Dir) Map(call func(index int, file string)) *Dir {
	err := Map(d.Path, call)
	d.AddError(err)
	return d
}

// RemoveContent  删除目录全部内容，但不包含目录自身
func (d *Dir) RemoveContent() *Dir {
	err := RemoveContent(d.Path)
	d.AddError(err)
	return d
}

// RemoveExt 删除目录下指定后缀名的文件
func (d *Dir) RemoveExt(ext string) *Dir {
	err := RemoveExt(d.Path, ext)
	d.AddError(err)
	return d
}

// RemoveRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件
func (d *Dir) RemoveRegexp(pattern string) *Dir {
	err := RemoveRegexp(d.Path, pattern)
	d.AddError(err)
	return d
}

// RemoveContain 目录下所有文件中，删除文件名包含给定字符的文件
func (d *Dir) RemoveContain(containStr string) *Dir {
	err := RemoveRegexp(d.Path, containStr)
	d.AddError(err)
	return d
}
