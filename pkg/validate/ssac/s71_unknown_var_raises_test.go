//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-71 test — 미선언 변수 query.page 참조 시 ERROR 발생

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS71UnknownVarRaises(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "ListWorkflows",
			FileName: "workflow.ssac.go",
			Sequences: []parsessac.Sequence{
				{
					Type:     parsessac.SeqAuth,
					Action:   "list",
					Resource: "workflow",
					Line:     5,
				},
				{
					Type:   parsessac.SeqGet,
					Model:  "Workflow.ListByOrgID",
					Inputs: map[string]string{"OrgID": "currentUser.OrgID", "Page": "query.page"},
					Result: &parsessac.Result{Type: "[]Workflow", Var: "items"},
					Line:   6,
				},
			},
		}},
	}
	diags := s71UnknownVariable(fs)
	if len(diags) != 1 {
		t.Fatalf("diag count: got %d want 1 (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, `"query"`) {
		t.Errorf("message should mention query: got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "[S-71]") {
		t.Errorf("message should contain [S-71]: got %q", diags[0].Message)
	}
}
