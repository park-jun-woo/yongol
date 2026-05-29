//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-71 test — @get 결과 변수를 후속 step 에서 참조 (PASS)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS71StepResultVarPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "UpdateUser",
			FileName: "user.ssac.go",
			Sequences: []parsessac.Sequence{
				{
					Type:   parsessac.SeqGet,
					Model:  "User.FindByID",
					Inputs: map[string]string{"ID": "request.Id"},
					Result: &parsessac.Result{Type: "User", Var: "user"},
					Line:   5,
				},
				{
					Type:   parsessac.SeqPut,
					Model:  "User.Update",
					Inputs: map[string]string{"ID": "user.ID", "Name": "request.Name"},
					Line:   6,
				},
			},
		}},
	}
	diags := s71UnknownVariable(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diags, got %d: %+v", len(diags), diags)
	}
}
