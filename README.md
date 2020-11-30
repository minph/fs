# fs

golang 文件操作库

`import "github.com/minph/fs"`

# 概览

- [func AppendByte(src string, content []byte) error](#AppendByte)
- [func AppendString(src, content string) error](#AppendString)
- [func CopyDir(folder, target string) error](#CopyDir)
- [func CopyFile(src, target string) error](#CopyFile)
- [func CopyFileSafe(src, target string) error](#CopyFileSafe)
- [func CreateDir(folder string) error](#CreateDir)
- [func CreateFile(src string) error](#CreateFile)
- [func Dir(folder string) ([]os.FileInfo, error)](#Dir)
- [func DirAll(folder string) ([]string, error)](#DirAll)
- [func DirSub(folder string) ([]string, error)](#DirSub)
- [func DirSubfile(folder string) ([]string, error)](#DirSubfile)
- [func DirSubfolder(folder string) ([]string, error)](#DirSubfolder)
- [func Exist(path string) bool](#Exist)
- [func Move(src, target string) error](#Move)
- [func MoveSafe(src, target string) error](#MoveSafe)
- [func ReadAt(src string, position int64) ([]byte, error)](#ReadAt)
- [func ReadByBytes(src string, bufferSize int64) ([][]byte, error)](#ReadByBytes)
- [func ReadByRow(src string) ([]string, error)](#ReadByRow)
- [func ReadByte(src string) ([]byte, error)](#ReadByte)
- [func ReadString(src string) (string, error)](#ReadString)
- [func ReadStringAt(src string, position int64) (string, error)](#ReadStringAt)
- [func Remove(src string) error](#Remove)
- [func RemoveAll(src string) error](#RemoveAll)
- [func RemoveDir(folder string) error](#RemoveDir)
- [func RemoveExt(folder, ext string) error](#RemoveExt)
- [func RemoveNamesContain(folder, containStr string) error](#RemoveNamesContain)
- [func RemoveNamesRegexp(folder, pattern string) error](#RemoveNamesRegexp)
- [func Rename(src, target string) error](#Rename)
- [func Rewrite(src, content string) error](#Rewrite)
- [func Truncate(src string, length int64) error](#Truncate)
- [func WriteAt(src string, position int64, content []byte) error](#WriteAt)
- [func WriteStringAt(src string, position int64, content string) error](#WriteStringAt)

# 详情

## <a name="AppendByte">func</a> AppendByte

```go
func AppendByte(src string, content []byte) error
```

AppendByte 追加文件内容

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

## <a name="Dir">func</a> Dir

```go
func Dir(folder string) ([]os.FileInfo, error)
```

Dir 原生方法读取目录信息

## <a name="DirAll">func</a> DirAll

```go
func DirAll(folder string) ([]string, error)
```

DirAll 读取目录和子目录下的所有文件名信息

全部为文件名，不再包含目录名

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

## <a name="ReadByte">func</a> ReadByte

```go
func ReadByte(src string) ([]byte, error)
```

ReadByte 读取文件全部内容

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

## <a name="RemoveDir">func</a> RemoveDir

```go
func RemoveDir(folder string) error
```

RemoveDir 删除目录全部内容，但不包含目录自身

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
