//ff:func feature=validate type=test-helper control=iteration dimension=2 topic=hurl-structural
//ff:what runRun — TestRun table-driven 개별 케이스 검증

package hurl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runRun(t *testing.T, c TestRunCase) {
	t.Helper()
	dir := t.TempDir()
	c.setup(dir)
	fs := &yongol.Fullstack{
		SpecsDir:  dir,
		Presences: c.presences,
	}
	diags := Run(fs)
	for _, code := range c.wantCodes {
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, code) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected diagnostic containing %q not found in %v", code, diags)
		}
	}
	// Also check no extra unexpected diags
	if len(c.wantCodes) == 0 && len(diags) != 0 {
		t.Errorf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
