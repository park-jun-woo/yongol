//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXMO20_Negative_AuthEndpointExcluded — auth endpoint는 XMO-20 제외

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO20_Negative_AuthEndpointExcluded(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/users":      {"get": {ID: "ListUsers"}},
			"/auth/login": {"post": {ID: "Login", Security: "[]"}},
		}),
	}
	pages := []stml.PageSpec{{
		FileName: "frontend/users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
	}}
	fs := makeMultiDomainFS(domains, opFiles, pages, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XMO-20]") {
		t.Errorf("unexpected XMO-20 for auth endpoint, got %v", diags)
	}
}
