//ff:func feature=chain type=test control=iteration dimension=2
//ff:what findFuncSpecLink 가 pkg/이름 매칭 시 Link 반환, 불일치 시 false 반환을 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestFindFuncSpecLink(t *testing.T) {
	specsDir := t.TempDir()
	pkgDir := filepath.Join(specsDir, "func", "auth")
	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "hash_password.go"), []byte("// @func\nfunc HashPassword() {}\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	specs := []funcspec.FuncSpec{{Package: "auth", Name: "hashPassword"}}

	// Matching package + case-insensitive name.
	link, ok := findFuncSpecLink("auth.HashPassword", "auth", "HashPassword", specs, specsDir)
	if !ok {
		t.Fatal("expected match")
	}
	if link.Kind != "FuncSpec" || link.File != "func/auth/hash_password.go" {
		t.Errorf("link fields: %+v", link)
	}
	if link.Summary != "@func auth.HashPassword" {
		t.Errorf("summary: %q", link.Summary)
	}
	if link.Line != 2 {
		t.Errorf("line: got %d, want 2", link.Line)
	}

	// Wrong package → no match.
	if _, ok := findFuncSpecLink("mail.HashPassword", "mail", "HashPassword", specs, specsDir); ok {
		t.Error("expected no match for wrong package")
	}
	// Wrong name → no match.
	if _, ok := findFuncSpecLink("auth.Other", "auth", "Other", specs, specsDir); ok {
		t.Error("expected no match for wrong name")
	}
}
