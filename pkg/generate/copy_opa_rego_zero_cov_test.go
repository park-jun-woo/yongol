//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCopyOPARego_ZeroCov(t *testing.T) {
	// missing policy dir → skip
	fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
	if err := copyOPARego(fs, t.TempDir()); err != nil {
		t.Fatalf("missing policy err: %v", err)
	}
	// real policy dir with .rego + a non-rego file (skipped) + subdir (skipped)
	specs := t.TempDir()
	pol := filepath.Join(specs, "policy")
	if err := os.MkdirAll(filepath.Join(pol, "sub"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pol, "auth.rego"), []byte("package auth"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(pol, "readme.md"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	arts := t.TempDir()
	if err := copyOPARego(&yongol.Fullstack{SpecsDir: specs}, arts); err != nil {
		t.Fatalf("real copy err: %v", err)
	}
	if _, err := os.Stat(filepath.Join(arts, "backend", "policy", "auth.rego")); err != nil {
		t.Errorf("expected auth.rego copied: %v", err)
	}
}
