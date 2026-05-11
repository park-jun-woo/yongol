//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what XMO-22 test — STML에서 다른 도메인의 operationId 호출 시 WARNING
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO22_Positive_CrossDomainCall(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
		"admin":  {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/users": {"get": {ID: "ListUsers"}},
		}),
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers"}},
		}),
	}
	// Public frontend page calls admin operationId.
	pages := []stml.PageSpec{{
		FileName: "frontend/users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListAdminUsers",
		}},
	}}
	fs := makeMultiDomainFS(domains, opFiles, pages, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XMO-22]") {
		t.Errorf("expected XMO-22 diagnostic for cross-domain call, got %v", diags)
	}
	if diagLevel(diags, "[XMO-22]") != diagnostic.LevelWarning {
		t.Errorf("expected WARNING level for XMO-22, got %v", diagLevel(diags, "[XMO-22]"))
	}
}

func TestXMO22_Negative_SameDomainCall(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
		"admin":  {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/users": {"get": {ID: "ListUsers"}},
		}),
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers"}},
		}),
	}
	// Public frontend page calls public operationId — no violation.
	pages := []stml.PageSpec{{
		FileName: "frontend/users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
	}}
	fs := makeMultiDomainFS(domains, opFiles, pages, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XMO-22]") {
		t.Errorf("unexpected XMO-22 diagnostic for same-domain call, got %v", diags)
	}
}
