//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-observability
//ff:what runObs02MetricsPathNotOpenAPI — TestObs02MetricsPathNotOpenAPI table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runObs02MetricsPathNotOpenAPI(t *testing.T, c TestObs02MetricsPathNotOpenAPICase) {
	t.Helper()
	diags := obs02MetricsPathNotOpenAPI(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[OBS-002]") {
			t.Errorf("expected [OBS-002], got %q", d.Message)
		}
	}
}
