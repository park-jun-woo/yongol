//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-manifest
//ff:what runXoh07CSRFOnMutation — TestXoh07CSRFOnMutation table-driven 개별 케이스 검증

package hurl_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runXoh07CSRFOnMutation(t *testing.T, c TestXoh07CSRFOnMutationCase) {
	t.Helper()
	diags := xoh07CSRFOnMutation(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[XOH-07]") {
			t.Errorf("message should contain [XOH-07], got %q", d.Message)
		}
	}
}
