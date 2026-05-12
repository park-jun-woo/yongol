//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDS82_Positive — public DELETE Rego 없을 때 XDS-82 ERROR

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestXDS82_Positive_PublicDeleteNoRego(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/workflows/{id}": {"delete": {ID: "DeleteWorkflow"}},
		}),
	}
	fs := makeMultiDomainFS(domains, opFiles, nil, nil)
	defer cleanupFS(fs)

	diags := Run(fs)
	if !hasDiag(diags, "[XDS-82]") {
		t.Errorf("expected XDS-82 diagnostic, got %v", diags)
	}
	if diagLevel(diags, "[XDS-82]") != diagnostic.LevelError {
		t.Errorf("expected ERROR level for XDS-82, got %v", diagLevel(diags, "[XDS-82]"))
	}
}
