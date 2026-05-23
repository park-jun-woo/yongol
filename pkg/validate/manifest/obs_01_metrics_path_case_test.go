//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-observability
//ff:what runObs01MetricsPath — TestObs01MetricsPath table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runObs01MetricsPath(t *testing.T, c TestObs01MetricsPathCase) {
	t.Helper()
	diags := obs01MetricsPath(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[OBS-001]") {
			t.Errorf("expected [OBS-001], got %q", d.Message)
		}
	}
}
