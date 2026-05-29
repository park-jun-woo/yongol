//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestParseGoFile — preserved 파일 AST 파싱 (fset 동반) 검증

package contract

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseGoFile(t *testing.T) {
	dir := t.TempDir()

	t.Run("valid file", func(t *testing.T) {
		p := filepath.Join(dir, "ok.go")
		if err := os.WriteFile(p, []byte("package p\n\nfunc F() {}\n"), 0o644); err != nil {
			t.Fatal(err)
		}
		fset, file, err := parseGoFile(p)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if fset == nil || file == nil {
			t.Fatal("expected non-nil fset and file")
		}
		if file.Name.Name != "p" {
			t.Fatalf("package name = %q, want p", file.Name.Name)
		}
	})

	t.Run("missing file", func(t *testing.T) {
		if _, _, err := parseGoFile(filepath.Join(dir, "missing.go")); err == nil {
			t.Fatal("expected error for missing file")
		}
	})

	t.Run("syntax error", func(t *testing.T) {
		p := filepath.Join(dir, "bad.go")
		if err := os.WriteFile(p, []byte("package p\nfunc {"), 0o644); err != nil {
			t.Fatal(err)
		}
		if _, _, err := parseGoFile(p); err == nil {
			t.Fatal("expected parse error for malformed source")
		}
	})
}
