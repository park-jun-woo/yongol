//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXMO21_Positive — admin 미소비 operationId 시 XMO-21 ERROR

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO21_Positive_AdminUnconsumed(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers"}},
			"/admin/logs":  {"get": {ID: "GetAdminLogs"}},
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
	if !hasDiag(diags, "[XMO-21]") {
		t.Errorf("expected XMO-21 diagnostic for GetAdminLogs, got %v", diags)
	}
	if diagLevel(diags, "[XMO-21]") != diagnostic.LevelError {
		t.Errorf("expected ERROR level for XMO-21, got %v", diagLevel(diags, "[XMO-21]"))
	}
}
