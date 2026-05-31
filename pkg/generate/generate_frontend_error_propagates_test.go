//ff:func feature=generate type=test control=sequence
//ff:what Generate 오케스트레이터의 migration no-op + backend 에러 경로 검증
package generate

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_FrontendErrorPropagates(t *testing.T) {
	// FastAPI backend succeeds on an empty Fullstack (writes scaffold), so an
	// unknown frontend then surfaces a wrapped "frontend:" error before the
	// hurl/opa steps run.
	fs := &yongol.Fullstack{SpecsDir: ""}
	err := Generate(fs, t.TempDir(), FastAPI, FrontendType("does-not-exist"))
	if err == nil || !strings.Contains(err.Error(), "frontend:") {
		t.Fatalf("expected frontend error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "unknown frontend") {
		t.Errorf("expected 'unknown frontend' in error, got: %v", err)
	}
}
