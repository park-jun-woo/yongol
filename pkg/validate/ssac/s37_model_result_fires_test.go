//ff:func feature=validate type=test control=sequence dimension=1 topic=ssac-structural
//ff:what S-37 model fire — @get Workflow FK 참조 + @empty 없음 → S-37 fire 확인

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS37ModelResultFires(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "ActivateWorkflow",
			FileName: "service/activateworkflow.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:   "get",
					Model:  "Organization.FindByID",
					Result: &parsessac.Result{Type: "Organization", Var: "org"},
					Args:   []parsessac.Arg{{Field: "ID", Source: "request"}},
					Line:   10,
				},
				{
					Type:    "empty",
					Target:  "org",
					Message: "Organization not found",
					Line:    11,
				},
				{
					Type:   "get",
					Model:  "Workflow.FindByOrgID",
					Result: &parsessac.Result{Type: "Workflow", Var: "wf"},
					Args:   []parsessac.Arg{{Field: "OrgID", Source: "org"}},
					Line:   12,
				},
			},
		}},
	}
	g := &rule.Ground{
		Lookup: map[string]rule.StringSet{},
		Types:  map[string]string{},
		Pairs:  map[string]rule.StringSet{},
		Config: map[string]bool{},
		Vars:   rule.StringSet{},
		Flags:  rule.StringSet{},
	}
	fs.SetGround(g)

	diags := s37FKReferenceGuard(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag for Model result without @empty, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[S-37]") {
		t.Errorf("expected S-37 prefix, got %q", diags[0].Message)
	}
}
