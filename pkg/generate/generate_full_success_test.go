//ff:func feature=generate type=test control=sequence
//ff:what Generate 오케스트레이터의 migration no-op + backend 에러 경로 검증
package generate

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_FullSuccess(t *testing.T) {
	// Empty Fullstack: migration is a no-op, FastAPI backend + React frontend
	// both succeed, and the hurl-mirror / opa-rego steps run to completion.
	fs := &yongol.Fullstack{SpecsDir: ""}
	if err := Generate(fs, t.TempDir(), FastAPI, React); err != nil {
		t.Fatalf("expected success, got: %v", err)
	}
}
