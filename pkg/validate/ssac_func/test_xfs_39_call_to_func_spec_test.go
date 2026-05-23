//ff:func feature=validate type=test control=sequence topic=func-check
//ff:what XFS-39 test — @call builtin package 함수 존재 검증 + advice 메시지 확인

package ssac_func

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/ground"
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXfs39_BuiltinNonExistent verifies that XFS-39 emits an ERROR when a
// @call targets a non-existent function in a builtin ssac package, and the
// advice lists available functions.
func TestXfs39_BuiltinNonExistent(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "RefreshToken",
			FileName: "service/auth/refreshtoken.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "auth.IssueTokenFromClaims",
				Line:  14,
			}},
		}},
		YongolPkgSpecs: []funcspec.FuncSpec{
			{Package: "auth", Name: "hashPassword"},
			{Package: "auth", Name: "verifyPassword"},
			{Package: "auth", Name: "logout"},
		},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs39CallToFuncSpec(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d (%+v)", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XFS-39]") {
		t.Errorf("expected [XFS-39] in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "auth.IssueTokenFromClaims") {
		t.Errorf("expected auth.IssueTokenFromClaims in message, got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Advice, "Available auth functions:") {
		t.Errorf("expected 'Available auth functions:' in advice, got %q", diags[0].Advice)
	}
	if !strings.Contains(diags[0].Advice, "HashPassword") {
		t.Errorf("expected HashPassword in advice, got %q", diags[0].Advice)
	}
}

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

// TestXfs39_NonBuiltinFallbackAdvice verifies that XFS-39 advice for a
// non-builtin package falls back to the generic "Define function under pkg/"
// message.
func TestXfs39_NonBuiltinFallbackAdvice(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "HandleOrder",
			FileName: "service/order/handle.ssac",
			Sequences: []parsessac.Sequence{{
				Type:  "call",
				Model: "billing.DoSomething",
				Line:  5,
			}},
		}},
	}
	fs.SetGround(ground.Build(fs))
	diags := xfs39CallToFuncSpec(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Advice, "Define function") {
		t.Errorf("expected generic advice, got %q", diags[0].Advice)
	}
}
