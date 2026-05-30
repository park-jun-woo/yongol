//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderVerifyPasswordOp — VerifyPasswordOp → lookup + bcrypt 검증 블록 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderVerifyPasswordOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderVerifyPasswordOp(&b, nil, "    ", "db")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Full", func(t *testing.T) {
		var b strings.Builder
		op := &ir.VerifyPasswordOp{
			Model:        "User",
			EmailCol:     "Email",
			EmailExpr:    "request.Email",
			PasswordExpr: "request.Password",
			HashCol:      "PasswordHash",
			ResultVar:    "user",
			Message:      "invalid credentials",
		}
		renderVerifyPasswordOp(&b, op, "    ", "db")
		out := b.String()
		for _, want := range []string{
			"result = await db.execute(select(User).where(User.email == body.email))",
			"user = result.scalars().first()",
			"if not user:",
			`raise HTTPException(status_code=401, detail="invalid credentials")`,
			"# TODO: bcrypt.checkpw(body.password, user.password_hash)",
		} {
			if !strings.Contains(out, want) {
				t.Errorf("missing %q in:\n%s", want, out)
			}
		}
	})
}
