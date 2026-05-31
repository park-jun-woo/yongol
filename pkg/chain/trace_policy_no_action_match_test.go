//ff:func feature=chain type=test control=sequence
//ff:what TestTracePolicyNoActionMatch — 빈 action 시 summary 에 action 브래킷이 없음
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTracePolicyNoActionMatch(t *testing.T) {
	specsDir, regoFile := tracePolicySetup(t)
	sfNoAction := &ssac.ServiceFunc{
		Name:      "ListProject",
		Sequences: []ssac.Sequence{{Type: "auth", Resource: "project", Action: ""}},
	}
	pol := []rego.Policy{
		{
			File:  regoFile,
			Rules: []rego.AllowRule{{Resource: "project", Actions: []string{"delete", "update"}}},
		},
	}
	links := tracePolicy(sfNoAction, pol, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 link, got %d: %+v", len(links), links)
	}
	if links[0].Summary != "resource: project" {
		t.Errorf("summary: got %q, want %q (no action brackets)", links[0].Summary, "resource: project")
	}
}
