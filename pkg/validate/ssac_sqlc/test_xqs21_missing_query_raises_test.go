//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXqs21_MissingQuery_Raises — @verify-password 쿼리 부재 → [XQS-21]

package ssac_sqlc

import (
	"strings"
	"testing"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXqs21_MissingQuery_Raises(t *testing.T) {
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
	}
	diags := xqs21VerifyPasswordQuery(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "XQS-21") {
		t.Errorf("diag missing XQS-21: %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "UserFindByEmail") {
		t.Errorf("diag missing UserFindByEmail: %s", diags[0].Message)
	}
}
