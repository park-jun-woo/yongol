//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDS80_Negative — admin에 보안 있을 때 XDS-80 진단 없음

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDS80_Negative_AdminWithSecurity(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers", Security: "[{bearerAuth: []}]"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XDS-80]") {
		t.Errorf("unexpected XDS-80 diagnostic, got %v", diags)
	}
}
