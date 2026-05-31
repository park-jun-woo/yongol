//ff:func feature=gen-splitter type=test control=sequence
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"testing"
)

func TestSuffixFor(t *testing.T) {
	typeDecl := firstDecl(t, "package p\ntype Row struct{}")
	if got := suffixFor(ToolOAPICodegen, false, typeDecl); got != ".gen.go" {
		t.Errorf("oapi = %q, want .gen.go", got)
	}
	if got := suffixFor(ToolSQLC, true, typeDecl); got != ".model.go" {
		t.Errorf("sqlc models = %q, want .model.go", got)
	}
	if got := suffixFor(Tool("unknown"), false, typeDecl); got != ".go" {
		t.Errorf("default = %q, want .go", got)
	}
}
