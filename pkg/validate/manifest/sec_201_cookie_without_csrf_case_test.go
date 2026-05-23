//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-auth
//ff:what runSec201CookieWithoutCsrf — TestSec201CookieWithoutCsrf table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runSec201CookieWithoutCsrf(t *testing.T, c TestSec201CookieWithoutCsrfCase) {
	t.Helper()
	diags := sec201CookieWithoutCsrf(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[SEC-201]") {
			t.Errorf("expected [SEC-201], got %q", d.Message)
		}
	}
}
