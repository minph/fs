# fs

一个全面的 Go 语言文件操作 package，API 参照 nodejs 中 fs-extra 的设计，简单易用

`import "github.com/minph/fs"`

# 概览

- [func Append(src, content string) error](#Append)
- [func Clear(src string) error](#Clear)
- [func Copy(src, target string) error](#Copy)
- [func CopySafe(src, target string) error](#CopySafe)
- [func Create(src string) error](#Create)
- [func Move(src, target string) error](#Move)
- [func MoveSafe(src, target string) error](#MoveSafe)
- [func ReadAll(src string) (string, error)](#ReadAll)
- [func ReadBytes(src string, bufferSize int64) ([][]byte, error)](#ReadBytes)
- [func ReadLineAt(src string, row int) (string, error)](#ReadLineAt)
- [func ReadLines(src string) ([]string, error)](#ReadLines)
- [func Remove(src string) error](#Remove)
- [func RemoveAll(src string) error](#RemoveAll)
- [func Rename(src, newName string) error](#Rename)
- [func Rewrite(src, content string) error](#Rewrite)
- [func Truncate(src string, length int64) error](#Truncate)
- [func WriteLine(src string, content string, row int, replace bool) error](#WriteLine)
- [type File](#File)
  - [func New(src string) \*File](#New)
  - [func (f *File) Append(content string) *File](#File.Append)
  - [func (f *File) Check(call func(err []error)) *File](#File.Check)
  - [func (f *File) Clear() *File](#File.Clear)
  - [func (f *File) Copy(target string) *File](#File.Copy)
  - [func (f *File) CopySafe(target string) *File](#File.CopySafe)
  - [func (f *File) Create() *File](#File.Create)
  - [func (f *File) Do(call func(f *File)) \*File](#File.Do)
  - [func (f *File) Move(target string) *File](#File.Move)
  - [func (f *File) ReadAll() *File](#File.ReadAll)
  - [func (f *File) ReadBytes(bufferSize int64) *File](#File.ReadBytes)
  - [func (f *File) ReadLineAt(row int) *File](#File.ReadLineAt)
  - [func (f *File) ReadLines() *File](#File.ReadLines)
  - [func (f *File) Rename(name string) *File](#File.Rename)
  - [func (f *File) Rewrite(content string) *File](#File.Rewrite)
  - [func (f *File) Truncate(length int64) *File](#File.Truncate)
  - [func (f *File) WriteLine(content string, row int, replace bool) *File](#File.WriteLine)

## <a name="Append">func</a> Append

```go
func Append(src, content string) error
```

Append 以字符串方式追加文件内容

## <a name="Clear">func</a> Clear

```go
func Clear(src string) error
```

Clear 清空文件内容

## <a name="Copy">func</a> Copy

```go
func Copy(src, target string) error
```

Copy 复制文件

## <a name="CopySafe">func</a> CopySafe

```go
func CopySafe(src, target string) error
```

CopySafe 安全地复制文件

保证复制位置文件存在

## <a name="Create">func</a> Create

```go
func Create(src string) error
```

Create 跨目录创建文件

## <a name="Move">func</a> Move

```go
func Move(src, target string) error
```

Move 移动文件或文件夹

## <a name="MoveSafe">func</a> MoveSafe

```go
func MoveSafe(src, target string) error
```

MoveSafe 安全地移动文件或文件夹
保证跨目录的安全性

## <a name="ReadAll">func</a> ReadAll

```go
func ReadAll(src string) (string, error)
```

ReadAll 读取文件全部内容并返回字符串

## <a name="ReadBytes">func</a> ReadBytes

```go
func ReadBytes(src string, bufferSize int64) ([][]byte, error)
```

ReadBytes 按 byte 块方式读取文件内容

## <a name="ReadLineAt">func</a> ReadLineAt

```go
func ReadLineAt(src string, row int) (string, error)
```

ReadLineAt 读取指定行的内容

## <a name="ReadLines">func</a> ReadLines

```go
func ReadLines(src string) ([]string, error)
```

ReadLines 按行读取文件内容

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

## <a name="Rename">func</a> Rename

```go
func Rename(src, newName string) error
```

Rename 重命名文件
不修改目录路径，只修改文件名

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

## <a name="WriteLine">func</a> WriteLine

```go
func WriteLine(src string, content string, row int, replace bool) error
```

WriteLine 在文件指定行写入内容
传入字符串数据即可
row 修改的行数
replace 是否替换当前行

## <a name="File">type</a> File

```go
type File struct {
    utils.Wrapper
    Path    string      //文件地址
    DirName string      // 文件目录名
    Res     interface{} // 函数返回值
}

```

File 文件类型结构体

### <a name="New">func</a> New

```go
func New(src string) *File
```

New 创建文件类型结构体

### <a name="File.Append">func</a> (\*File) Append

```go
func (f *File) Append(content string) *File
```

Append 以字符串方式追加文件内容

### <a name="File.Check">func</a> (\*File) Check

```go
func (f *File) Check(call func(err []error)) *File
```

Check 检查错误

### <a name="File.Clear">func</a> (\*File) Clear

```go
func (f *File) Clear() *File
```

Clear 清空文件内容

### <a name="File.Copy">func</a> (\*File) Copy

```go
func (f *File) Copy(target string) *File
```

Copy 复制文件

### <a name="File.CopySafe">func</a> (\*File) CopySafe

```go
func (f *File) CopySafe(target string) *File
```

CopySafe 安全地复制文件

### <a name="File.Create">func</a> (\*File) Create

```go
func (f *File) Create() *File
```

Create 创建文件

### <a name="File.Do">func</a> (\*File) Do

```go
func (f *File) Do(call func(f *File)) *File
```

Do 对结构体进行链式操作

### <a name="File.Move">func</a> (\*File) Move

```go
func (f *File) Move(target string) *File
```

Rename 重命名文件

### <a name="File.ReadAll">func</a> (\*File) ReadAll

```go
func (f *File) ReadAll() *File
```

ReadAll 读取文件全部内容并返回字符串

### <a name="File.ReadBytes">func</a> (\*File) ReadBytes

```go
func (f *File) ReadBytes(bufferSize int64) *File
```

ReadBytes 按 byte 块方式读取文件内容

### <a name="File.ReadLineAt">func</a> (\*File) ReadLineAt

```go
func (f *File) ReadLineAt(row int) *File
```

ReadLineAt 读取指定行的内容

### <a name="File.ReadLines">func</a> (\*File) ReadLines

```go
func (f *File) ReadLines() *File
```

ReadLines 按行读取文件内容

### <a name="File.Rename">func</a> (\*File) Rename

```go
func (f *File) Rename(name string) *File
```

Rename 重命名文件

### <a name="File.Rewrite">func</a> (\*File) Rewrite

```go
func (f *File) Rewrite(content string) *File
```

Rewrite 对文件覆盖写入数据

### <a name="File.Truncate">func</a> (\*File) Truncate

```go
func (f *File) Truncate(length int64) *File
```

Truncate 截短文件内容，使文件为指定长度
传入 0 则清空文件
传入负数则截短该数值长度

### <a name="File.WriteLine">func</a> (\*File) WriteLine

```go
func (f *File) WriteLine(content string, row int, replace bool) *File
```

WriteLine 在文件指定行写入内容
传入字符串数据即可
row 修改的行数
replace 是否替换当前行

# dir 目录操作包

`import "github.com/minph/fs/dir"`

## 概览

- [func All(folder string) ([]string, error)](#All)
- [func Content(folder string) ([]string, error)](#Content)
- [func Copy(folder, target string) error](#Copy)
- [func Create(folder string) error](#Create)
- [func Files(folder string) ([]string, error)](#Files)
- [func Folders(folder string) ([]string, error)](#Folders)
- [func Info(folder string) ([]os.FileInfo, error)](#Info)
- [func Is(path string) bool](#Is)
- [func Map(folder string, call func(index int, file string)) error](#Map)
- [func RemoveContain(folder, containStr string) error](#RemoveContain)
- [func RemoveContent(folder string) error](#RemoveContent)
- [func RemoveExt(folder, ext string) error](#RemoveExt)
- [func RemoveRegexp(folder, pattern string) error](#RemoveRegexp)
- [type Dir](#Dir)
  - [func New(src string) \*Dir](#New)
  - [func (d *Dir) All() *Dir](#Dir.All)
  - [func (d *Dir) Content() *Dir](#Dir.Content)
  - [func (d *Dir) Copy(target string) *Dir](#Dir.Copy)
  - [func (d *Dir) Create() *Dir](#Dir.Create)
  - [func (d *Dir) Files() *Dir](#Dir.Files)
  - [func (d *Dir) Folders() *Dir](#Dir.Folders)
  - [func (d *Dir) Info() *Dir](#Dir.Info)
  - [func (d *Dir) Is() *Dir](#Dir.Is)
  - [func (d *Dir) Map(call func(index int, file string)) *Dir](#Dir.Map)
  - [func (d *Dir) RemoveContain(containStr string) *Dir](#Dir.RemoveContain)
  - [func (d *Dir) RemoveContent() *Dir](#Dir.RemoveContent)
  - [func (d *Dir) RemoveExt(ext string) *Dir](#Dir.RemoveExt)
  - [func (d *Dir) RemoveRegexp(pattern string) *Dir](#Dir.RemoveRegexp)

## <a name="All">func</a> All

```go
func All(folder string) ([]string, error)
```

All 读取目录和子目录下的所有文件名信息

全部为文件名，不再包含目录名

## <a name="Content">func</a> Content

```go
func Content(folder string) ([]string, error)
```

Content 读取目录下所有文件和目录信息

## <a name="Copy">func</a> Copy

```go
func Copy(folder, target string) error
```

Copy 复制文件夹

## <a name="Create">func</a> Create

```go
func Create(folder string) error
```

Create 跨目录创建文件夹

## <a name="Files">func</a> Files

```go
func Files(folder string) ([]string, error)
```

Files 读取目录下的所有子文件
即不包含子目录

## <a name="Folders">func</a> Folders

```go
func Folders(folder string) ([]string, error)
```

Folders 读取目录下的子文件夹
即不包含子文件

## <a name="Info">func</a> Info

```go
func Info(folder string) ([]os.FileInfo, error)
```

Info 原生方法读取目录信息

## <a name="Is">func</a> Is

```go
func Is(path string) bool
```

Is 判断是不是文件夹

## <a name="Map">func</a> Map

```go
func Map(folder string, call func(index int, file string)) error
```

Map 读取目录和子目录下的所有文件名信息，然后进行操作

全部为文件名，不再包含目录名

## <a name="RemoveContain">func</a> RemoveContain

```go
func RemoveContain(folder, containStr string) error
```

RemoveContain 目录下所有文件中，删除文件名包含给定字符的文件

## <a name="RemoveContent">func</a> RemoveContent

```go
func RemoveContent(folder string) error
```

RemoveContent 删除目录全部内容，但不包含目录自身

## <a name="RemoveExt">func</a> RemoveExt

```go
func RemoveExt(folder, ext string) error
```

RemoveExt 删除目录下指定后缀名的文件

## <a name="RemoveRegexp">func</a> RemoveRegexp

```go
func RemoveRegexp(folder, pattern string) error
```

RemoveRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件

## <a name="Dir">type</a> Dir

```go
type Dir struct {
    utils.Wrapper
    Path string      //目录地址
    Res  interface{} // 函数返回值
}

```

Dir 目录类型结构体

### <a name="New">func</a> New

```go
func New(src string) *Dir
```

New 创建目录类型结构体

### <a name="Dir.All">func</a> (\*Dir) All

```go
func (d *Dir) All() *Dir
```

All 读取目录和子目录下的所有文件名信息

全部为文件名，不再包含目录名

### <a name="Dir.Content">func</a> (\*Dir) Content

```go
func (d *Dir) Content() *Dir
```

Content 读取目录下所有文件和目录信息

### <a name="Dir.Copy">func</a> (\*Dir) Copy

```go
func (d *Dir) Copy(target string) *Dir
```

Copy 复制文件夹

### <a name="Dir.Create">func</a> (\*Dir) Create

```go
func (d *Dir) Create() *Dir
```

Create 跨目录创建文件夹

### <a name="Dir.Files">func</a> (\*Dir) Files

```go
func (d *Dir) Files() *Dir
```

Files 读取目录下的所有子文件
即不包含子目录

### <a name="Dir.Folders">func</a> (\*Dir) Folders

```go
func (d *Dir) Folders() *Dir
```

Folders 读取目录下的子文件夹
即不包含子文件

### <a name="Dir.Info">func</a> (\*Dir) Info

```go
func (d *Dir) Info() *Dir
```

Info 原生方法读取目录信息

### <a name="Dir.Is">func</a> (\*Dir) Is

```go
func (d *Dir) Is() *Dir
```

Is 判断是不是文件夹

### <a name="Dir.Map">func</a> (\*Dir) Map

```go
func (d *Dir) Map(call func(index int, file string)) *Dir
```

Map 读取目录和子目录下的所有文件名信息，然后进行操作

全部为文件名，不再包含目录名

### <a name="Dir.RemoveContain">func</a> (\*Dir) RemoveContain

```go
func (d *Dir) RemoveContain(containStr string) *Dir
```

RemoveContain 目录下所有文件中，删除文件名包含给定字符的文件

### <a name="Dir.RemoveContent">func</a> (\*Dir) RemoveContent

```go
func (d *Dir) RemoveContent() *Dir
```

RemoveContent 删除目录全部内容，但不包含目录自身

### <a name="Dir.RemoveExt">func</a> (\*Dir) RemoveExt

```go
func (d *Dir) RemoveExt(ext string) *Dir
```

RemoveExt 删除目录下指定后缀名的文件

### <a name="Dir.RemoveRegexp">func</a> (\*Dir) RemoveRegexp

```go
func (d *Dir) RemoveRegexp(pattern string) *Dir
```

RemoveRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件
