//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestWriteHTTPHandlerPreAuth — @verify-password 있는 login 핸들러에 @Req/req.user 제외 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHTTPHandlerPreAuth(t *testing.T) {
	t.Run("LoginSkipsReqUser", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "Login",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/auth/login",
			Feature:     "auth",
			BodyFields: []ir.BodyFieldMeta{
				{Name: "email"},
				{Name: "password"},
			},
			Ops: []ir.Op{
				{Kind: ir.OpVerifyPassword, VerifyPW: &ir.VerifyPasswordOp{
					Model:        "User",
					EmailCol:     "email",
					EmailExpr:    "request.Email",
					HashCol:      "password_hash",
					PasswordExpr: "request.Password",
					ResultVar:    "user",
				}},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if strings.Contains(got, "@Req()") {
			t.Errorf("login handler should NOT have @Req() decorator, got:\n%s", got)
		}
		if strings.Contains(got, "req.user") {
			t.Errorf("login handler should NOT pass req.user, got:\n%s", got)
		}
		if !strings.Contains(got, "@Body()") {
			t.Errorf("login handler should have @Body() decorator, got:\n%s", got)
		}
	})

	t.Run("NormalEndpointKeepsReqUser", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "GetProfile",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/auth/profile",
			Feature:     "auth",
			Ops: []ir.Op{
				{Kind: ir.OpGet, Get: &ir.GetOp{Model: "User"}},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()
		if !strings.Contains(got, "@Req()") {
			t.Errorf("normal handler should have @Req(), got:\n%s", got)
		}
		if !strings.Contains(got, "req.user") {
			t.Errorf("normal handler should pass req.user, got:\n%s", got)
		}
	})
}
