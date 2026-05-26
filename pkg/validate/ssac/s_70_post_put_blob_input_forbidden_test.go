//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what S-70 test — @post/@put Inputs 단독 reserved source 거절 (6 cases)

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS70PostPutBlobInputForbidden(t *testing.T) {
	tt := []struct {
		name      string
		seqs      []parsessac.Sequence
		wantDiag  int
		wantSubst string
	}{
		{
			name: "1-dotted-OK: @post Email=currentUser.Email",
			seqs: []parsessac.Sequence{{
				Type:   parsessac.SeqPost,
				Model:  "User.Create",
				Inputs: map[string]string{"Email": "currentUser.Email"},
				Result: &parsessac.Result{Type: "User", Var: "u"},
				Line:   10,
			}},
			wantDiag: 0,
		},
		{
			name: "2-standalone-currentUser: @post Claims=currentUser → ERROR",
			seqs: []parsessac.Sequence{{
				Type:   parsessac.SeqPost,
				Model:  "User.Create",
				Inputs: map[string]string{"Claims": "currentUser"},
				Result: &parsessac.Result{Type: "User", Var: "u"},
				Line:   11,
			}},
			wantDiag:  1,
			wantSubst: `reserved source "currentUser"`,
		},
		{
			name: "3-standalone-request: @put Meta=request → ERROR",
			seqs: []parsessac.Sequence{{
				Type:   parsessac.SeqPut,
				Model:  "Workflow.UpdateMeta",
				Inputs: map[string]string{"ID": "wf.ID", "Meta": "request"},
				Line:   12,
			}},
			wantDiag:  1,
			wantSubst: `reserved source "request"`,
		},
		{
			name: "4-call-exempt: @call Subject=currentUser → 0 (예외)",
			seqs: []parsessac.Sequence{{
				Type:    parsessac.SeqCall,
				Package: "audit",
				Model:   "audit.Log",
				Inputs:  map[string]string{"Subject": "currentUser"},
				Line:    13,
			}},
			wantDiag: 0,
		},
		{
			name: "5-get-dotted-OK: @get FindByID with request.id",
			seqs: []parsessac.Sequence{{
				Type:   parsessac.SeqGet,
				Model:  "Model.FindByID",
				Inputs: map[string]string{"ID": "request.id"},
				Result: &parsessac.Result{Type: "Model", Var: "m"},
				Line:   14,
			}},
			wantDiag: 0,
		},
		{
			name: "6-mixed: @post Claims=currentUser + Email=request.email → 1 (Claims만)",
			seqs: []parsessac.Sequence{{
				Type:   parsessac.SeqPost,
				Model:  "User.Create",
				Inputs: map[string]string{"Claims": "currentUser", "Email": "request.email"},
				Result: &parsessac.Result{Type: "User", Var: "u"},
				Line:   15,
			}},
			wantDiag:  1,
			wantSubst: `reserved source "currentUser"`,
		},
	}

	for _, tc := range tt {
		fs := &yongol.Fullstack{
			ServiceFuncs: []parsessac.ServiceFunc{{
				Name:      "TestFunc",
				FileName:  "test.ssac.go",
				Sequences: tc.seqs,
			}},
		}
		diags := s70PostPutBlobInputForbidden(fs)
		if len(diags) != tc.wantDiag {
			t.Errorf("[%s] diag count: got %d want %d (%+v)", tc.name, len(diags), tc.wantDiag, diags)
			continue
		}
		if tc.wantDiag > 0 && tc.wantSubst != "" && !strings.Contains(diags[0].Message, tc.wantSubst) {
			t.Errorf("[%s] message: got %q want substring %q", tc.name, diags[0].Message, tc.wantSubst)
		}
	}
}
