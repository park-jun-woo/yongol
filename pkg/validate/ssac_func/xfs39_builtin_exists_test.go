//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what TestXfs39_BuiltinExists — @call 대상이 존재하는 builtin 함수면 에러 미발생 확인

package ssac_func

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs39_BuiltinExists verifies that XFS-39 does NOT emit an error when a
// @call targets an existing function in a builtin ssac package.
func TestXfs39_BuiltinExists(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "Login",
			FileName: "service/auth/login.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "auth.VerifyPassword",
				Line:  10,
			}},
		}},
		YongolPkgSpecs: []funcspec.FuncSpec{
			{Package: "auth", Name: "verifyPassword"},
		},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs39CallToFuncSpec(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics for valid builtin call, got %d (%+v)", len(diags), diags)
	}
}
