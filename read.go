package fs

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func readFile(src string) (*os.File, error) {
	file, err := os.OpenFile(src, os.O_CREATE|os.O_RDONLY, os.ModePerm)

	if err != nil {
		return nil, err
	}

	return file, nil
}

// ReadByte 读取文件全部内容
func ReadByte(src string) ([]byte, error) {

	// 打开文件
	file, err := readFile(src)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(file)
}

// ReadString 读取文件全部内容并返回字符串
func ReadString(src string) (string, error) {

	// 打开文件
	file, err := readFile(src)
	defer file.Close()
	if err != nil {
		return "", err
	}

	// 读取结果并转为字符串类型
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	return string(content), err
}

// ReadByRow 按行读取文件内容
func ReadByRow(src string) ([]string, error) {

	// 打开文件
	file, err := readFile(src)
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
			break
		}

		// 存入每行数据
		content = append(content, line)
	}

	return content, nil
}

// ReadByBytes 按块方式读取文件内容
func ReadByBytes(src string, bufferSize int64) ([][]byte, error) {

	// 打开文件
	file, err := readFile(src)
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

// ReadAt 读取指定位置之后内容
//
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func ReadAt(src string, position int64) ([]byte, error) {
	// 打开文件
	file, err := readFile(src)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	// 获取文件字节数

	bytes, err := ReadByte(src)
	if err != nil {
		return nil, err
	}

	// 获取定位之后的内容
	var length int64
	if position >= 0 {
		file.Seek(position, 0)
		length = int64(len(bytes)) - position
	} else {
		file.Seek(int64(len(bytes))+position, 0)
		length = -1 * position
	}

	buffer := make([]byte, length)

	// 读取内容
	_, err = file.Read(buffer)
	if err != nil {
		return nil, err
	}

	return buffer, nil
}

// ReadStringAt 读取指定位置之后内容
// 返回字符串
// 正数则从开始到最后定位
// 负数则从最后到开始定位
func ReadStringAt(src string, position int64) (string, error) {
	content, err := ReadAt(src, position)
	if err != nil {
		return "", err
	}

	return string(content), nil
}
