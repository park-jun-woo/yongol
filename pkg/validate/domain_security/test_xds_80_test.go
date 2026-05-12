//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDS80_Positive — admin security: [] 시 XDS-80 ERROR

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDS80_Positive_AdminPublicAccess(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers", Security: "[]"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XDS-80]") {
		t.Errorf("expected XDS-80 diagnostic, got %v", diags)
	}
	if diagLevel(diags, "[XDS-80]") != diagnostic.LevelError {
		t.Errorf("expected ERROR level for XDS-80, got %v", diagLevel(diags, "[XDS-80]"))
	}
}
