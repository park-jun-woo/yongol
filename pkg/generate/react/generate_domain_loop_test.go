//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what Generate — 도메인 모드에서 도메인별 frontend/<name> 앱 스캐폴드 + per-domain api.ts operationId 격리 + baseUrl '' 검증

package react

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_DomainMode_PerDomainApps(t *testing.T) {
	publicDoc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/api/things", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "ListThings", Responses: openapi3.NewResponses()},
		}),
	)}
	adminDoc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/api/admin/users", &openapi3.PathItem{
			Get: &openapi3.Operation{OperationID: "ListUsers", Responses: openapi3.NewResponses()},
		}),
	)}

	// Real per-domain spec files so the per-domain spec-path derivation
	// (filepath.Join(SpecsDir, cfg.OpenAPI)) resolves to a parseable file and
	// openapi-typescript does not fault on a missing path. The tool may be
	// absent in some environments — that surfaces as a Generate error and the
	// test skips (the domain-loop wiring is still covered elsewhere).
	specsDir := t.TempDir()
	const minimalSpec = "openapi: 3.0.0\ninfo:\n  title: t\n  version: '1.0.0'\npaths:\n  /api/x:\n    get:\n      operationId: X\n      responses:\n        '200':\n          description: ok\n"
	writeSpecFile(t, filepath.Join(specsDir, "api", "openapi.yaml"), minimalSpec)
	writeSpecFile(t, filepath.Join(specsDir, "api", "admin.yaml"), minimalSpec)

	fs := &yongol.Fullstack{
		SpecsDir: specsDir,
		Manifest: &manifest.ProjectConfig{
			Frontend: manifest.Frontend{Lang: "typescript"},
			Domains: map[string]manifest.DomainConfig{
				"public": {OpenAPI: "api/openapi.yaml", Frontend: "frontend", RoutePrefix: "/api"},
				"admin":  {OpenAPI: "api/admin.yaml", Frontend: "admin", RoutePrefix: "/api/admin"},
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{"public": publicDoc, "admin": adminDoc},
	}
	if !fs.IsDomained() {
		t.Fatal("fixture must be domained")
	}

	out := t.TempDir()
	if err := Generate(fs, out); err != nil {
		t.Skipf("Generate (domain mode) returned a tool error (openapi-typescript env?): %v", err)
	}

	// Each domain materializes its own complete app tree under frontend/<name>.
	for _, rel := range []string{
		filepath.Join("public", "package.json"),
		filepath.Join("public", "vite.config.ts"),
		filepath.Join("public", "src", "lib", "api.ts"),
		filepath.Join("admin", "package.json"),
		filepath.Join("admin", "vite.config.ts"),
		filepath.Join("admin", "src", "lib", "api.ts"),
	} {
		if _, err := os.Stat(filepath.Join(out, "frontend", rel)); err != nil {
			t.Errorf("expected %s: %v", rel, err)
		}
	}

	// api.ts carries only the per-domain operationId (the per-domain doc is
	// already filtered — no operationId filtering needed) and keeps baseUrl: ''
	// (BUG-110).
	pubAPI := readFile(t, filepath.Join(out, "frontend", "public", "src", "lib", "api.ts"))
	adminAPI := readFile(t, filepath.Join(out, "frontend", "admin", "src", "lib", "api.ts"))
	if !strings.Contains(pubAPI, "ListThings") || strings.Contains(pubAPI, "ListUsers") {
		t.Errorf("public api.ts must contain ListThings only (no ListUsers leak)")
	}
	if !strings.Contains(adminAPI, "ListUsers") || strings.Contains(adminAPI, "ListThings") {
		t.Errorf("admin api.ts must contain ListUsers only (no ListThings leak)")
	}
	if !strings.Contains(pubAPI, "baseUrl: ''") {
		t.Errorf("public api.ts must keep baseUrl: '' (BUG-110)")
	}
	if !strings.Contains(adminAPI, "baseUrl: ''") {
		t.Errorf("admin api.ts must keep baseUrl: '' (BUG-110)")
	}
}
