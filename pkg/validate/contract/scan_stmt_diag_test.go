//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanStmtDiag — 단일 stmt 가 PRV-13 위반이면 Diagnostic 생성 검증

package contract

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"
)

func TestScanStmtDiag(t *testing.T) {
	dir := t.TempDir()

	t.Run("bare scan → diag", func(t *testing.T) {
		p := filepath.Join(dir, "s.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(row R) {\n  row.Scan(&x)\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, err := parseGoFile(p)
		if err != nil {
			t.Fatal(err)
		}
		fn := file.Decls[0].(*ast.FuncDecl)
		_, ok := scanStmtDiag(fset, file, p, fn.Body, 0, fn.Body.List[0])
		if !ok {
			t.Fatal("expected diag for bare scan")
		}
	})

	t.Run("if-init scan → no diag", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(row R) {\n  if err := row.Scan(&x); err != nil { return }\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(p)
		fn := file.Decls[0].(*ast.FuncDecl)
		if _, ok := scanStmtDiag(fset, file, p, fn.Body, 0, fn.Body.List[0]); ok {
			t.Error("if-init scan should be discarded (no diag)")
		}
	})

	t.Run("non-scan stmt → no diag", func(t *testing.T) {
		p := filepath.Join(dir, "n.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F() {\n  x := 1\n  _ = x\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(p)
		fn := file.Decls[0].(*ast.FuncDecl)
		if _, ok := scanStmtDiag(fset, file, p, fn.Body, 0, fn.Body.List[0]); ok {
			t.Error("non-scan stmt should produce no diag")
		}
	})
}
