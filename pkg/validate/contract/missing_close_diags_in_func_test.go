//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestMissingCloseDiagsInFunc — 함수 body 획득-then-close 매칭, 누락 시 PRV-17 검증

package contract

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"
)

func TestMissingCloseDiagsInFunc(t *testing.T) {
	dir := t.TempDir()

	t.Run("missing close", func(t *testing.T) {
		p := filepath.Join(dir, "m.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F() {\n  f, _ := os.Open(\"x\")\n  _ = f\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, err := parseGoFile(p)
		if err != nil {
			t.Fatal(err)
		}
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := missingCloseDiagsInFunc(fset, file, p, fn.Body); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("with defer close", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F() {\n  f, _ := os.Open(\"x\")\n  defer f.Close()\n  _ = f\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(p)
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := missingCloseDiagsInFunc(fset, file, p, fn.Body); len(d) != 0 {
			t.Errorf("defer close should be safe, got %+v", d)
		}
	})
}
