//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs21_QueryExists_NoDiag — @verify-password 쿼리 존재 → PASS

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs21_QueryExists_NoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name:     "Login",
			FileName: "auth.ssac",
			Sequences: []ssacparser.Sequence{{
				Type:     "verify-password",
				Model:    "User",
				EmailCol: "email",
			}},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{{
			Name:  "UserFindByEmail",
			Model: "User",
		}},
	}
	diags := xqs21VerifyPasswordQuery(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
