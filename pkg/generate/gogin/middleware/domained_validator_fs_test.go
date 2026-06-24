//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what domainedValidatorFS — 2 도메인(public/admin) + 디스크 스펙 픽스처 (BUG-142 검증)

package middleware

import (
	"os"
	"path/filepath"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// domainedValidatorFS writes per-domain OpenAPI specs to a temp specs dir and
// returns a 2-domain Fullstack (public/admin) pointing at them.
func domainedValidatorFS(t *testing.T) *yongol.Fullstack {
	t.Helper()
	specs := t.TempDir()
	apiDir := filepath.Join(specs, "api")
	if err := os.MkdirAll(apiDir, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	for _, d := range []struct{ file, title string }{
		{"public.yaml", "public"}, {"admin.yaml", "admin"},
	} {
		spec := "openapi: 3.0.0\ninfo:\n  title: " + d.title + "\n  version: '1'\npaths: {}\n"
		if err := os.WriteFile(filepath.Join(apiDir, d.file), []byte(spec), 0o644); err != nil {
			t.Fatalf("setup %s: %v", d.file, err)
		}
	}
	return &yongol.Fullstack{
		SpecsDir: specs,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{Module: "example.com/app"},
			Domains: map[string]pmanifest.DomainConfig{
				"public": {OpenAPI: "api/public.yaml", RoutePrefix: "/api"},
				"admin":  {OpenAPI: "api/admin.yaml", RoutePrefix: "/api/admin"},
			},
		},
	}
}
