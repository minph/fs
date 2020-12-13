package fs

import (
	"fmt"
	"io"
	"os"
)

func truncOpen(src string) (*os.File, error) {
	//打开文件
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func appendOpen(src string) (*os.File, error) {
	//打开文件
	file, err := os.OpenFile(src, os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return nil, err
	}

	return file, nil
}

// Truncate 截短文件内容，使文件为指定长度
// 传入 0 则清空文件
// 传入负数则截短该数值长度
func Truncate(src string, length int64) error {

	// 传入负数
	if length < 0 {
		content, err := readAll(src)
		if err != nil {
			return err
		}

		// 改为截短该长度
		// 切记 length 是负数，用加号
		length = int64(len(content)) + length
	}
	return os.Truncate(src, length)
}

// Rewrite 对文件覆盖写入数据
func Rewrite(src, content string) error {
	//打开文件
	file, err := truncOpen(src)
	defer file.Close()

	if err != nil {
		return err
	}
	// 直接写入字符串
	_, err = io.WriteString(file, content)
	return err
}

// Clear 清空文件内容
func Clear(src string) error {
	return Rewrite(src, "")
}

// Append 以字符串方式追加文件内容
func Append(src, content string) error {
	// 追加方式打开文件
	file, err := appendOpen(src)
	defer file.Close()

	if err != nil {
		return err
	}

	// 写入数据
	_, err = file.Write([]byte(content))

	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

// WriteLine 在文件指定行写入内容
// 传入字符串数据即可
// row 修改的行数
// replace 是否替换当前行
func WriteLine(src string, content string, row int, replace bool) error {
	// 打开文件
	file, err := read(src)
	defer file.Close()
	if err != nil {
		return err
	}

	// 获取文件字节数
	rows, err := ReadLines(src)

	if err != nil {
		return err
	}

	// 重写文件
	err = Clear(src)
	if err != nil {
		return err
	}

	for line, str := range rows {
		if line == row {
			err := Append(src, content)
			if err != nil {
				return err
			}

			if replace {
				continue
			}
		}
		err := Append(src, str)
		if err != nil {
			return err
		}

	}

	return nil
}
