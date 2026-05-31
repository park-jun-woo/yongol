//ff:func feature=agent type=test control=sequence
//ff:what TestFixSplitFile — operationId 부재 에러 / op 추출 후 per-op 루프+writeFile 분기 검증
package agent

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixSplitFileNoOpID(t *testing.T) {
	// Diagnostics without any operationId reference → early "no operationId" error.
	dir := t.TempDir()
	abs := filepath.Join(dir, "api", "openapi.yaml")
	err := fixSplitFile(dir, &features.FeaturesFile{}, "api/openapi.yaml", abs, "paths: {}\n",
		[]diagnostic.Diagnostic{{Message: "generic error"}},
		func(b, m, s, u string) (string, error) { return "", nil }, Config{}, layerOpenAPI)
	if err == nil {
		t.Fatal("expected error when no operationId present")
	}
	if !strings.Contains(err.Error(), "no operationId") {
		t.Errorf("expected no-operationId error, got: %v", err)
	}
}
