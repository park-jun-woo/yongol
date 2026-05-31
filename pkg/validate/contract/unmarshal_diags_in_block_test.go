//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestUnmarshalDiagsInBlock — 단일 블록 내 Unmarshal 에러 미처리 진단 생성 검증
package contract

import (
	"go/ast"
	"os"
	"path/filepath"
	"testing"
)

func TestUnmarshalDiagsInBlock(t *testing.T) {
	dir := t.TempDir()

	t.Run("ignored unmarshal flagged", func(t *testing.T) {
		p := filepath.Join(dir, "u.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(b []byte) {\n  var v T\n  json.Unmarshal(b, &v)\n  _ = v\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, err := parseGoFile(p)
		if err != nil {
			t.Fatal(err)
		}
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := unmarshalDiagsInBlock(fset, file, p, fn.Body); len(d) != 1 {
			t.Fatalf("expected 1 diag, got %+v", d)
		}
	})

	t.Run("explicit discard → safe", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		if err := os.WriteFile(p, []byte("package p\nfunc F(b []byte) {\n  var v T\n  _ = json.Unmarshal(b, &v)\n  _ = v\n}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, _ := parseGoFile(p)
		fn := file.Decls[0].(*ast.FuncDecl)
		if d := unmarshalDiagsInBlock(fset, file, p, fn.Body); len(d) != 0 {
			t.Errorf("explicit discard should be safe, got %+v", d)
		}
	})
}
