//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-71 test — @auth 전에 currentUser 참조 시 ERROR 발생

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS71CurrentUserBeforeAuth(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "GetProject",
			FileName: "project.ssac.go",
			Sequences: []parsessac.Sequence{
				{
					Type:   parsessac.SeqGet,
					Model:  "Project.FindByOwner",
					Inputs: map[string]string{"OwnerID": "currentUser.ID"},
					Result: &parsessac.Result{Type: "Project", Var: "project"},
					Line:   5,
				},
				{
					Type:     parsessac.SeqAuth,
					Action:   "read",
					Resource: "project",
					Line:     6,
				},
			},
		}},
	}
	diags := s71UnknownVariable(fs)
	if len(diags) != 1 {
		t.Fatalf("diag count: got %d want 1 (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, `"currentUser"`) {
		t.Errorf("message should mention currentUser: got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "[S-71]") {
		t.Errorf("message should contain [S-71]: got %q", diags[0].Message)
	}
}
