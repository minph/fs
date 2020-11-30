package fs

import (
	"fmt"
	"io"
	"os"
)

func openFile(src string) (*os.File, error) {
	//打开文件
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func appendFile(src string) (*os.File, error) {
	//打开文件
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Rewrite 对文件覆盖写入数据
func Rewrite(src, content string) error {
	//打开文件
	file, err := openFile(src)
	defer file.Close()

	if err != nil {
		return err
	}
	// 直接写入字符串
	_, err = io.WriteString(file, content)
	return err
}

// AppendBytes 追加文件内容
func AppendBytes(src string, content []byte) error {

	// 追加方式打开文件
	file, err := appendFile(src)
	defer file.Close()

	if err != nil {
		return err
	}

	// 写入数据
	_, err = file.Write(content)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// AppendString 以字符串方式追加文件内容
func AppendString(src, content string) error {
	return AppendBytes(src, []byte(content))
}

// WriteAt 在文件指定位置写入内容
//
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func WriteAt(src string, position int64, content []byte) error {
	// 打开文件
	file, err := readFile(src)
	defer file.Close()
	if err != nil {
		return err
	}

	// 获取文件字节数
	bytes, err := ReadBytes(src)
	if err != nil {
		return err
	}

	// 获取定位
	if position >= 0 {
		file.Seek(position, 0)
	} else {
		file.Seek(int64(len(bytes))+position, 0)
	}

	// 写入内容
	_, err = file.Write(content)
	if err != nil {
		return err
	}

	return nil
}

// WriteStringAt 在文件指定位置写入内容
// 传入字符串数据即可
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func WriteStringAt(src string, position int64, content string) error {
	// 打开文件
	file, err := readFile(src)
	defer file.Close()
	if err != nil {
		return err
	}

	// 获取文件字节数
	bytes, err := ReadBytes(src)
	if err != nil {
		return err
	}

	// 获取定位
	if position < 0 {
		position = int64(len(bytes)) + position
	}

	// 写入内容
	_, err = file.WriteAt([]byte(content), position)
	if err != nil {
		return err
	}

	return nil
}
