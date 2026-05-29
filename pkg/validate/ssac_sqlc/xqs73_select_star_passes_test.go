//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS73_SelectStarPasses — SELECT * 쿼리는 항상 통과 확인

package ssac_sqlc

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestXQS73_SelectStarPasses verifies that SELECT * queries are always allowed.
func TestXQS73_SelectStarPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "GetUser", FileName: "getuser.ssac",
			Sequences: []ssacparser.Sequence{
				{
					Type:   "get",
					Model:  "User.FindByID",
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
			Name:        "UserFindByID",
			Model:       "User",
			Method:      "FindByID",
			Cardinality: "one",
			RowType:     "UserFindByIDRow",
			Body:        "SELECT * FROM users WHERE id = @id",
			SelectStar:  true,
			SelectCols:  nil,
			File:        "users.sql",
			Line:        1,
		}},
	}
	diags := xqs73PartialSelectField(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
