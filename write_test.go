package fs

import "testing"

func TestWriteStringAtLine(t *testing.T) {
	src := "examples/a/a2.txt"
	_ = WriteStringAtLine(src, 5, "Add at line 5\n", true)
}
