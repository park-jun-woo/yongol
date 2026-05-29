//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestPanicDiagsInBody — 함수 body 내 panic() 호출마다 PRV-10 Diagnostic 생성 검증

package contract

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"
)

func TestPanicDiagsInBody(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "p.go")
	if err := os.WriteFile(p, []byte("package service\nfunc F() { panic(\"x\"); panic(\"y\") }\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	fset, file, err := parseGoFile(p)
	if err != nil {
		t.Fatal(err)
	}
	fn := file.Decls[0].(*ast.FuncDecl)
	diags := panicDiagsInBody(fset, file, p, fn.Body)
	if len(diags) != 2 {
		t.Fatalf("expected 2 panic diags, got %d (%+v)", len(diags), diags)
	}

	t.Run("nolint suppresses", func(t *testing.T) {
		q := filepath.Join(dir, "q.go")
		if err := os.WriteFile(q, []byte("package service\nfunc G() {\n  panic(\"x\") // nolint:panic\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(q)
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := panicDiagsInBody(fset, file, q, fn.Body); len(d) != 0 {
			t.Errorf("nolint should suppress, got %+v", d)
		}
	})
}
