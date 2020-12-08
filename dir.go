package fs

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

// IsDir 判断是不是文件夹
func IsDir(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// DirInfo 原生方法读取目录信息
func DirInfo(folder string) ([]os.FileInfo, error) {

	// 读取目录
	info, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// DirSub 读取目录下所有文件和目录信息
func DirSub(folder string) ([]string, error) {

	var result []string

	// 读取目录
	info, err := DirInfo(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {
		result = append(result, FormatPath(path.Join(folder, item.Name())))
	}

	return result, nil
}

// DirSubfolder 读取目录下的子文件夹
// 即不包含子文件
func DirSubfolder(folder string) ([]string, error) {

	var result []string

	// 读取目录
	info, err := DirInfo(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {
		if item.IsDir() {
			result = append(result, FormatPath(path.Join(folder, item.Name())))
		}
	}

	return result, nil
}

// DirSubfile 读取目录下的所有子文件
// 即不包含子目录
func DirSubfile(folder string) ([]string, error) {

	var result []string

	// 读取目录
	info, err := DirInfo(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {
		if !item.IsDir() {
			result = append(result, FormatPath(path.Join(folder, item.Name())))
		}
	}

	return result, nil
}

// DirAll 读取目录和子目录下的所有文件名信息
//
// 全部为文件名，不再包含目录名
func DirAll(folder string) ([]string, error) {
	var result []string

	// 读取目录
	info, err := DirInfo(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {

		// 获取文件绝对路径
		fullName := path.Join(folder, item.Name())
		if !item.IsDir() {
			result = append(result, fullName)
		} else {

			// 递归获取子目录信息
			subInfo, err := DirAll(fullName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, subItem := range subInfo {
				result = append(result, FormatPath(subItem))
			}
		}
	}

	return result, nil
}
