//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-security-headers
//ff:what runSec301CspPermissive — TestSec301CspPermissive table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runSec301CspPermissive(t *testing.T, c TestSec301CspPermissiveCase) {
	t.Helper()
	diags := sec301CspPermissive(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[SEC-301]") {
			t.Errorf("expected [SEC-301], got %q", d.Message)
		}
	}
}
