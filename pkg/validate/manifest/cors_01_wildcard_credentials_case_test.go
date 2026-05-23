//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-cors
//ff:what runCors01WildcardCredentials — TestCors01WildcardCredentials table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runCors01WildcardCredentials(t *testing.T, c TestCors01WildcardCredentialsCase) {
	t.Helper()
	diags := cors01WildcardCredentials(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[CORS-01]") {
			t.Errorf("expected [CORS-01], got %q", d.Message)
		}
	}
}
