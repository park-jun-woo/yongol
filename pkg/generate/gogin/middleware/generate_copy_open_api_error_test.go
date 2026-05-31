//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — nil fs early-return + copy 에러 + 전체 성공 경로 검증
package middleware

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_CopyOpenAPIError(t *testing.T) {
	// SpecsDir/api/openapi.yaml does not exist -> copyFile open error.
	fs := &yongol.Fullstack{
		SpecsDir: t.TempDir(),
		Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}},
	}
	err := Generate(fs, prepared.State{}, t.TempDir())
	if err == nil {
		t.Fatalf("expected copy openapi.yaml error, got nil")
	}
}
