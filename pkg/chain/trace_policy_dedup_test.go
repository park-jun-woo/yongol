//ff:func feature=chain type=test control=sequence
//ff:what TestTracePolicyDedup — 같은 파일+resource 중복 rule 은 하나의 link 로 dedup
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTracePolicyDedup(t *testing.T) {
	specsDir, regoFile := tracePolicySetup(t)
	sf := &ssac.ServiceFunc{
		Name:      "DeleteProject",
		Sequences: []ssac.Sequence{{Type: "auth", Resource: "project", Action: "delete"}},
	}
	dupPolicies := []rego.Policy{
		{
			File: regoFile,
			Rules: []rego.AllowRule{
				{Resource: "project", Actions: []string{"delete"}},
				{Resource: "project", Actions: []string{"update"}}, // same file+resource
			},
		},
	}
	links := tracePolicy(sf, dupPolicies, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 deduped rego link, got %d: %+v", len(links), links)
	}
	if links[0].File != "authz/project.rego" {
		t.Errorf("link file: got %q, want authz/project.rego", links[0].File)
	}
}
