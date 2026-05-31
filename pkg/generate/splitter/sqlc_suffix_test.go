//ff:func feature=gen-splitter type=test control=sequence
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"testing"
)

func TestSqlcSuffix(t *testing.T) {
	typeDecl := firstDecl(t, "package p\ntype Row struct{}")
	funcDecl := firstDecl(t, "package p\nfunc f() {}")
	if got := sqlcSuffix(false, typeDecl); got != ".sql.go" {
		t.Errorf("non-models type = %q, want .sql.go", got)
	}
	if got := sqlcSuffix(true, typeDecl); got != ".model.go" {
		t.Errorf("models type = %q, want .model.go", got)
	}
	if got := sqlcSuffix(true, funcDecl); got != ".sql.go" {
		t.Errorf("models func = %q, want .sql.go", got)
	}
}
