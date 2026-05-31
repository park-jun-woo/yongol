//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"testing"
)

func TestImportName(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want string
	}{
		{"plain", "package p\nimport \"fmt\"", "fmt"},
		{"path base", "package p\nimport \"net/http\"", "http"},
		{"alias", "package p\nimport x \"net/http\"", "x"},
		{"version suffix", "package p\nimport \"github.com/foo/bar/v2\"", "bar"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := importName(importSpecOf(t, c.src)); got != c.want {
				t.Errorf("importName = %q, want %q", got, c.want)
			}
		})
	}
}
