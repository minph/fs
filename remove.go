package fs

import (
	"os"
	"path"
	"regexp"
	"strings"
)

// Remove 同os.Remove
//
// 只用于删除单个文件和空目录
func Remove(src string) error {

	return os.Remove(src)
}

// RemoveAll 同os.RemoveAll
func RemoveAll(src string) error {

	return os.RemoveAll(src)
}

//RemoveDir  删除目录全部内容，包含目录自身
func RemoveDir(folder string) error {
	content, err := DirSub(folder)

	if err != nil {
		return err
	}

	for _, item := range content {
		err := RemoveAll(item)
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveExt 删除目录下指定后缀名的文件
func RemoveExt(folder, ext string) error {
	files, err := DirAll(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		// 判断文件名是否相匹配后缀名
		if path.Ext(file) == ext || path.Ext(file) == "."+ext {
			err := Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveNamesRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件
func RemoveNamesRegexp(folder, pattern string) error {
	files, err := DirAll(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		// 判断文件名是否匹配正则表达式
		filename := path.Base(file)
		matched, err := regexp.MatchString(filename, pattern)
		if err != nil {
			return err
		}
		if matched {
			err := Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveNamesContain 目录下所有文件中，删除文件名包含给定字符的文件
func RemoveNamesContain(folder, containStr string) error {
	files, err := DirAll(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		// 判断文件名是否包含给定字符
		filename := path.Base(file)
		if strings.Index(filename, containStr) != -1 {
			err := Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
