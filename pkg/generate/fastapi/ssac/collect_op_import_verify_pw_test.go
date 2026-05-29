//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectOpImportVerifyPW — OpVerifyPassword → UsesSelect + Model 등록 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectOpImportVerifyPW(t *testing.T) {
	t.Run("VerifyPasswordAddsSelectAndModel", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpVerifyPassword,
			VerifyPW: &ir.VerifyPasswordOp{
				Model:        "User",
				EmailCol:     "email",
				EmailExpr:    "request.Email",
				HashCol:      "password_hash",
				PasswordExpr: "request.Password",
				ResultVar:    "user",
			},
		}
		collectOpImport(&d, op, "auth")
		if !d.UsesSelect {
			t.Error("expected UsesSelect true for @verify-password")
		}
		if !d.Models["User"] {
			t.Errorf("expected User model in imports, got %v", d.Models)
		}
	})

	t.Run("VerifyPasswordNilOp", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpVerifyPassword,
		}
		collectOpImport(&d, op, "auth")
		if !d.UsesSelect {
			t.Error("expected UsesSelect true even with nil VerifyPW")
		}
		if len(d.Models) != 0 {
			t.Errorf("expected no models with nil VerifyPW, got %v", d.Models)
		}
	})

	t.Run("VerifyPasswordSnakeCaseModel", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpVerifyPassword,
			VerifyPW: &ir.VerifyPasswordOp{
				Model:    "admin_user",
				EmailCol: "email",
			},
		}
		collectOpImport(&d, op, "auth")
		if !d.Models["AdminUser"] {
			t.Errorf("expected PascalCase model AdminUser, got %v", d.Models)
		}
	})
}
