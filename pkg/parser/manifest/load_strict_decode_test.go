//ff:func feature=manifest type=test control=sequence topic=rate-limit
//ff:what Load — rate_limit 의 모르는 키(requests/window)가 엄격 디코딩으로 parse ERROR 가 되는지 검증

package manifest

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func TestLoad_UnknownRateLimitKeyIsParseError(t *testing.T) {
	dir := t.TempDir()
	// `requests`/`window` are not part of the rate_limit schema (rate/period/key).
	// Strict decoding must reject them instead of silently dropping (BUG-115).
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: strict
backend:
  module: github.com/test/strict
  auth:
    type: jwt
  rate_limit:
    Login:
      requests: 10
      window: 60
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0o644); err != nil {
		t.Fatalf("write fixture: %v", err)
	}

	cfg, diags := Load(dir)
	if len(diags) == 0 {
		t.Fatalf("expected a parse ERROR for unknown rate_limit keys, got 0 diagnostics (cfg=%+v)", cfg)
	}
	d := diags[0]
	if d.Phase != diagnostic.PhaseParse {
		t.Errorf("Phase = %v, want PhaseParse", d.Phase)
	}
	if d.Level != diagnostic.LevelError {
		t.Errorf("Level = %v, want LevelError", d.Level)
	}
	if !strings.Contains(d.Message, "field requests") {
		t.Errorf("Message should name the unknown field, got %q", d.Message)
	}
}
