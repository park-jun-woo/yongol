//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case1 — @post User + RETURNING * → diag 0 (full RETURNING ↔ Model)

package ssac_sqlc

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case1_PostModelFullStarPasses(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{userTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Register", FileName: "register.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("post", "User", "User.Create")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"UserCreate", "User", "Create", "one",
			"INSERT INTO users (email) VALUES (@email) RETURNING *;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags, got %d: %v", len(diags), diags)
	}
}
