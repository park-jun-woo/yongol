//ff:func feature=agent type=test control=iteration dimension=1
//ff:what TestFixSplitFile — operationId 부재 에러 / op 추출 후 per-op 루프+writeFile 분기 검증

package agent

import (
	"os"
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

func TestFixSplitFileWithOpID(t *testing.T) {
	// A diagnostic referencing an operationId drives the per-op loop. The op-level
	// fix fails (unsupported backend → llmCall error inside fixSingleBlock), the
	// loop continues, and the (unchanged) content is written back to absPath.
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	abs := filepath.Join(apiDir, "openapi.yaml")
	content := "paths:\n  /users:\n    get:\n      operationId: ListUsers\n"
	if err := os.WriteFile(abs, []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}

	diags := []diagnostic.Diagnostic{{Message: "operationId: ListUsers has an error"}}
	err := fixSplitFile(dir, &features.FeaturesFile{}, "api/openapi.yaml", abs, content, diags,
		func(b, m, s, u string) (string, error) { return "", nil },
		Config{Backend: "unsupported-backend", Model: "none"}, layerOpenAPI)
	if err != nil {
		t.Fatalf("expected nil error after writing content back, got: %v", err)
	}
	// File should still exist and be readable.
	if _, statErr := os.Stat(abs); statErr != nil {
		t.Errorf("expected file written back: %v", statErr)
	}
}
