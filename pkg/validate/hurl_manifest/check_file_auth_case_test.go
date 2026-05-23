//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-manifest
//ff:what runCheckFileAuth — TestCheckFileAuth table-driven 개별 케이스 검증

package hurl_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runCheckFileAuth(t *testing.T, c TestCheckFileAuthCase) {
	t.Helper()
	diags := checkFileAuth(c.entries)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[XOH-06]") {
			t.Errorf("message should contain [XOH-06], got %q", d.Message)
		}
	}
}
