//ff:func feature=ground type=test control=sequence
//ff:what TestBuild_DomainedMerge — 도메인 2개 OpenAPI 의 operationId/path/security 가 덮어쓰기 없이 합산(MERGE)되는지 검증

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuild_DomainedMerge(t *testing.T) {
	fs := &yongol.Fullstack{
		Manifest: &manifest.ProjectConfig{
			Domains: map[string]manifest.DomainConfig{
				"public": {OpenAPI: "api/public.yaml"},
				"admin":  {OpenAPI: "api/admin.yaml"},
			},
		},
		DomainOpenAPIDocs: map[string]*openapi3.T{
			"public": makeDomainDoc("/users", "ListUsers", "publicAuth"),
			"admin":  makeDomainDoc("/admin/users", "ListAdminUsers", "adminAuth"),
		},
	}

	g := Build(fs)

	ops := g.Lookup["OpenAPI.operationId"]
	if !ops["ListUsers"] || !ops["ListAdminUsers"] {
		t.Fatalf("operationId merge dropped a domain: %v", ops)
	}
	paths := g.Lookup["OpenAPI.path"]
	if !paths["/users"] || !paths["/admin/users"] {
		t.Fatalf("path merge dropped a domain: %v", paths)
	}
	sec := g.Lookup["OpenAPI.security"]
	if !sec["publicAuth"] || !sec["adminAuth"] {
		t.Fatalf("security merge dropped a domain: %v", sec)
	}
}
