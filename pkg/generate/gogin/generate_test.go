//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — 첫 단계(oapi-codegen) 실패 시 래핑 에러 반환 검증

package gogin

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestGenerate_OpenAPIStepErrors verifies that when the very first pipeline
// step (oapi-codegen strict-server) fails — because specs/api/openapi.yaml is
// absent or the oapi-codegen binary is unavailable — Generate surfaces the
// wrapped error and aborts before any downstream stage runs.
func TestGenerate_OpenAPIStepErrors(t *testing.T) {
	fs := &yongol.Fullstack{
		SpecsDir: t.TempDir(), // no api/openapi.yaml present
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}
	err := Generate(fs, t.TempDir())
	if err == nil || !strings.Contains(err.Error(), "oapi-codegen strict-server") {
		t.Fatalf("expected oapi-codegen step error, got: %v", err)
	}
}
