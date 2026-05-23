//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-structural
//ff:what runH01DeprecatedFeature — TestH01DeprecatedFeature table-driven 개별 케이스 검증

package hurl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runH01DeprecatedFeature(t *testing.T, c TestH01DeprecatedFeatureCase) {
	t.Helper()
	dir := t.TempDir()
	os.MkdirAll(filepath.Join(dir, "tests"), 0o755)
	os.MkdirAll(filepath.Join(dir, "scenario"), 0o755)
	for rel, content := range c.files {
		path := filepath.Join(dir, rel)
		os.MkdirAll(filepath.Dir(path), 0o755)
		os.WriteFile(path, []byte(content), 0o644)
	}
	fs := &yongol.Fullstack{SpecsDir: dir}
	diags := h01DeprecatedFeature(fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[H-1]") {
			t.Errorf("message should contain [H-1], got %q", d.Message)
		}
	}
}
