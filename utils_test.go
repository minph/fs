package fs

import "testing"

func TestFormatPath(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test0",
			args: struct{ path string }{path: "../../a\\b\\c"},
			want: "../../a/b/c",
		},
		{
			name: "test1",
			args: struct{ path string }{path: "a//b\\c\\d/e/f"},
			want: "a/b/c/d/e/f",
		},
		{
			name: "test2",
			args: struct{ path string }{path: `a\b//c\d\e/f`},
			want: "a/b/c/d/e/f",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatPath(tt.args.path); got != tt.want {
				t.Errorf("FormatPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
