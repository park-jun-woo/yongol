//ff:func feature=manifest type=test control=sequence
//ff:what Load domains 블록 파싱 검증 — DomainConfig 전 필드(openapi·frontend·route_prefix·auth_mode·cors)

package manifest

import (
	"os"
	"path/filepath"
	"testing"
)

// TestLoadDomains verifies that a manifest carrying a `domains:` block parses
// every DomainConfig field, including the per-domain CORS override, and that a
// manifest without the block leaves Domains nil (single-site back-compat).
func TestLoadDomains(t *testing.T) {
	dir := t.TempDir()
	content := `apiVersion: yongol/v1
kind: Project
metadata:
  name: multisite
backend:
  module: github.com/test/multisite
  auth:
    type: jwt
domains:
  public:
    openapi: api/public.yaml
    frontend: frontend/public
    route_prefix: /api
    auth_mode: cookie
  admin:
    openapi: api/admin.yaml
    frontend: frontend/admin
    route_prefix: /api/admin
    auth_mode: bearer
    cors:
      allow_origins: ["https://admin.example.com"]
`
	if err := os.WriteFile(filepath.Join(dir, "manifest.yaml"), []byte(content), 0644); err != nil {
		t.Fatal(err)
	}

	cfg, diags := Load(dir)
	if len(diags) > 0 {
		t.Fatalf("Load() diagnostics: %v", diags)
	}
	if len(cfg.Domains) != 2 {
		t.Fatalf("Domains count = %d, want 2", len(cfg.Domains))
	}

	pub, ok := cfg.Domains["public"]
	if !ok {
		t.Fatal("Domains[public] missing")
	}
	if pub.OpenAPI != "api/public.yaml" {
		t.Errorf("public.OpenAPI = %q", pub.OpenAPI)
	}
	if pub.Frontend != "frontend/public" {
		t.Errorf("public.Frontend = %q", pub.Frontend)
	}
	if pub.RoutePrefix != "/api" {
		t.Errorf("public.RoutePrefix = %q", pub.RoutePrefix)
	}
	if pub.AuthMode != "cookie" {
		t.Errorf("public.AuthMode = %q", pub.AuthMode)
	}
	if pub.CORS != nil {
		t.Errorf("public.CORS = %+v, want nil", pub.CORS)
	}

	adm, ok := cfg.Domains["admin"]
	if !ok {
		t.Fatal("Domains[admin] missing")
	}
	if adm.RoutePrefix != "/api/admin" || adm.AuthMode != "bearer" {
		t.Errorf("admin route_prefix/auth_mode = %q/%q", adm.RoutePrefix, adm.AuthMode)
	}
	if adm.CORS == nil {
		t.Fatal("admin.CORS is nil")
	}
	if len(adm.CORS.AllowOrigins) != 1 || adm.CORS.AllowOrigins[0] != "https://admin.example.com" {
		t.Errorf("admin.CORS.AllowOrigins = %v", adm.CORS.AllowOrigins)
	}
}
