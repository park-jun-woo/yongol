//ff:func feature=chain type=test control=sequence
//ff:what traceFuncSpecs 가 @call sequence 를 func spec 과 매칭하고 @call 없을 때 nil 을 반환하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTraceFuncSpecs(t *testing.T) {
	specsDir := t.TempDir()
	pkgDir := filepath.Join(specsDir, "func", "auth")
	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "verify_password.go"), []byte("// @func\nfunc VerifyPassword() {}\n"), 0o644); err != nil {
		t.Fatalf("write: %v", err)
	}

	specs := []funcspec.FuncSpec{{Package: "auth", Name: "VerifyPassword"}}

	sf := &ssac.ServiceFunc{
		Name: "Login",
		Sequences: []ssac.Sequence{
			{Type: "call", Model: "auth.VerifyPassword"},
			{Type: "get", Model: "User.FindByEmail"}, // not a call
			{Type: "call", Model: "NoDot"},           // no "." → skipped
		},
	}

	links := traceFuncSpecs(sf, specs, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 funcspec link, got %d: %+v", len(links), links)
	}
	if links[0].Kind != "FuncSpec" || links[0].File != "func/auth/verify_password.go" {
		t.Errorf("link fields: %+v", links[0])
	}
	if links[0].Summary != "@func auth.VerifyPassword" {
		t.Errorf("summary: %q", links[0].Summary)
	}

	// No @call sequences → nil.
	sfNone := &ssac.ServiceFunc{Name: "X", Sequences: []ssac.Sequence{{Type: "get", Model: "User.X"}}}
	if traceFuncSpecs(sfNone, specs, specsDir) != nil {
		t.Error("expected nil when no @call sequences")
	}
}
