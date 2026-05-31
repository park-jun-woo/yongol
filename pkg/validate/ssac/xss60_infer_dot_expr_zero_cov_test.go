//ff:func feature=validate type=test control=sequence
//ff:what TestXss60ByName_ZeroCov — XSS-60 publish/subscribe 타입 수집·비교 헬퍼 직접 호출
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestXss60InferDotExpr_ZeroCov(t *testing.T) {
	fn := parsessac.ServiceFunc{
		Name: "Pub",
		Sequences: []parsessac.Sequence{
			{Type: "get", Result: &parsessac.Result{Var: "user", Type: "User"}},
		},
	}
	tables := map[string]*ddl.Table{
		"users": {Name: "users", Columns: map[string]ddl.Column{
			"email": {Name: "email", RawType: "TEXT"},
		}},
	}
	// var.Field resolved through model→table→column.
	got := xss60InferDotExpr("user.Email", 4, fn, tables)
	if got == "" {
		t.Errorf("expected Go type for user.Email, got empty")
	}
	// unresolvable variable model → "".
	if v := xss60InferDotExpr("ghost.X", 5, fn, tables); v != "" {
		t.Errorf("unknown var should yield empty, got %q", v)
	}
}
