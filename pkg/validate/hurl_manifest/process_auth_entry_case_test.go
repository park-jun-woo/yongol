//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-manifest
//ff:what runProcessAuthEntry — TestProcessAuthEntry table-driven 개별 케이스 검증

package hurl_manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runProcessAuthEntry(t *testing.T, c TestProcessAuthEntryCase) {
	t.Helper()
	var diags []diagnostic.Diagnostic
	gotAuth := processAuthEntry(c.entry, c.authIssued, &diags)
	if gotAuth != c.wantAuth {
		t.Errorf("authIssued = %v, want %v", gotAuth, c.wantAuth)
	}
	if len(diags) != c.wantDiagCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantDiagCount)
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
