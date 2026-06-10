//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-auth
//ff:what runSec202RuntimeModeCsrf — TestSec202RuntimeModeCsrf table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runSec202RuntimeModeCsrf(t *testing.T, c TestSec202RuntimeModeCsrfCase) {
	t.Helper()
	diags := sec202RuntimeModeCsrf(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[SEC-202]") {
			t.Errorf("expected [SEC-202], got %q", d.Message)
		}
	}
}
