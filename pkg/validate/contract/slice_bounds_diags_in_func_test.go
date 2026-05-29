//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestSliceBoundsDiagsInFunc — 함수 body 내 가드 없는 x[0] 접근 진단 생성 검증

package contract

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"
)

func TestSliceBoundsDiagsInFunc(t *testing.T) {
	dir := t.TempDir()

	t.Run("unguarded x[0]", func(t *testing.T) {
		p := filepath.Join(dir, "s.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(xs []int) int { return xs[0] }\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, err := parseGoFile(p)
		if err != nil {
			t.Fatal(err)
		}
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := sliceBoundsDiagsInFunc(fset, file, p, fn.Body); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("len-guarded", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(xs []int) int {\n  if len(xs) == 0 { return 0 }\n  return xs[0]\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(p)
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := sliceBoundsDiagsInFunc(fset, file, p, fn.Body); len(d) != 0 {
			t.Errorf("guarded access should be safe, got %+v", d)
		}
	})
}
