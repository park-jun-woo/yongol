//ff:func feature=chain type=test control=sequence
//ff:what TestTracePolicy — @auth resource ↔ Rego rule 매칭 + action 필터링 검증
package chain

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestTracePolicy(t *testing.T) {
	specsDir, regoFile := tracePolicySetup(t)
	policies := []rego.Policy{
		{
			File: regoFile,
			Rules: []rego.AllowRule{
				{Resource: "project", Actions: []string{"delete", "update"}},
				{Resource: "other", Actions: []string{"read"}}, // not referenced
			},
		},
	}
	sf := &ssac.ServiceFunc{
		Name: "DeleteProject",
		Sequences: []ssac.Sequence{
			{Type: "auth", Resource: "project", Action: "delete"},
			{Type: "get", Model: "Project.FindByID"}, // ignored
		},
	}

	links := tracePolicy(sf, policies, specsDir)
	if len(links) != 1 {
		t.Fatalf("expected 1 rego link, got %d: %+v", len(links), links)
	}
	l := links[0]
	if l.Kind != "Rego" || l.File != "authz/project.rego" {
		t.Errorf("link fields: %+v", l)
	}
	if l.Line != 1 {
		t.Errorf("line: got %d, want 1", l.Line)
	}
	if l.Summary != "resource: project [delete]" {
		t.Errorf("summary: got %q, want %q", l.Summary, "resource: project [delete]")
	}
}
