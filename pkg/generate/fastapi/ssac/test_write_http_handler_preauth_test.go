//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHTTPHandlerPreAuth — @verify-password 있는 login 핸들러에 current_user 의존성 제외 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHTTPHandlerPreAuth(t *testing.T) {
	t.Run("LoginSkipsCurrentUser", func(t *testing.T) {
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
		if strings.Contains(got, "Depends(get_current_user)") {
			t.Errorf("login handler should NOT have Depends(get_current_user), got:\n%s", got)
		}
		if !strings.Contains(got, "Depends(get_session)") {
			t.Errorf("login handler should still have Depends(get_session), got:\n%s", got)
		}
		if strings.Contains(got, "current_user") {
			t.Errorf("login handler should not pass current_user to service, got:\n%s", got)
		}
		if !strings.Contains(got, "body: LoginRequest") {
			t.Errorf("login handler should have body parameter, got:\n%s", got)
		}
	})

	t.Run("NormalEndpointKeepsCurrentUser", func(t *testing.T) {
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
		if !strings.Contains(got, "Depends(get_current_user)") {
			t.Errorf("normal handler should have Depends(get_current_user), got:\n%s", got)
		}
		if !strings.Contains(got, "current_user") {
			t.Errorf("normal handler should pass current_user, got:\n%s", got)
		}
	})
}
