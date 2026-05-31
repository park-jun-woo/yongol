//ff:func feature=generate type=test control=sequence
//ff:what Generate 오케스트레이터의 migration no-op + backend 에러 경로 검증
package generate

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_UnknownBackendErrors(t *testing.T) {
	// Empty SpecsDir makes runMigrationStep a no-op; an unknown backend then
	// surfaces a wrapped "backend:" error before frontend/hurl/opa steps run.
	fs := &yongol.Fullstack{SpecsDir: ""}
	err := Generate(fs, t.TempDir(), BackendType("does-not-exist"), FrontendType("react"))
	if err == nil || !strings.Contains(err.Error(), "backend:") {
		t.Fatalf("expected backend error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "unknown backend") {
		t.Errorf("expected 'unknown backend' in error, got: %v", err)
	}
}
