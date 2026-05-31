//ff:func feature=gen-splitter type=test control=sequence
//ff:what zz_zerocov_test — splitter 패키지의 0% 커버리지 함수(cleanOriginal/preserveComments/isPreservedFile/SplitDirectory) 단위 테스트
package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestPreserveComments_ZeroCov(t *testing.T) {
	src := `package p

// doc comment
//go:generate something
func Foo() {}

func Bar() {}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	cmap := ast.NewCommentMap(fset, file, file.Comments)

	fooDecl := findFuncDecl(file, "Foo")
	barDecl := findFuncDecl(file, "Bar")
	if fooDecl == nil || barDecl == nil {
		t.Fatal("decls not found")
	}

	groups := preserveComments(cmap, fooDecl)
	if len(groups) == 0 {
		t.Fatal("expected comment groups for Foo")
	}
	assertNoNilGroups(t, groups)

	// Bar has no associated comments → empty slice.
	if got := preserveComments(cmap, barDecl); len(got) != 0 {
		t.Errorf("expected no comments for Bar, got %d", len(got))
	}
}
