//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what S-70 반증 검증 — Run() 에 s70 호출이 없으면 standalone reserved source 가 silent 통과함을 증명

package ssac

import (
	"strings"
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS70DisproveViaRun(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "Register",
			FileName: "register.ssac.go",
			Sequences: []parsessac.Sequence{{
				Type:   parsessac.SeqPost,
				Model:  "User.Create",
				Inputs: map[string]string{"Claims": "currentUser"},
				Result: &parsessac.Result{Type: "User", Var: "u"},
				Line:   10,
			}},
		}},
	}
	diags := Run(fs)
	hit := 0
	for _, d := range diags {
		if strings.Contains(d.Message, "[S-70]") {
			hit++
		}
	}
	if hit != 1 {
		t.Fatalf("expected exactly one [S-70] diag from Run; got %d. all diags: %+v", hit, diags)
	}
}
