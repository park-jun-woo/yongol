//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case4 — @post UserNewRow + RETURNING * → diag 1 (Row ↔ full)

package ssac_sqlc

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case4_PostRowFullEmits(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{userTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Register", FileName: "register.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("post", "UserNewRow", "User.New")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"UserNew", "User", "New", "one",
			"INSERT INTO users (email) VALUES (@email) RETURNING *;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0], `"UserNewRow"`) || !strings.Contains(diags[0], `"User"`) {
		t.Errorf("diag missing declared/expected pair: %s", diags[0])
	}
	if !strings.Contains(diags[0], "RETURNING * → model") {
		t.Errorf("diag missing reason: %s", diags[0])
	}
}
