//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS73_ResponseFieldMissing — 부분 SELECT에서 참조 필드 누락 시 에러 발생 검증

package ssac_sqlc

import (
	"strings"
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXQS73_ResponseFieldMissing verifies that @response { field: var.Field }
// raises an error when the field is not in the query's SELECT column list.
func TestXQS73_ResponseFieldMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "GetUser", FileName: "getuser.ssac",
			Sequences: []ssacparser.Sequence{
				{
					Type:   "get",
					Model:  "User.FindByEmail",
					Result: &ssacparser.Result{Type: "User", Var: "user"},
					Line:   5,
				},
				{
					Type:   "response",
					Fields: map[string]string{"email": "user.Email", "name": "user.Name"},
					Line:   10,
				},
			},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{{
			Name:        "UserFindByEmail",
			Model:       "User",
			Method:      "FindByEmail",
			Cardinality: "one",
			RowType:     "UserFindByEmailRow",
			Body:        "SELECT email FROM users WHERE email = @email",
			SelectStar:  false,
			SelectCols:  []string{"email"},
			File:        "users.sql",
			Line:        1,
		}},
	}
	diags := xqs73PartialSelectField(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "Name") {
		t.Fatalf("expected error about 'Name' field, got: %s", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "XQS-73") {
		t.Fatalf("expected XQS-73 rule ID, got: %s", diags[0].Message)
	}
}
