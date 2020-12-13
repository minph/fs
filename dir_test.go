package fs

import (
	"fmt"
	"testing"
)

func TestDirAll(t *testing.T) {
	dir, _ := DirAll(`examples\\next`)
	for _, path := range dir {
		fmt.Println(path)
	}
}
