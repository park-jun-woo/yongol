//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"path/filepath"
	"testing"
)

func TestIsPreservedFile_ZeroCov(t *testing.T) {
	cases := []struct {
		name string
		tool Tool
		want bool
	}{
		{"querier.go", ToolSQLC, true},
		{"db.go", ToolSQLC, true},
		{"models.go", ToolSQLC, false},
		{"foo.sql.go", ToolSQLC, false},
		{"querier.go", ToolOAPICodegen, false},
		{"anything.go", Tool("unknown"), false},
		// basename is taken even with a path prefix.
		{filepath.Join("sub", "querier.go"), ToolSQLC, true},
	}
	for _, c := range cases {
		if got := isPreservedFile(c.name, c.tool); got != c.want {
			t.Errorf("isPreservedFile(%q,%q)=%v want %v", c.name, c.tool, got, c.want)
		}
	}
}
