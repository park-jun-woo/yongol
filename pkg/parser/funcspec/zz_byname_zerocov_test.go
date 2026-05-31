//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용

package funcspec

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func bnFile(t *testing.T, src string) (*token.FileSet, *ast.File) {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return fset, f
}

const bnStructSrc = `package p
type FooRequest struct {
	Name string ` + "`json:\"name\"`" + `
	Age  int
}
type FooResponse struct {
	ID int
}
func g() {}
`

func TestBuildFields_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	gd := f.Decls[0].(*ast.GenDecl)
	st := gd.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	fields := buildFields(st.Fields.List[0])
	if len(fields) != 1 || fields[0].Name != "Name" || fields[0].JSONName != "name" {
		t.Errorf("buildFields wrong: %#v", fields)
	}
}

func TestExtractFields_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	st := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	if got := extractFields(st); len(got) != 2 {
		t.Errorf("expected 2 fields, got %d", len(got))
	}
}

func TestCollectStructsFromGenDecl_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	result := map[string][]Field{}
	collectStructsFromGenDecl(f.Decls[0].(*ast.GenDecl), result)
	if _, ok := result["FooRequest"]; !ok {
		t.Errorf("FooRequest not collected: %v", result)
	}
}

func TestCollectStructsFromFile_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	result := map[string][]Field{}
	collectStructsFromFile(f, result)
	if len(result) != 2 {
		t.Errorf("expected 2 structs, got %d", len(result))
	}
}

func TestFillSpecFromTypeMap_ZeroCov(t *testing.T) {
	spec := &FuncSpec{Name: "foo"}
	tm := map[string][]Field{
		"FooRequest":  {{Name: "Name"}},
		"FooResponse": {{Name: "ID"}},
	}
	fillSpecFromTypeMap(spec, tm)
	if len(spec.RequestFields) != 1 || len(spec.ResponseFields) != 1 {
		t.Errorf("spec not filled: %#v", spec)
	}
}

func TestProcessTypeSpecs_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	spec := &FuncSpec{Name: "foo"}
	for _, d := range f.Decls {
		if gd, ok := d.(*ast.GenDecl); ok {
			processTypeSpecs(gd, spec, "FooRequest", "FooResponse")
		}
	}
	if len(spec.RequestFields) != 2 || len(spec.ResponseFields) != 1 {
		t.Errorf("type specs not processed: %#v", spec)
	}
}

func TestProcessDecl_ZeroCov(t *testing.T) {
	fset, f := bnFile(t, bnStructSrc)
	spec := &FuncSpec{Name: "foo"}
	for _, d := range f.Decls {
		processDecl(d, fset, spec, "FooRequest", "FooResponse")
	}
	if len(spec.RequestFields) != 2 {
		t.Errorf("processDecl did not handle type decl: %#v", spec)
	}
}

func TestLoadTypeMapForDir_ZeroCov(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "m.go"), []byte(bnStructSrc), 0644); err != nil {
		t.Fatal(err)
	}
	cache := map[string]map[string][]Field{}
	seen := map[string]struct{}{}
	tm, diags := loadTypeMapForDir(dir, cache, seen, nil)
	if tm == nil {
		t.Errorf("expected type map")
	}
	_ = diags
	// second call hits cache.
	tm2, _ := loadTypeMapForDir(dir, cache, seen, []diagnostic.Diagnostic{})
	if tm2 == nil {
		t.Errorf("expected cached type map")
	}
}
