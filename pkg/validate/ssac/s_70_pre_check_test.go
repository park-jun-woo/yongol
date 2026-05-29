//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what S-70 사전 점검 — 기존 S-29/S-33/S-59 가 @post 단독 reserved source 를 막지 않음을 증명 (silent pass)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestS70PreCheckSilentPass demonstrates that prior to S-70, a standalone
// reserved source (currentUser) used as an @post Inputs value is *not* caught
// by S-29 (currentUser is isImplicitVar → exempt), S-33 (only checks Result
// variable name, not Inputs), or S-59 (only checks dotted forms; standalone
// has no dot and is skipped). This Case-B evidence justifies introducing
// S-70.
func TestS70PreCheckSilentPass(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "Register",
			FileName: "register.ssac.go",
			Sequences: []parsessac.Sequence{
				{
					Type:   parsessac.SeqPost,
					Model:  "User.Create",
					Inputs: map[string]string{"Claims": "currentUser"},
					Result: &parsessac.Result{Type: "User", Var: "u"},
					Line:   10,
				},
			},
		}},
	}
	if d := s29InputDeclared(fs); len(d) != 0 {
		t.Errorf("S-29 unexpectedly fired: %v", d)
	}
	if d := s33ReservedSource(fs); len(d) != 0 {
		t.Errorf("S-33 unexpectedly fired: %v", d)
	}
	if d := s59DottedField(fs); len(d) != 0 {
		t.Errorf("S-59 unexpectedly fired: %v", d)
	}
}
