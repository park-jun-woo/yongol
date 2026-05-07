//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-71 test — request.Id 참조는 항상 유효 (PASS)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS71RequestVarPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "GetUser",
			FileName: "user.ssac.go",
			Sequences: []parsessac.Sequence{
				{
					Type:   parsessac.SeqGet,
					Model:  "User.FindByID",
					Inputs: map[string]string{"ID": "request.Id"},
					Result: &parsessac.Result{Type: "User", Var: "user"},
					Line:   5,
				},
			},
		}},
	}
	diags := s71UnknownVariable(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diags, got %d: %+v", len(diags), diags)
	}
}
