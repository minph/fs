package fs

import (
	"fmt"
	"os"
	"path"
)

// File 文件
type File struct {
	path string
	Rows int
}

// Common 文件夹和文件公用方法接口
type Common interface {
	Create() error
	Copy(target string) error
	Remove() error
	Move(target string) error
	Rename(newName string) error
	Exist() bool
	IsDir() bool
	Path() string
	Join(src string) string
	JoinBase(src string) string
}

// New 在不判断文件夹还是文件的情况下，支持通用方法接口
func New(path string) Common {
	if IsDir(path) {
		return NewDir(path)
	}
	return NewFile(path)
}

// AsDir 类型断言[*fs.Dir]
func AsDir(c Common) (*Dir, bool) {
	value, ok := c.(*Dir)

	return value, ok
}

// AsFile 类型断言[*fs.File]
func AsFile(c Common) (*File, bool) {
	value, ok := c.(*File)
	return value, ok
}

// FileKind 文件方法接口
type FileKind interface {
	Common
	Clear() error
	ReadBytes() ([]byte, error)
	ReadString() (string, error)
	ReadByRow() ([]string, error)
	ReadByBytes(bufferSize int64) ([][]byte, error)
	ReadAt(position int64) ([]byte, error)
	ReadStringAt(position int64) (string, error)
	AppendBytes(content []byte) error
	AppendString(content string) error
	Rewrite(content string) error
	WriteAt(position int64, content []byte) error
	WriteStringAt(src string, position int64, content string) error
	WriteStringAtLine(row int, content string, replace bool) error
	Truncate(length int64)
}

// NewFile  创建一个 File 对象
func NewFile(path string) *File {
	rows, err := ReadByRow(path)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &File{path, len(rows)}
}

// IsDir 判断是文件夹或文件
func (f *File) IsDir() bool {
	return IsDir(f.path)
}

// Path 获取文件路径
func (f *File) Path() string {
	return f.path
}

// Create 创建文件
func (f *File) Create() error {
	return CreateFile(f.path)
}

// Copy 复制文件
func (f *File) Copy(target string) error {
	return CopyFileSafe(f.path, target)
}

// Clear 清空文件
func (f *File) Clear() error {
	return Clear(f.path)
}

// Remove 删除文件
func (f *File) Remove() error {
	return RemoveAll(f.path)
}

// Move 移动文件
func (f *File) Move(target string) error {
	return MoveSafe(f.path, target)
}

// Rename 重命名文件
func (f *File) Rename(newName string) error {
	return Rename(f.path, newName)
}

// Exist 判断文件是否存在
func (f *File) Exist() bool {
	return Exist(f.path)
}

// ReadBytes 读取文件返回字符串
func (f *File) ReadBytes() ([]byte, error) {
	return ReadBytes(f.path)
}

// ReadString 读取文件返回字符串
func (f *File) ReadString() (string, error) {
	return ReadString(f.path)
}

// ReadByRow 按行读取文件内容
func (f *File) ReadByRow() ([]string, error) {
	return ReadByRow(f.path)
}

// ReadByBytes 按块方式读取文件内容
func (f *File) ReadByBytes(bufferSize int64) ([][]byte, error) {
	return ReadByBytes(f.path, bufferSize)
}

// ReadAt 读取指定位置之后内容
//
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) ReadAt(position int64) ([]byte, error) {
	return ReadAt(f.path, position)
}

// ReadStringAt 读取指定位置之后内容
// 返回字符串
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) ReadStringAt(position int64) (string, error) {
	return ReadStringAt(f.path, position)
}

// AppendBytes 追加文件内容
func (f *File) AppendBytes(content []byte) error {
	return AppendBytes(f.path, content)
}

// AppendString 追加文件内容
func (f *File) AppendString(content string) error {
	return AppendString(f.path, content)
}

// Rewrite 重写文件内容
func (f *File) Rewrite(content string) error {
	return Rewrite(f.path, content)
}

// WriteAt 在文件指定位置写入内容
//
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) WriteAt(position int64, content []byte) error {
	return WriteAt(f.path, position, content)
}

// WriteStringAt 在文件指定位置写入内容
// 传入字符串数据即可
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) WriteStringAt(src string, position int64, content string) error {
	return WriteStringAt(f.path, position, content)
}

// WriteStringAtLine 在文件指定行写入内容
// 传入字符串数据即可
func (f *File) WriteStringAtLine(row int, content string, replace bool) error {
	return WriteStringAtLine(f.path, row, content, replace)
}

// Truncate 截短文件内容，使文件为指定长度
// 传入 0 则清空文件
// 传入负数则截短该数值长度
func (f *File) Truncate(length int64) error {
	return Truncate(f.path, length)
}

// Join 以此目录为基础目录，拼接完整路径
func (f *File) Join(folder string) string {
	return path.Join(folder, f.path)
}

// JoinBase 以此目录为基础目录，只拼接一级路径
func (f *File) JoinBase(folder string) string {
	return path.Join(folder, path.Base(f.path))
}

// Dir 目录
type Dir struct {
	path string
}

// DirKind 文件夹方法接口
type DirKind interface {
	Common
	Info() ([]os.FileInfo, error)
	Sub() ([]string, error)
	Subfolder() ([]string, error)
	Subfile() ([]string, error)
	All() ([]string, error)
	RemoveContent() error
	RemoveExt(ext string) error
	RemoveNamesRegexp(pattern string) error
}

// NewDir  创建一个 Dir 对象
func NewDir(path string) *Dir {
	return &Dir{path}
}

// Path 获取文件路径
func (d *Dir) Path() string {
	return d.path
}

// IsDir 判断是文件夹或文件
func (d *Dir) IsDir() bool {
	return IsDir(d.path)
}

// Create 创建文件夹
func (d *Dir) Create() error {
	return CreateDir(d.path)
}

// Info 原生方法读取目录信息
func (d *Dir) Info() ([]os.FileInfo, error) {

	return DirInfo(d.path)
}

// Exist 判断文件夹是否存在
func (d *Dir) Exist() bool {
	return Exist(d.path)
}

// Sub 读取目录下所有文件和目录信息
func (d *Dir) Sub() ([]string, error) {
	return DirSub(d.path)
}

// Subfolder 读取目录下的子文件夹
// 即不包含子文件
func (d *Dir) Subfolder() ([]string, error) {
	return DirSubfolder(d.path)
}

// Subfile 读取目录下的所有子文件
// 即不包含子目录
func (d *Dir) Subfile() ([]string, error) {
	return DirSubfile(d.path)
}

// All 读取目录和子目录下的所有文件名信息
//
// 全部为文件名，不再包含目录名
func (d *Dir) All() ([]string, error) {
	return DirAll(d.path)
}

// Copy 复制文件夹
func (d *Dir) Copy(target string) error {
	return CopyDir(d.path, target)
}

// Remove 删除文件夹
func (d *Dir) Remove() error {
	return RemoveAll(d.path)
}

// RemoveContent 删除目录全部内容，但不包含目录自身
func (d *Dir) RemoveContent() error {
	return RemoveDirContent(d.path)
}

// RemoveExt 删除目录下指定后缀名的文件
func (d *Dir) RemoveExt(ext string) error {
	return RemoveExt(d.path, ext)
}

// RemoveNamesRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件
func (d *Dir) RemoveNamesRegexp(pattern string) error {
	return RemoveNamesRegexp(d.path, pattern)
}

// RemoveNamesContain 目录下所有文件中，删除文件名包含给定字符的文件
func (d *Dir) RemoveNamesContain(containStr string) error {
	return RemoveNamesContain(d.path, containStr)
}

// Move 移动文件夹
func (d *Dir) Move(target string) error {
	return MoveSafe(d.path, target)
}

// Rename 重命名文件夹
func (d *Dir) Rename(newName string) error {
	return Rename(d.path, newName)
}

// Join 以此目录为基础目录，拼接完整路径
func (d *Dir) Join(src string) string {
	return path.Join(d.path, src)
}

// JoinBase 以此目录为基础目录，只拼接一级路径
func (d *Dir) JoinBase(src string) string {
	return path.Join(d.path, path.Base(src))
}
