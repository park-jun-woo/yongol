//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runC11IpKeyRequiresProxy — TestC11IpKeyRequiresProxy table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runC11IpKeyRequiresProxy(t *testing.T, c TestC11IpKeyRequiresProxyCase) {
	t.Helper()
	diags := c11IpKeyRequiresProxy(c.fs)
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
		if !strings.Contains(d.Message, "[C-11]") {
			t.Errorf("expected [C-11], got %q", d.Message)
		}
		if !strings.Contains(d.Advice, "trusted_proxies") {
			t.Errorf("expected advice to mention trusted_proxies, got %q", d.Advice)
		}
	}
}
