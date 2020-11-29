package fs

import (
	"fmt"
	"io"
	"os"
)

// Rewrite 对文件覆盖写入数据
func Rewrite(src, content string) error {
	//打开文件
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return err
	}
	// 直接写入字符串
	_, err = io.WriteString(file, content)
	return err
}

// AppendByte 追加文件内容
func AppendByte(src string, content []byte) error {

	// 追加方式打开文件
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_APPEND|os.O_CREATE, os.ModePerm)
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
	return AppendByte(src, []byte(content))
}
