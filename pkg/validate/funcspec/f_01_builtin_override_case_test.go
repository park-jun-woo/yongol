//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=funcspec-structural
//ff:what runF01BuiltinOverride — TestF01BuiltinOverride table-driven 개별 케이스 검증

package funcspec

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func runF01BuiltinOverride(t *testing.T, c TestF01BuiltinOverrideCase) {
	t.Helper()
	fs := &yongol.Fullstack{
		YongolPkgSpecs:   c.builtin,
		ProjectFuncSpecs: c.project,
	}
	diags := f01BuiltinOverride(fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if d.Phase != diagnostic.PhaseValidate {
			t.Errorf("expected PhaseValidate, got %q", d.Phase)
		}
	}
}
