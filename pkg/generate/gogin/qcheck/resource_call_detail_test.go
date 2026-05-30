//ff:func feature=gen-gogin type=test control=branch topic=defensive
//ff:what TestResourceCallDetail — os.Open / Query* / Prepare* 매칭 + 비매칭/비-selector 분기 검증

package qcheck

import "testing"

func TestResourceCallDetail(t *testing.T) {
	cases := []struct {
		expr string
		want string
	}{
		{"os.Open(p)", "os.Open"},
		{"db.Query(q)", "db.Query"},
		{"conn.QueryContext(ctx, q)", "conn.QueryContext"},
		{"tx.Prepare(q)", "tx.Prepare"},
		{"qtx.PrepareContext(ctx, q)", "qtx.PrepareContext"},
		{"os.Create(p)", ""},     // os but not Open
		{"db.Exec(q)", ""},       // not a resource-returning method
		{"plain(x)", ""},         // not a selector
		{"a.b.Query(q)", ""},     // receiver not a plain ident
	}
	for _, c := range cases {
		if got := resourceCallDetail(callExpr(t, c.expr)); got != c.want {
			t.Errorf("resourceCallDetail(%q) = %q, want %q", c.expr, got, c.want)
		}
	}
}
