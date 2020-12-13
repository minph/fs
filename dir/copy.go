package dir

import (
	"github.com/minph/fs"
	"github.com/minph/fs/utils"
	"os"
	"path"
	"path/filepath"
)

// Create 跨目录创建文件夹
func Create(folder string) error {
	if !utils.Exist(folder) {
		err := os.MkdirAll(folder, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

// Copy 复制文件夹
func Copy(folder, target string) error {

	// 获取文件夹全部文件信息
	info, err := All(folder)
	if err != nil {
		return err
	}

	for _, file := range info {

		// 获取相对路径
		rel, err := filepath.Rel(folder, file)

		if err != nil {
			return err
		}

		// 逐一复制文件
		// 使用 ToSlash 防止反斜杠
		copyFile := path.Join(target, filepath.ToSlash(rel))
		err = fs.CopySafe(file, copyFile)
		if err != nil {
			return err
		}
	}
	return nil
}
