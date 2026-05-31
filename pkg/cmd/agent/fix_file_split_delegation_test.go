//ff:func feature=agent type=test control=sequence
//ff:what TestFixFile — read 에러 / 전체파일 수정 성공·LLM에러·빈응답 / split 레이어 위임 분기 검증
package agent

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestFixFileSplitDelegation(t *testing.T) {
	// An OpenAPI file routes to fixSplitFile. With diagnostics carrying no
	// operationId, fixSplitFile returns its "no operationId" error, proving the
	// split-layer delegation branch was taken.
	dir := t.TempDir()
	apiDir := filepath.Join(dir, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(apiDir, "openapi.yaml"), []byte("paths: {}\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	rel := filepath.Join("api", "openapi.yaml")
	diags := []diagnostic.Diagnostic{{Message: "some error", Level: diagnostic.LevelError}}
	llm := func(b, m, s, u string) (string, error) { return "x", nil }
	err := fixFile(dir, &features.FeaturesFile{}, rel, diags, llm, Config{})
	if err == nil {
		t.Fatal("expected delegation to fixSplitFile to surface an error")
	}
	if !strings.Contains(err.Error(), "operationId") {
		t.Errorf("expected fixSplitFile no-operationId error, got: %v", err)
	}
}
