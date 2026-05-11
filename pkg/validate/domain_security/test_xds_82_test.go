//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what XDS-82 test — public DELETE endpoint without admin Rego rule triggers ERROR
package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
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
	// No Rego policies at all.
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

func TestXDS82_Negative_PublicDeleteWithRego(t *testing.T) {
	domains := map[string]manifest.DomainConfig{
		"public": {OpenAPI: "api/public.yaml", Frontend: "frontend/"},
	}
	opFiles := map[string]string{
		"api/public.yaml": minimalOpenAPI(map[string]map[string]opDef{
			"/workflows/{id}": {"delete": {ID: "DeleteWorkflow"}},
		}),
	}
	policies := []rego.Policy{{
		File: "policy/workflow.rego",
		Rules: []rego.AllowRule{{
			Actions:  []string{"delete"},
			Resource: "workflow",
		}},
	}}
	fs := makeMultiDomainFS(domains, opFiles, nil, policies)
	defer cleanupFS(fs)

	diags := Run(fs)
	if hasDiag(diags, "[XDS-82]") {
		t.Errorf("unexpected XDS-82 diagnostic, got %v", diags)
	}
}

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
