//ff:func feature=chain type=test control=sequence
//ff:what resolveFuncSpecPath 가 snake_case 직접경로 / glob 매칭 / fallback 을 올바르게 선택하는지 검증
package chain

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

func TestResolveFuncSpecPath(t *testing.T) {
	specsDir := t.TempDir()

	// Direct snake_case path exists.
	pkgDir := filepath.Join(specsDir, "func", "auth")
	if err := os.MkdirAll(pkgDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(pkgDir, "hash_password.go"), []byte("// @func\nfunc hashPassword() {}\n"), 0o644); err != nil {
		t.Fatalf("write direct: %v", err)
	}
	spec := funcspec.FuncSpec{Package: "auth", Name: "hashPassword"}
	if got := resolveFuncSpecPath(spec, "hashPassword", specsDir); got != "func/auth/hash_password.go" {
		t.Errorf("direct path: got %q", got)
	}

	// Glob fallback: file named differently but contains @func + funcName.
	mailDir := filepath.Join(specsDir, "func", "mail")
	if err := os.MkdirAll(mailDir, 0o755); err != nil {
		t.Fatalf("mkdir mail: %v", err)
	}
	if err := os.WriteFile(filepath.Join(mailDir, "bundle.go"), []byte("// @func\nfunc SendWelcome() {}\n"), 0o644); err != nil {
		t.Fatalf("write glob: %v", err)
	}
	specGlob := funcspec.FuncSpec{Package: "mail", Name: "SendWelcome"}
	if got := resolveFuncSpecPath(specGlob, "SendWelcome", specsDir); got != "func/mail/bundle.go" {
		t.Errorf("glob fallback: got %q, want func/mail/bundle.go", got)
	}

	// No file at all → returns computed snake_case relPath.
	specNone := funcspec.FuncSpec{Package: "queue", Name: "PushJob"}
	if got := resolveFuncSpecPath(specNone, "PushJob", specsDir); got != "func/queue/push_job.go" {
		t.Errorf("no file fallback: got %q", got)
	}
}
