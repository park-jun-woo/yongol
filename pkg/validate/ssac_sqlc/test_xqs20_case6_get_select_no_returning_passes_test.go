//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case6 — @get User + plain SELECT (no RETURNING) → diag 0 (skip)

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case6_GetSelectNoReturningPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{userTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "GetUser", FileName: "get.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("get", "User", "User.FindByID")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"UserFindByID", "User", "FindByID", "one",
			"SELECT * FROM users WHERE id = @id;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
