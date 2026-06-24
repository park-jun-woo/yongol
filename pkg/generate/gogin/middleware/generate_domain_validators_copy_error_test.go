//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerateDomainValidators_CopyError — 누락 스펙 복사 실패가 에러로 표면화되는지 검증 (BUG-142)

package middleware

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateDomainValidators_CopyError(t *testing.T) {
	// Missing spec file → copyFile fails → error surfaced (not a silent skip).
	fs := &yongol.Fullstack{
		SpecsDir: t.TempDir(),
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/app"},
			Domains: map[string]pmanifest.DomainConfig{
				"public": {OpenAPI: "api/missing.yaml", RoutePrefix: "/api"},
			},
		},
	}
	if err := generateDomainValidators(fs, t.TempDir()); err == nil {
		t.Fatal("expected error for missing spec, got nil")
	}
}
