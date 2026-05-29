//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDO90_Positive — 중복 operationId 시 XDO-90 ERROR 진단

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDO90_Positive_DuplicateOpID(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
		"admin":  {OpenAPI: "api/admin.yaml", Frontend: "admin-frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/users": {"get": {ID: "ListUsers"}},
		}),
		"api/admin.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/admin/users": {"get": {ID: "ListUsers"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XDO-90]") {
		t.Errorf("expected XDO-90 diagnostic, got %v", diags)
	}
	if diagLevel(diags, "[XDO-90]") != diagnostic.LevelError {
		t.Errorf("expected ERROR level for XDO-90, got %v", diagLevel(diags, "[XDO-90]"))
	}
}
