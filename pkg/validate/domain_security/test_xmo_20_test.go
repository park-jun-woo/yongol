//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what XMO-20 test — public OpenAPI operationId가 public STML에서 미소비되면 ERROR
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO20_Positive_PublicUnconsumed(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/users":      {"get": {ID: "ListUsers"}},
			"/users/{id}": {"get": {ID: "GetUser"}},
		}),
	}
	// Only ListUsers is consumed; GetUser is not.
	pages := []stml.PageSpec{{
		FileName: "frontend/users.html",
		Fetches: []stml.FetchBlock{{
			OperationID: "ListUsers",
		}},
	}}
	fs := makeMultiDomainFS(domains, opFiles, pages, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XMO-20]") {
		t.Errorf("expected XMO-20 diagnostic for GetUser, got %v", diags)
	}
	if diagLevel(diags, "[XMO-20]") != diagnostic.LevelError {
		t.Errorf("expected ERROR level for XMO-20, got %v", diagLevel(diags, "[XMO-20]"))
	}
}

func TestXMO20_Negative_AllConsumed(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/users": {"get": {ID: "ListUsers"}},
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
		t.Errorf("unexpected XMO-20 diagnostic, got %v", diags)
	}
}

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
