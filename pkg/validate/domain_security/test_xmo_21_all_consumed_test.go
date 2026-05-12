//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXMO21_Negative_AllAdminConsumed — admin 전체 소비 시 XMO-21 진단 없음

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO21_Negative_AllAdminConsumed(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers"}},
		}),
	}
	pages := []stml.PageSpec{{
		FileName: "admin-frontend/users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListAdminUsers",
		}},
	}}
	fs := makeMultiDomainFS(domains, opFiles, pages, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XMO-21]") {
		t.Errorf("unexpected XMO-21 diagnostic when all consumed, got %v", diags)
	}
}
