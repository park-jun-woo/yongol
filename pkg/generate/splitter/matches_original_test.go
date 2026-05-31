//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"testing"
)

func TestMatchesOriginal(t *testing.T) {
	cases := []struct {
		name string
		tool Tool
		want bool
	}{
		{"api.gen.go", ToolOAPICodegen, true},
		{"api.go", ToolOAPICodegen, false},
		{"models.go", ToolSQLC, true},
		{"users.sql.go", ToolSQLC, true},
		{"users.model.go", ToolSQLC, true},
		{"querier.go", ToolSQLC, false},
		{"x.go", Tool("other"), false},
	}
	for _, c := range cases {
		if got := matchesOriginal(c.name, c.tool); got != c.want {
			t.Errorf("matchesOriginal(%q,%v) = %v, want %v", c.name, c.tool, got, c.want)
		}
	}
}
