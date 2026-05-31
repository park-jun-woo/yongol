//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestZeroCov — 0% render/util 함수 (controllerRoutePrefix / formatCallTarget / render*Op / resolveDataKey 등) 회귀
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderVerifyPasswordOp(t *testing.T) {
	var b strings.Builder
	renderVerifyPasswordOp(&b, nil, "  ", "this.prisma")
	if b.String() != "" {
		t.Errorf("nil verify")
	}
	b.Reset()
	renderVerifyPasswordOp(&b, &ir.VerifyPasswordOp{
		Model:        "User",
		EmailCol:     "Email",
		EmailExpr:    "request.email",
		HashCol:      "PasswordHash",
		PasswordExpr: "request.password",
		ResultVar:    "user",
		Message:      "invalid",
	}, "  ", "this.prisma")
	out := b.String()
	if !strings.Contains(out, "findUnique({ where: { email: body.email } })") {
		t.Errorf("verify lookup = %q", out)
	}
	if !strings.Contains(out, "bcrypt.compare(body.password") {
		t.Errorf("verify bcrypt = %q", out)
	}
}
