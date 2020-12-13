package fs

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func read(src string) (*os.File, error) {
	file, err := os.OpenFile(src, os.O_RDONLY, os.ModePerm)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// readAll 读取文件全部内容
func readAll(src string) ([]byte, error) {

	// 打开文件
	file, err := read(src)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

// ReadAll 读取文件全部内容并返回字符串
func ReadAll(src string) (string, error) {
	// 读取结果并转为字符串类型
	content, err := readAll(src)
	if err != nil {
		return "", err
	}
	return string(content), err
}

// ReadLines 按行读取文件内容
func ReadLines(src string) ([]string, error) {

	// 打开文件
	file, err := read(src)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// 创建行阅读器
	var content []string
	reader := bufio.NewReader(file)

	// 进行循环
	for {
		//以每行为结束符读入一行
		line, err := reader.ReadString('\n')

		if err != nil || io.EOF == err {
			content = append(content, line)
			break
		}

		// 存入每行数据
		content = append(content, line)
	}

	return content, nil
}

// ReadBytes 按 byte 块方式读取文件内容
func ReadBytes(src string, bufferSize int64) ([][]byte, error) {

	// 打开文件
	file, err := read(src)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// 创建文件块
	var content [][]byte
	buffer := make([]byte, bufferSize)

	// 进行循环
	for {
		chunk, err := file.Read(buffer)
		if err != nil || io.EOF == err {
			break
		}
		content = append(content, buffer[:chunk])
	}

	return content, nil
}

// ReadLineAt 读取指定行的内容
func ReadLineAt(src string, row int) (string, error) {
	rows, err := ReadLines(src)
	if err != nil {
		return "", err
	}

	for i := range rows {
		if i == row {
			return rows[i], nil
		}
	}

	return "", nil
}
