//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what runXoh02StatusDeclared — TestXoh02StatusDeclared table-driven 개별 케이스 검증

package hurl_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runXoh02StatusDeclared(t *testing.T, c TestXoh02StatusDeclaredCase) {
	t.Helper()
	diags := xoh02StatusDeclared(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[XOH-02]") {
			t.Errorf("message should contain [XOH-02], got %q", d.Message)
		}
	}
}
