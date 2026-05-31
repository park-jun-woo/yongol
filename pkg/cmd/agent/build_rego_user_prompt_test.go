//ff:func feature=agent type=test control=sequence
//ff:what TestBuildRegoUserPrompt — Table 유무에 따른 resource 결정(테이블/경로유래) 검증
package agent

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestBuildRegoUserPrompt(t *testing.T) {
	feats := []features.Feature{
		{Op: "DeleteWorkflow", Table: "workflows", Path: "/v1/workflows/{id}"},
		{Op: "GetProject", Table: "", Path: "/v1/projects/{id}"}, // no Table → domainFromPath
	}
	got := buildRegoUserPrompt(feats)

	if !strings.Contains(got, "Non-public features requiring authorization:") {
		t.Errorf("missing header, got:\n%s", got)
	}
	// Explicit Table is used verbatim.
	if !strings.Contains(got, "op: DeleteWorkflow, resource: workflows") {
		t.Errorf("expected table-derived resource, got:\n%s", got)
	}
	// Empty Table falls back to the path-derived domain (non-empty).
	if !strings.Contains(got, "op: GetProject, resource: ") {
		t.Errorf("expected GetProject entry, got:\n%s", got)
	}
	wantDomain := domainFromPath("/v1/projects/{id}")
	if !strings.Contains(got, "op: GetProject, resource: "+wantDomain) {
		t.Errorf("expected path-derived resource %q, got:\n%s", wantDomain, got)
	}
	if !strings.Contains(got, "Generate OPA Rego allow rules") {
		t.Errorf("missing instruction, got:\n%s", got)
	}
}
