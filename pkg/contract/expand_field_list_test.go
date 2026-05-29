//ff:func feature=contract type=test control=iteration dimension=1
//ff:what test: TestExpandFieldList — nil 리스트, 다중 이름 그룹, 익명 반환 타입 전개 분기 검증

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// fieldListFromFunc parses a single func declaration and returns its
// parameter and result field lists for exercising expandFieldList.
func fieldListFromFunc(t *testing.T, src string) (*token.FileSet, *ast.FuncDecl) {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", "package p\n"+src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse error: %v", err)
	}
	for _, decl := range f.Decls {
		if fd, ok := decl.(*ast.FuncDecl); ok {
			return fset, fd
		}
	}
	t.Fatalf("no func decl in %q", src)
	return nil, nil
}

func TestExpandFieldListNil(t *testing.T) {
	fset := token.NewFileSet()
	if got := expandFieldList(fset, nil, false); got != nil {
		t.Errorf("nil list: got %v want nil", got)
	}
}

func TestExpandFieldListParams(t *testing.T) {
	fset, fd := fieldListFromFunc(t, "func F(a, b int, name string) {}")
	got := expandFieldList(fset, fd.Type.Params, true)
	want := []FuncParam{
		{Name: "a", Type: "int"},
		{Name: "b", Type: "int"},
		{Name: "name", Type: "string"},
	}
	if len(got) != len(want) {
		t.Fatalf("got %d params (%v) want %d", len(got), got, len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("param[%d]: got %+v want %+v", i, got[i], want[i])
		}
	}
}

func TestExpandFieldListUnnamedResults(t *testing.T) {
	fset, fd := fieldListFromFunc(t, "func F() (string, error) {}")
	got := expandFieldList(fset, fd.Type.Results, false)
	if len(got) != 2 {
		t.Fatalf("got %d results (%v) want 2", len(got), got)
	}
	if got[0].Type != "string" || got[0].Name != "" {
		t.Errorf("result[0]: got %+v want unnamed string", got[0])
	}
	if got[1].Type != "error" || got[1].Name != "" {
		t.Errorf("result[1]: got %+v want unnamed error", got[1])
	}
}
