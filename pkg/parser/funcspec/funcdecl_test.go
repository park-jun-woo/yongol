//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what extractReturnTypes / isStubBody / processFuncDecl / findFuncDeclLine / extractGoParseErrorLine

package funcspec

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseDeclT(t *testing.T, src string) (*token.FileSet, *ast.File) {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return fset, f
}

func firstFunc(f *ast.File) *ast.FuncDecl {
	for _, d := range f.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok {
			return fn
		}
	}
	return nil
}

func TestExtractReturnTypes(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want []string
	}{
		{"no results", "package p\nfunc f() {}", nil},
		{"single", "package p\nfunc f() error { return nil }", []string{"error"}},
		{"two anon", "package p\nfunc f() (Resp, error) { return Resp{}, nil }", []string{"Resp", "error"}},
		{"named group expands", "package p\nfunc f() (a, b int) { return }", []string{"int", "int"}},
		{"pointer", "package p\nfunc f() *T { return nil }", []string{"*T"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fset, f := parseDeclT(t, c.src)
			got := extractReturnTypes(fset, firstFunc(f))
			if len(got) != len(c.want) {
				t.Fatalf("got %v, want %v", got, c.want)
			}
			for i := range got {
				if got[i] != c.want[i] {
					t.Errorf("got[%d] = %q, want %q", i, got[i], c.want[i])
				}
			}
		})
	}
}

func TestProcessFuncDecl(t *testing.T) {
	t.Run("matching name with body", func(t *testing.T) {
		fset, f := parseDeclT(t, `package p
func HashPassword() (Resp, error) { return Resp{Status: "x"}, nil }`)
		spec := &FuncSpec{Name: "hash_password"}
		processFuncDecl(firstFunc(f), fset, spec)
		if !spec.HasBody {
			t.Errorf("expected HasBody true")
		}
		if len(spec.ReturnTypes) != 2 {
			t.Errorf("ReturnTypes = %v", spec.ReturnTypes)
		}
	})
	t.Run("name mismatch is no-op", func(t *testing.T) {
		fset, f := parseDeclT(t, "package p\nfunc Other() {}")
		spec := &FuncSpec{Name: "hash_password"}
		processFuncDecl(firstFunc(f), fset, spec)
		if spec.HasBody || spec.ReturnTypes != nil {
			t.Errorf("expected no-op, got %+v", spec)
		}
	})
}

func TestFindFuncDeclLine(t *testing.T) {
	fset, f := parseDeclT(t, "package p\n\nfunc Target() {}\n")
	if line := findFuncDeclLine(f, fset, "Target"); line != 3 {
		t.Errorf("line = %d, want 3", line)
	}
	if line := findFuncDeclLine(f, fset, "Missing"); line != 0 {
		t.Errorf("missing func line = %d, want 0", line)
	}
}
