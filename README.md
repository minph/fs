# fs

一个全面的 golang 文件操作库

`import "github.com/minph/fs"`

# 概览

- [func AppendBytes(src string, content []byte) error](#AppendBytes)
- [func AppendString(src, content string) error](#AppendString)
- [func CopyDir(folder, target string) error](#CopyDir)
- [func CopyFile(src, target string) error](#CopyFile)
- [func CopyFileSafe(src, target string) error](#CopyFileSafe)
- [func CreateDir(folder string) error](#CreateDir)
- [func CreateFile(src string) error](#CreateFile)
- [func DirAll(folder string) ([]string, error)](#DirAll)
- [func DirInfo(folder string) ([]os.FileInfo, error)](#DirInfo)
- [func DirSub(folder string) ([]string, error)](#DirSub)
- [func DirSubfile(folder string) ([]string, error)](#DirSubfile)
- [func DirSubfolder(folder string) ([]string, error)](#DirSubfolder)
- [func Exist(path string) bool](#Exist)
- [func IsDir(path string) bool](#IsDir)
- [func Move(src, target string) error](#Move)
- [func MoveSafe(src, target string) error](#MoveSafe)
- [func ReadAt(src string, position int64) ([]byte, error)](#ReadAt)
- [func ReadByBytes(src string, bufferSize int64) ([][]byte, error)](#ReadByBytes)
- [func ReadByRow(src string) ([]string, error)](#ReadByRow)
- [func ReadBytes(src string) ([]byte, error)](#ReadBytes)
- [func ReadString(src string) (string, error)](#ReadString)
- [func ReadStringAt(src string, position int64) (string, error)](#ReadStringAt)
- [func Remove(src string) error](#Remove)
- [func RemoveAll(src string) error](#RemoveAll)
- [func RemoveDirContent(folder string) error](#RemoveDirContent)
- [func RemoveExt(folder, ext string) error](#RemoveExt)
- [func RemoveNamesContain(folder, containStr string) error](#RemoveNamesContain)
- [func RemoveNamesRegexp(folder, pattern string) error](#RemoveNamesRegexp)
- [func Rename(src, target string) error](#Rename)
- [func Rewrite(src, content string) error](#Rewrite)
- [func Truncate(src string, length int64) error](#Truncate)
- [func WriteAt(src string, position int64, content []byte) error](#WriteAt)
- [func WriteStringAt(src string, position int64, content string) error](#WriteStringAt)
- [type Common](#Common)
  - [func New(path string) Common](#New)
- [type Dir](#Dir)
  - [func NewDir(path string) \*Dir](#NewDir)
  - [func (d \*Dir) All() ([]string, error)](#Dir.All)
  - [func (d \*Dir) Copy(target string) error](#Dir.Copy)
  - [func (d \*Dir) Create() error](#Dir.Create)
  - [func (d \*Dir) Exist() bool](#Dir.Exist)
  - [func (d \*Dir) Info() ([]os.FileInfo, error)](#Dir.Info)
  - [func (d \*Dir) IsDir() bool](#Dir.IsDir)
  - [func (d \*Dir) Move(target string) error](#Dir.Move)
  - [func (d \*Dir) Path() string](#Dir.Path)
  - [func (d \*Dir) Remove() error](#Dir.Remove)
  - [func (d \*Dir) RemoveContent() error](#Dir.RemoveContent)
  - [func (d \*Dir) RemoveExt(ext string) error](#Dir.RemoveExt)
  - [func (d \*Dir) RemoveNamesContain(containStr string) error](#Dir.RemoveNamesContain)
  - [func (d \*Dir) RemoveNamesRegexp(pattern string) error](#Dir.RemoveNamesRegexp)
  - [func (d \*Dir) Rename(target string) error](#Dir.Rename)
  - [func (d \*Dir) Sub() ([]string, error)](#Dir.Sub)
  - [func (d \*Dir) Subfile() ([]string, error)](#Dir.Subfile)
  - [func (d \*Dir) Subfolder() ([]string, error)](#Dir.Subfolder)
- [type DirKind](#DirKind)
- [type File](#File)
  - [func NewFile(path string) \*File](#NewFile)
  - [func (f \*File) AppendBytes(content []byte) error](#File.AppendBytes)
  - [func (f \*File) AppendString(content string) error](#File.AppendString)
  - [func (f \*File) Copy(target string) error](#File.Copy)
  - [func (f \*File) Create() error](#File.Create)
  - [func (f \*File) Exist() bool](#File.Exist)
  - [func (f \*File) IsDir() bool](#File.IsDir)
  - [func (f \*File) Move(target string) error](#File.Move)
  - [func (f \*File) Path() string](#File.Path)
  - [func (f \*File) ReadAt(position int64) ([]byte, error)](#File.ReadAt)
  - [func (f \*File) ReadByBytes(bufferSize int64) ([][]byte, error)](#File.ReadByBytes)
  - [func (f \*File) ReadByRow() ([]string, error)](#File.ReadByRow)
  - [func (f \*File) ReadBytes() ([]byte, error)](#File.ReadBytes)
  - [func (f \*File) ReadString() (string, error)](#File.ReadString)
  - [func (f \*File) ReadStringAt(position int64) (string, error)](#File.ReadStringAt)
  - [func (f \*File) Remove() error](#File.Remove)
  - [func (f \*File) Rename(target string) error](#File.Rename)
  - [func (f \*File) Rewrite(content string) error](#File.Rewrite)
  - [func (f \*File) Truncate(length int64) error](#File.Truncate)
  - [func (f \*File) WriteAt(position int64, content []byte) error](#File.WriteAt)
  - [func (f \*File) WriteStringAt(src string, position int64, content string) error](#File.WriteStringAt)
- [type FileKind](#FileKind)

# 详情

```go
func AppendBytes(src string, content []byte) error
```

AppendBytes 追加文件内容

## <a name="AppendString">func</a> AppendString

```go
func AppendString(src, content string) error
```

AppendString 以字符串方式追加文件内容

## <a name="CopyDir">func</a> CopyDir

```go
func CopyDir(folder, target string) error
```

CopyDir 复制文件夹

## <a name="CopyFile">func</a> CopyFile

```go
func CopyFile(src, target string) error
```

CopyFile 复制文件

## <a name="CopyFileSafe">func</a> CopyFileSafe

```go
func CopyFileSafe(src, target string) error
```

CopyFileSafe 安全地复制文件

保证复制位置文件存在

## <a name="CreateDir">func</a> CreateDir

```go
func CreateDir(folder string) error
```

CreateDir 跨目录创建文件夹

## <a name="CreateFile">func</a> CreateFile

```go
func CreateFile(src string) error
```

CreateFile 跨目录创建文件

## <a name="DirAll">func</a> DirAll

```go
func DirAll(folder string) ([]string, error)
```

DirAll 读取目录和子目录下的所有文件名信息

全部为文件名，不再包含目录名

## <a name="DirInfo">func</a> DirInfo

```go
func DirInfo(folder string) ([]os.FileInfo, error)
```

DirInfo 原生方法读取目录信息

## <a name="DirSub">func</a> DirSub

```go
func DirSub(folder string) ([]string, error)
```

DirSub 读取目录下所有文件和目录信息

## <a name="DirSubfile">func</a> DirSubfile

```go
func DirSubfile(folder string) ([]string, error)
```

DirSubfile 读取目录下的所有子文件
即不包含子目录

## <a name="DirSubfolder">func</a> DirSubfolder

```go
func DirSubfolder(folder string) ([]string, error)
```

DirSubfolder 读取目录下的子文件夹
即不包含子文件

## <a name="Exist">func</a> Exist

```go
func Exist(path string) bool
```

Exist 判断所给路径文件或文件夹是否存在

## <a name="IsDir">func</a> IsDir

```go
func IsDir(path string) bool
```

IsDir 判断是不是文件夹

## <a name="Move">func</a> Move

```go
func Move(src, target string) error
```

Move 移动文件或文件夹

和 Rename 同义，函数名不一致是为了语义的区别

## <a name="MoveSafe">func</a> MoveSafe

```go
func MoveSafe(src, target string) error
```

MoveSafe 安全地移动文件或文件夹
保证跨目录的安全性

## <a name="ReadAt">func</a> ReadAt

```go
func ReadAt(src string, position int64) ([]byte, error)
```

ReadAt 读取指定位置之后内容

正数则从开始到最后定位
负数则从最后到开始定位

## <a name="ReadByBytes">func</a> ReadByBytes

```go
func ReadByBytes(src string, bufferSize int64) ([][]byte, error)
```

ReadByBytes 按块方式读取文件内容

## <a name="ReadByRow">func</a> ReadByRow

```go
func ReadByRow(src string) ([]string, error)
```

ReadByRow 按行读取文件内容

## <a name="ReadBytes">func</a> ReadBytes

```go
func ReadBytes(src string) ([]byte, error)
```

ReadBytes 读取文件全部内容

## <a name="ReadString">func</a> ReadString

```go
func ReadString(src string) (string, error)
```

ReadString 读取文件全部内容并返回字符串

## <a name="ReadStringAt">func</a> ReadStringAt

```go
func ReadStringAt(src string, position int64) (string, error)
```

ReadStringAt 读取指定位置之后内容
返回字符串
正数则从开始到最后定位
负数则从最后到开始定位

## <a name="Remove">func</a> Remove

```go
func Remove(src string) error
```

Remove 同 os.Remove

只用于删除单个文件和空目录

## <a name="RemoveAll">func</a> RemoveAll

```go
func RemoveAll(src string) error
```

RemoveAll 同 os.RemoveAll

## <a name="RemoveDirContent">func</a> RemoveDirContent

```go
func RemoveDirContent(folder string) error
```

RemoveDirContent 删除目录全部内容，但不包含目录自身

## <a name="RemoveExt">func</a> RemoveExt

```go
func RemoveExt(folder, ext string) error
```

RemoveExt 删除目录下指定后缀名的文件

## <a name="RemoveNamesContain">func</a> RemoveNamesContain

```go
func RemoveNamesContain(folder, containStr string) error
```

RemoveNamesContain 目录下所有文件中，删除文件名包含给定字符的文件

## <a name="RemoveNamesRegexp">func</a> RemoveNamesRegexp

```go
func RemoveNamesRegexp(folder, pattern string) error
```

RemoveNamesRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件

## <a name="Rename">func</a> Rename

```go
func Rename(src, target string) error
```

Rename 重命名文件

## <a name="Rewrite">func</a> Rewrite

```go
func Rewrite(src, content string) error
```

Rewrite 对文件覆盖写入数据

## <a name="Truncate">func</a> Truncate

```go
func Truncate(src string, length int64) error
```

Truncate 截短文件内容，使文件为指定长度
传入 0 则清空文件
传入负数则截短该数值长度

## <a name="WriteAt">func</a> WriteAt

```go
func WriteAt(src string, position int64, content []byte) error
```

WriteAt 在文件指定位置写入内容

正数则从开始到最后定位
负数则从最后到开始定位

## <a name="WriteStringAt">func</a> WriteStringAt

```go
func WriteStringAt(src string, position int64, content string) error
```

WriteStringAt 在文件指定位置写入内容
传入字符串数据即可
正数则从开始到最后定位
负数则从最后到开始定位

## <a name="Common">type</a> Common

```go
type Common interface {
    Create() error
    Copy(target string) error
    Remove() error
    Move(target string) error
    Rename(target string) error
    Exist() bool
    IsDir() bool
    Path() string
}
```

Common 文件夹和文件公用方法接口

### <a name="New">func</a> New

```go
func New(path string) Common
```

New 在不判断文件夹还是文件的情况下，支持通用方法接口

## <a name="Dir">type</a> Dir

```go
type Dir struct {
    // contains filtered or unexported fields
}

```

Dir 目录

### <a name="NewDir">func</a> NewDir

```go
func NewDir(path string) *Dir
```

NewDir 创建一个 Dir 对象

### <a name="Dir.All">func</a> (\*Dir) All

```go
func (d *Dir) All() ([]string, error)
```

All 读取目录和子目录下的所有文件名信息

全部为文件名，不再包含目录名

### <a name="Dir.Copy">func</a> (\*Dir) Copy

```go
func (d *Dir) Copy(target string) error
```

Copy 复制文件夹

### <a name="Dir.Create">func</a> (\*Dir) Create

```go
func (d *Dir) Create() error
```

Create 创建文件夹

### <a name="Dir.Exist">func</a> (\*Dir) Exist

```go
func (d *Dir) Exist() bool
```

Exist 判断文件夹是否存在

### <a name="Dir.Info">func</a> (\*Dir) Info

```go
func (d *Dir) Info() ([]os.FileInfo, error)
```

Info 原生方法读取目录信息

### <a name="Dir.IsDir">func</a> (\*Dir) IsDir

```go
func (d *Dir) IsDir() bool
```

IsDir 判断是文件夹或文件

### <a name="Dir.Move">func</a> (\*Dir) Move

```go
func (d *Dir) Move(target string) error
```

Move 移动文件夹

### <a name="Dir.Path">func</a> (\*Dir) Path

```go
func (d *Dir) Path() string
```

Path 获取文件路径

### <a name="Dir.Remove">func</a> (\*Dir) Remove

```go
func (d *Dir) Remove() error
```

Remove 删除文件夹

### <a name="Dir.RemoveContent">func</a> (\*Dir) RemoveContent

```go
func (d *Dir) RemoveContent() error
```

RemoveContent 删除目录全部内容，但不包含目录自身

### <a name="Dir.RemoveExt">func</a> (\*Dir) RemoveExt

```go
func (d *Dir) RemoveExt(ext string) error
```

RemoveExt 删除目录下指定后缀名的文件

### <a name="Dir.RemoveNamesContain">func</a> (\*Dir) RemoveNamesContain

```go
func (d *Dir) RemoveNamesContain(containStr string) error
```

RemoveNamesContain 目录下所有文件中，删除文件名包含给定字符的文件

### <a name="Dir.RemoveNamesRegexp">func</a> (\*Dir) RemoveNamesRegexp

```go
func (d *Dir) RemoveNamesRegexp(pattern string) error
```

RemoveNamesRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件

### <a name="Dir.Rename">func</a> (\*Dir) Rename

```go
func (d *Dir) Rename(target string) error
```

Rename 重命名文件夹

### <a name="Dir.Sub">func</a> (\*Dir) Sub

```go
func (d *Dir) Sub() ([]string, error)
```

Sub 读取目录下所有文件和目录信息

### <a name="Dir.Subfile">func</a> (\*Dir) Subfile

```go
func (d *Dir) Subfile() ([]string, error)
```

Subfile 读取目录下的所有子文件
即不包含子目录

### <a name="Dir.Subfolder">func</a> (\*Dir) Subfolder

```go
func (d *Dir) Subfolder() ([]string, error)
```

Subfolder 读取目录下的子文件夹
即不包含子文件

## <a name="DirKind">type</a> DirKind

```go
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
```

DirKind 文件夹方法接口

## <a name="File">type</a> File

```go
type File struct {
    // contains filtered or unexported fields
}

```

File 文件

### <a name="NewFile">func</a> NewFile

```go
func NewFile(path string) *File
```

NewFile 创建一个 File 对象

### <a name="File.AppendBytes">func</a> (\*File) AppendBytes

```go
func (f *File) AppendBytes(content []byte) error
```

AppendBytes 追加文件内容

### <a name="File.AppendString">func</a> (\*File) AppendString

```go
func (f *File) AppendString(content string) error
```

AppendString 追加文件内容

### <a name="File.Copy">func</a> (\*File) Copy

```go
func (f *File) Copy(target string) error
```

Copy 复制文件

### <a name="File.Create">func</a> (\*File) Create

```go
func (f *File) Create() error
```

Create 创建文件

### <a name="File.Exist">func</a> (\*File) Exist

```go
func (f *File) Exist() bool
```

Exist 判断文件是否存在

### <a name="File.IsDir">func</a> (\*File) IsDir

```go
func (f *File) IsDir() bool
```

IsDir 判断是文件夹或文件

### <a name="File.Move">func</a> (\*File) Move

```go
func (f *File) Move(target string) error
```

Move 移动文件

### <a name="File.Path">func</a> (\*File) Path

```go
func (f *File) Path() string
```

Path 获取文件路径

### <a name="File.ReadAt">func</a> (\*File) ReadAt

```go
func (f *File) ReadAt(position int64) ([]byte, error)
```

ReadAt 读取指定位置之后内容

正数则从开始到最后定位
负数则从最后到开始定位

### <a name="File.ReadByBytes">func</a> (\*File) ReadByBytes

```go
func (f *File) ReadByBytes(bufferSize int64) ([][]byte, error)
```

ReadByBytes 按块方式读取文件内容

### <a name="File.ReadByRow">func</a> (\*File) ReadByRow

```go
func (f *File) ReadByRow() ([]string, error)
```

ReadByRow 按行读取文件内容

### <a name="File.ReadBytes">func</a> (\*File) ReadBytes

```go
func (f *File) ReadBytes() ([]byte, error)
```

ReadBytes 读取文件返回字符串

### <a name="File.ReadString">func</a> (\*File) ReadString

```go
func (f *File) ReadString() (string, error)
```

ReadString 读取文件返回字符串

### <a name="File.ReadStringAt">func</a> (\*File) ReadStringAt

```go
func (f *File) ReadStringAt(position int64) (string, error)
```

ReadStringAt 读取指定位置之后内容
返回字符串
正数则从开始到最后定位
负数则从最后到开始定位

### <a name="File.Remove">func</a> (\*File) Remove

```go
func (f *File) Remove() error
```

Remove 删除文件

### <a name="File.Rename">func</a> (\*File) Rename

```go
func (f *File) Rename(target string) error
```

Rename 重命名文件

### <a name="File.Rewrite">func</a> (\*File) Rewrite

```go
func (f *File) Rewrite(content string) error
```

Rewrite 重写文件内容

### <a name="File.Truncate">func</a> (\*File) Truncate

```go
func (f *File) Truncate(length int64) error
```

Truncate 截短文件内容，使文件为指定长度
传入 0 则清空文件
传入负数则截短该数值长度

### <a name="File.WriteAt">func</a> (\*File) WriteAt

```go
func (f *File) WriteAt(position int64, content []byte) error
```

WriteAt 在文件指定位置写入内容

正数则从开始到最后定位
负数则从最后到开始定位

### <a name="File.WriteStringAt">func</a> (\*File) WriteStringAt

```go
func (f *File) WriteStringAt(src string, position int64, content string) error
```

WriteStringAt 在文件指定位置写入内容
传入字符串数据即可
正数则从开始到最后定位
负数则从最后到开始定位

## <a name="FileKind">type</a> FileKind

```go
type FileKind interface {
    Common

    // ReadBytes 读取文件返回字符串
    ReadBytes() ([]byte, error)

    // ReadString 读取文件返回字符串
    ReadString() (string, error)

    // ReadByRow 按行读取文件内容
    ReadByRow() ([]string, error)

    // ReadByBytes 按块方式读取文件内容
    ReadByBytes(bufferSize int64) ([][]byte, error)

    // ReadAt 读取指定位置之后内容
    //
    // 正数则从开始到最后定位
    // 负数则从最后到开始定位
    ReadAt(position int64) ([]byte, error)

    // ReadStringAt 读取指定位置之后内容
    // 返回字符串
    // 正数则从开始到最后定位
    // 负数则从最后到开始定位
    ReadStringAt(position int64) (string, error)

    // AppendBytes 追加文件内容
    AppendBytes(content []byte) error

    // AppendString 追加文件内容
    AppendString(content string) error

    // Rewrite 重写文件内容
    Rewrite(content string) error

    // WriteAt 在文件指定位置写入内容
    //
    // 正数则从开始到最后定位
    // 负数则从最后到开始定位
    WriteAt(position int64, content []byte) error

    // WriteStringAt 在文件指定位置写入内容
    // 传入字符串数据即可
    // 正数则从开始到最后定位
    // 负数则从最后到开始定位
    WriteStringAt(src string, position int64, content string) error

    // Truncate 截短文件内容，使文件为指定长度
    // 传入 0 则清空文件
    // 传入负数则截短该数值长度
    Truncate(length int64)
}
```

FileKind 文件方法接口
