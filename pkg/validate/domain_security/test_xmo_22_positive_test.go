//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXMO22_Positive — 도메인 경계 위반 호출 시 XMO-22 WARNING

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
