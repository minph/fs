package dir

import (
	"github.com/minph/fs"
	"path"
	"regexp"
	"strings"
)

// RemoveContent  删除目录全部内容，但不包含目录自身
func RemoveContent(folder string) error {
	content, err := Content(folder)

	if err != nil {
		return err
	}

	for _, item := range content {
		err := fs.RemoveAll(item)
		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveExt 删除目录下指定后缀名的文件
func RemoveExt(folder, ext string) error {
	files, err := Content(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		// 判断文件名是否相匹配后缀名
		if path.Ext(file) == ext || path.Ext(file) == "."+ext {
			err := fs.Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveRegexp 目录下所有文件中，删除文件名匹配正则表达式的文件
func RemoveRegexp(folder, pattern string) error {
	files, err := Content(folder)
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
			err := fs.Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// RemoveContain 目录下所有文件中，删除文件名包含给定字符的文件
func RemoveContain(folder, containStr string) error {
	files, err := Content(folder)
	if err != nil {
		return err
	}
	for _, file := range files {
		// 判断文件名是否包含给定字符
		filename := path.Base(file)
		if strings.Index(filename, containStr) != -1 {
			err := fs.Remove(file)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
