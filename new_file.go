package fs

// File 文件
type File struct {
	Path string
}

// NewFile  创建一个 File 对象
func NewFile(path string) *File {
	return &File{path}
}

// Create 创建文件
func (f *File) Create() error {
	return CreateFile(f.Path)
}

// Copy 复制文件
func (f *File) Copy(target string) error {
	return CopyFileSafe(f.Path, target)
}

// Remove 删除文件
func (f *File) Remove() error {
	return RemoveAll(f.Path)
}

// Move 移动文件
func (f *File) Move(target string) error {
	return MoveSafe(f.Path, target)
}

// Rename 重命名文件
func (f *File) Rename(target string) error {
	return MoveSafe(f.Path, target)
}

// Exist 判断文件是否存在
func (f *File) Exist() bool {
	return Exist(f.Path)
}

// ReadBytes 读取文件返回字符串
func (f *File) ReadBytes() ([]byte, error) {
	return ReadBytes(f.Path)
}

// ReadString 读取文件返回字符串
func (f *File) ReadString() (string, error) {
	return ReadString(f.Path)
}

// ReadByRow 按行读取文件内容
func (f *File) ReadByRow() ([]string, error) {
	return ReadByRow(f.Path)
}

// ReadByBytes 按块方式读取文件内容
func (f *File) ReadByBytes(bufferSize int64) ([][]byte, error) {
	return ReadByBytes(f.Path, bufferSize)
}

// ReadAt 读取指定位置之后内容
//
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) ReadAt(position int64) ([]byte, error) {
	return ReadAt(f.Path, position)
}

// ReadStringAt 读取指定位置之后内容
// 返回字符串
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) ReadStringAt(position int64) (string, error) {
	return ReadStringAt(f.Path, position)
}

// AppendBytes 追加文件内容
func (f *File) AppendBytes(content []byte) error {
	return AppendBytes(f.Path, content)
}

// AppendString 追加文件内容
func (f *File) AppendString(content string) error {
	return AppendString(f.Path, content)
}

// Rewrite 重写文件内容
func (f *File) Rewrite(content string) error {
	return Rewrite(f.Path, content)
}

// WriteAt 在文件指定位置写入内容
//
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) WriteAt(position int64, content []byte) error {
	return WriteAt(f.Path, position, content)
}

// WriteStringAt 在文件指定位置写入内容
// 传入字符串数据即可
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func (f *File) WriteStringAt(src string, position int64, content string) error {
	return WriteStringAt(f.Path, position, content)
}

// Truncate 截短文件内容，使文件为指定长度
// 传入 0 则清空文件
// 传入负数则截短该数值长度
func (f *File) Truncate(length int64) error {
	return Truncate(f.Path, length)
}
