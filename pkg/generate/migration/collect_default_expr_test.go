//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

import (
	"testing"
)

func TestCollectDefaultExpr(t *testing.T) {
	cases := []struct {
		name     string
		toks     []string
		wantExpr string
		wantN    int
	}{
		{"stops at NOT", []string{"0", "NOT", "NULL"}, "0", 1},
		{"function default", []string{"now()", "NOT", "NULL"}, "now()", 1},
		{"multi token expr stops at CHECK", []string{"'a'", "::", "text", "CHECK"}, "'a' :: text", 3},
		{"consumes all when no stop", []string{"42"}, "42", 1},
		{"empty", nil, "", 0},
		{"stops at GENERATED", []string{"x", "GENERATED"}, "x", 1},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			gotExpr, gotN := collectDefaultExpr(c.toks)
			if gotExpr != c.wantExpr || gotN != c.wantN {
				t.Errorf("collectDefaultExpr(%#v) = (%q,%d), want (%q,%d)", c.toks, gotExpr, gotN, c.wantExpr, c.wantN)
			}
		})
	}
}
