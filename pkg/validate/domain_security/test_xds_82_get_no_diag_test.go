//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDS82_Negative_PublicGetNoDiag — GET endpoint는 XDS-82 대상 아님

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDS82_Negative_PublicGetNoDiag(t *testing.T) {
	// Non-DELETE methods should not trigger XDS-82.
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/workflows": {"get": {ID: "ListWorkflows"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XDS-82]") {
		t.Errorf("unexpected XDS-82 diagnostic for GET endpoint, got %v", diags)
	}
}
