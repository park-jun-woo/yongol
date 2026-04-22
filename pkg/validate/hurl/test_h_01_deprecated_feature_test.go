//ff:func feature=validate type=test control=sequence topic=hurl-structural
//ff:what H-1 테스트 — .feature 존재 시 에러 생성 확인

package hurl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestH01DeprecatedFeature(t *testing.T) {
	dir := t.TempDir()
	testsDir := filepath.Join(dir, "tests")
	if err := os.MkdirAll(testsDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(testsDir, "sample.feature"), []byte(""), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := h01DeprecatedFeature(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
}
