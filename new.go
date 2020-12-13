package fs

import (
	"github.com/minph/fs/utils"
	"path"
)

// File 文件类型结构体
type File struct {
	utils.Wrapper
	Path    string      //文件地址
	DirName string      // 文件目录名
	Res     interface{} // 函数返回值
}

// New 创建文件类型结构体
func New(src string) *File {
	return &File{
		Path:    src,
		DirName: path.Dir(src),
	}
}

// Create 创建文件
func (f *File) Create() *File {
	err := Create(f.Path)
	f.AddError(err)
	return f
}

// Rename 重命名文件
func (f *File) Rename(name string) *File {
	err := Rename(f.Path, name)
	f.AddError(err)
	return f
}

// Rename 重命名文件
func (f *File) Move(target string) *File {
	err := Move(f.Path, target)
	f.AddError(err)
	return f
}

// Copy 复制文件
func (f *File) Copy(target string) *File {
	err := Copy(f.Path, target)
	f.AddError(err)
	return f
}

// CopySafe 安全地复制文件
func (f *File) CopySafe(target string) *File {
	err := CopySafe(f.Path, target)
	f.AddError(err)
	return f
}

// ReadAll 读取文件全部内容并返回字符串
func (f *File) ReadAll() *File {
	res, err := ReadAll(f.Path)
	f.AddError(err)
	f.Res = res
	return f
}

// ReadLines 按行读取文件内容
func (f *File) ReadLines() *File {
	res, err := ReadLines(f.Path)
	f.AddError(err)
	f.Res = res
	return f
}

// ReadBytes 按 byte 块方式读取文件内容
func (f *File) ReadBytes(bufferSize int64) *File {
	res, err := ReadBytes(f.Path, bufferSize)
	f.AddError(err)
	f.Res = res
	return f
}

// ReadLineAt 读取指定行的内容
func (f *File) ReadLineAt(row int) *File {
	res, err := ReadLineAt(f.Path, row)
	f.AddError(err)
	f.Res = res
	return f
}

// Truncate 截短文件内容，使文件为指定长度
// 传入 0 则清空文件
// 传入负数则截短该数值长度
func (f *File) Truncate(length int64) *File {
	err := Truncate(f.Path, length)
	f.AddError(err)
	return f
}

// Rewrite 对文件覆盖写入数据
func (f *File) Rewrite(content string) *File {
	err := Rewrite(f.Path, content)
	f.AddError(err)
	return f
}

// Clear 清空文件内容
func (f *File) Clear() *File {
	err := Clear(f.Path)
	f.AddError(err)
	return f
}

// Append 以字符串方式追加文件内容
func (f *File) Append(content string) *File {
	err := Append(f.Path, content)
	f.AddError(err)
	return f
}

// WriteLine 在文件指定行写入内容
// 传入字符串数据即可
// row 修改的行数
// replace 是否替换当前行
func (f *File) WriteLine(content string, row int, replace bool) *File {
	err := WriteLine(f.Path, content, row, replace)
	f.AddError(err)
	return f
}

// Check 检查错误
func (f *File) Check(call func(err []error)) *File {
	if f.GetError() {
		call(f.Err)
	}
	return f
}

// Do 对结构体进行链式操作
func (f *File) Do(call func(f *File)) *File {
	call(f)
	return f
}
