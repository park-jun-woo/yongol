//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case7 — @get UserSummaryRow + partial SELECT (no RETURNING) → diag 0 (skip)

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Rationale: SELECT bodies without RETURNING are out of scope for XQS-20.
// XDS-12 already covers result-type ↔ DDL coverage on the SELECT path.
func TestXQS20_Case7_GetRowPartialSelectPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{userTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "GetUserSummary", FileName: "get.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("get", "UserGetSummaryRow", "User.GetSummary")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"UserGetSummary", "User", "GetSummary", "one",
			"SELECT id, email FROM users WHERE id = @id;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
