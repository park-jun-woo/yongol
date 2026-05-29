//ff:func feature=validate type=test control=sequence topic=domain-security
//ff:what TestXDS82_Negative_WithRego — public DELETE에 Rego 룰 있으면 XDS-82 진단 없음

package domain_security

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
)

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
