//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXMO21_Negative_NoAdminFrontend — admin frontend 없으면 XMO-21 skip

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXMO21_Negative_NoAdminFrontend(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"admin": {OpenAPI: "api/admin.yaml"},
	}
	opFiles := map[string]string{
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListAdminUsers"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XMO-21]") {
		t.Errorf("unexpected XMO-21 diagnostic when admin has no frontend, got %v", diags)
	}
}
