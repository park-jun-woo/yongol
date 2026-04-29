//ff:func feature=validate type=test control=sequence topic=ssac-sqlc
//ff:what TestXQS20_Case2 — @post User + RETURNING id, email → diag 1 (Model ↔ partial)

package ssac_sqlc

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXQS20_Case2_PostModelPartialEmits(t *testing.T) {
	fs := &yongol.Fullstack{
		DDLTables: []ddl.Table{userTable()},
		ServiceFuncs: []ssacparser.ServiceFunc{{
			Name: "Register", FileName: "register.ssac",
			Sequences: []ssacparser.Sequence{makeSeq("post", "User", "User.Create")},
		}},
		SQLcQueries: []sqlcparser.QuerySpec{makeQuery(
			"UserCreate", "User", "Create", "one",
			"INSERT INTO users (email) VALUES (@email) RETURNING id, email;",
		)},
	}
	diags := runXQS20(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %v", len(diags), diags)
	}
	if !strings.Contains(diags[0], "[XQS-20]") {
		t.Errorf("diag missing rule tag: %s", diags[0])
	}
	if !strings.Contains(diags[0], `"User"`) || !strings.Contains(diags[0], `"UserCreateRow"`) {
		t.Errorf("diag missing declared/expected pair: %s", diags[0])
	}
	if !strings.Contains(diags[0], "partial Row") {
		t.Errorf("diag missing reason: %s", diags[0])
	}
}
