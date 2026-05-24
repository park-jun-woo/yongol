//ff:func feature=validate type=test control=sequence topic=rego-structural
//ff:what Run — Rego 전체 검증 (빈 fs) 검증

package rego

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestRun_Rego(t *testing.T) {
	t.Run("empty fullstack with empty policy dir returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		policyDir := filepath.Join(tmp, "policy")
		os.MkdirAll(policyDir, 0o755)
		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := Run(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})
}
