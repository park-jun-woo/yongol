//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestLoadDomainOpenAPIDocs_NoDisk — loadDomainOpenAPIDocs 가 디스크 재파싱 없이 fs.DomainOpenAPIDocs 에서 조립하고 Cfg 를 보존하는지 검증

package domain_security

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestLoadDomainOpenAPIDocs_NoDisk(t *testing.T) {
	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	fs := &yongol.Fullstack{
		// A path that does not exist on disk: if the function re-parsed from
		// SpecsDir it would find nothing and return zero docs.
		SpecsDir: "/nonexistent/specs/dir",
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/public"},
				// No pre-parsed doc for "admin" → must be skipped.
				"admin": {OpenAPI: "api/admin.yaml", Frontend: "frontend/admin"},
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{
			"public": doc,
		},
	}

	docs := loadDomainOpenAPIDocs(fs)
	if len(docs) != 1 {
		t.Fatalf("expected 1 assembled domainDoc (admin has no pre-parsed doc), got %d", len(docs))
	}
	dd := docs[0]
	if dd.Name != "public" {
		t.Errorf("Name = %q, want public", dd.Name)
	}
	if dd.Doc != doc {
		t.Error("Doc should be the pre-parsed fs.DomainOpenAPIDocs entry, not a re-parsed copy")
	}
	if dd.Cfg.OpenAPI != "api/public.yaml" || dd.Cfg.Frontend != "frontend/public" {
		t.Errorf("Cfg not preserved from manifest: %+v", dd.Cfg)
	}
}
