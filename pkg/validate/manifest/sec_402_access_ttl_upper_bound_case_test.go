//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-auth
//ff:what runSec402AccessTTLUpperBound — TestSec402AccessTTLUpperBound table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runSec402AccessTTLUpperBound(t *testing.T, c TestSec402AccessTTLUpperBoundCase) {
	t.Helper()
	diags := sec402AccessTTLUpperBound(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[SEC-402]") {
			t.Errorf("expected [SEC-402], got %q", d.Message)
		}
	}
}
