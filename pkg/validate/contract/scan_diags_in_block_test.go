//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestScanDiagsInBlock — 단일 블록 내 Scan 호출 에러 누락 진단 생성 검증
package contract

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"
)

func TestScanDiagsInBlock(t *testing.T) {
	dir := t.TempDir()

	t.Run("ignored scan flagged", func(t *testing.T) {
		p := filepath.Join(dir, "s.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(row R) {\n  var x int\n  row.Scan(&x)\n  _ = x\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, err := parseGoFile(p)
		if err != nil {
			t.Fatal(err)
		}
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := scanDiagsInBlock(fset, file, p, fn.Body); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("later err check → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(row R) error {\n  var x int\n  err := row.Scan(&x)\n  if err != nil { return err }\n  _ = x\n  return nil\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(p)
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := scanDiagsInBlock(fset, file, p, fn.Body); len(d) != 0 {
			t.Errorf("later err check should be safe, got %+v", d)
		}
	})
}
