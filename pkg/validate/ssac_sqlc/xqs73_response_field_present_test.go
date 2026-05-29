//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS73_ResponseFieldPresent — 참조 필드가 SELECT에 존재하면 에러 미발생 확인

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXQS73_ResponseFieldPresent verifies that no error is raised when
// all referenced fields exist in the query's SELECT column list.
func TestXQS73_ResponseFieldPresent(t *testing.T) {
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
			Body:        "SELECT email, name FROM users WHERE email = @email",
			SelectStar:  false,
			SelectCols:  []string{"email", "name"},
			File:        "users.sql",
			Line:        1,
		}},
	}
	diags := xqs73PartialSelectField(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
