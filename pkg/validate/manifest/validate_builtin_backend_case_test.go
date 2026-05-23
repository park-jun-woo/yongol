//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-infra
//ff:what runValidateBuiltinBackend — TestValidateBuiltinBackend table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runValidateBuiltinBackend(t *testing.T, c TestValidateBuiltinBackendCase) {
	t.Helper()
	diags := validateBuiltinBackend(c.fs, c.spec)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[XNC-90]") {
			t.Errorf("expected [XNC-90], got %q", d.Message)
		}
	}
}
