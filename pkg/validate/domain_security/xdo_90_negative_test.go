//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDO90_Negative — 고유 operationId 시 XDO-90 진단 없음

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDO90_Negative_UniqueOpIDs(t *testing.T) {
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
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XDO-90]") {
		t.Errorf("unexpected XDO-90 diagnostic, got %v", diags)
	}
}
