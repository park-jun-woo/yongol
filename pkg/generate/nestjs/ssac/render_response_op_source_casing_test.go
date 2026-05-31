//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderResponseOpSourceCasing -- tsSourceExpr 으로 PascalCase → camelCase 변환 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderResponseOpSourceCasing(t *testing.T) {
	op := &ir.ResponseOp{
		Fields: []ir.ResponseField{
			{Name: "access_token", Source: "token.AccessToken"},
			{Name: "refresh_token", Source: "token.RefreshToken"},
			{Name: "email", Source: "user.Email"},
		},
	}

	var b strings.Builder
	renderResponseOp(&b, op, "    ")
	got := b.String()

	// PascalCase fields should be converted to camelCase.
	if !strings.Contains(got, "token.accessToken") {
		t.Errorf("expected token.accessToken in output, got:\n%s", got)
	}
	if !strings.Contains(got, "token.refreshToken") {
		t.Errorf("expected token.refreshToken in output, got:\n%s", got)
	}
	if !strings.Contains(got, "user.email") {
		t.Errorf("expected user.email in output, got:\n%s", got)
	}

	// PascalCase should NOT appear in output.
	if strings.Contains(got, "AccessToken") {
		t.Errorf("PascalCase AccessToken should not appear in output, got:\n%s", got)
	}
}
