package dir

import (
	"fmt"
	"github.com/minph/fs/utils"
	"io/ioutil"
	"os"
	"path"
)

// Is 判断是不是文件夹
func Is(path string) bool {
	stat, err := os.Stat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

// Info 原生方法读取目录信息
func Info(folder string) ([]os.FileInfo, error) {

	// 读取目录
	info, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}

	return info, nil
}

// Content 读取目录下所有文件和目录信息
func Content(folder string) ([]string, error) {

	var result []string

	// 读取目录
	info, err := Info(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {
		result = append(result, utils.Format(path.Join(folder, item.Name())))
	}

	return result, nil
}

// Folders 读取目录下的子文件夹
// 即不包含子文件
func Folders(folder string) ([]string, error) {

	var result []string

	// 读取目录
	info, err := Info(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {
		if item.IsDir() {
			result = append(result, utils.Format(path.Join(folder, item.Name())))
		}
	}

	return result, nil
}

// Files 读取目录下的所有子文件
// 即不包含子目录
func Files(folder string) ([]string, error) {

	var result []string

	// 读取目录
	info, err := Info(folder)
	if err != nil {
		return nil, err
	}

	for _, item := range info {
		if !item.IsDir() {
			result = append(result, utils.Format(path.Join(folder, item.Name())))
		}
	}

	return result, nil
}

// All 读取目录和子目录下的所有文件名信息
//
// 全部为文件名，不再包含目录名
func All(folder string) ([]string, error) {
	var result []string

	// 读取目录
	info, err := Info(folder)
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
			subInfo, err := All(fullName)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, subItem := range subInfo {
				result = append(result, utils.Format(subItem))
			}
		}
	}

	return result, nil
}
