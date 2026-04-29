//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case3 — @post UserNewRow + RETURNING id → diag 0 (Row ↔ partial)

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case3_PostRowPartialPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{userTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Register", FileName: "register.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("post", "UserNewRow", "User.New")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"UserNew", "User", "New", "one",
			"INSERT INTO users (email) VALUES (@email) RETURNING id;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
