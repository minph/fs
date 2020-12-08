package fs

import (
	"path"
	"strings"
)

// FormatPath 统一路径格式
func FormatPath(src string) string {
	res := path.Clean(src)
	res = strings.ReplaceAll(res, `\\`, "/")
	res = strings.ReplaceAll(res, "\\", "/")
	return res
}
