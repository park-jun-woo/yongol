//ff:func feature=validate type=test-helper control=iteration dimension=2 topic=funcspec-structural
//ff:what runRun — TestRun table-driven 개별 케이스 검증

package funcspec

import (
	"strings"
	"testing"
)

func runRun(t *testing.T, c TestRunCase) {
	t.Helper()
	diags := Run(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), c.wantCount, diags)
	}
	for _, code := range c.wantCodes {
		found := false
		for _, d := range diags {
			if strings.Contains(d.Message, code) {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("expected diagnostic containing %q not found", code)
		}
	}
}
