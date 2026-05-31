//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestTokenizeColumnDef(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{"simple", "id INTEGER NOT NULL", []string{"id", "INTEGER", "NOT", "NULL"}},
		{"paren preserved", "amount NUMERIC(10, 2) NOT NULL", []string{"amount", "NUMERIC(10, 2)", "NOT", "NULL"}},
		{"single quoted default", "status TEXT DEFAULT 'active'", []string{"status", "TEXT", "DEFAULT", "'active'"}},
		{"escaped single quote", "msg TEXT DEFAULT 'it''s'", []string{"msg", "TEXT", "DEFAULT", "'it''s'"}},
		{"double quoted ident", `"Order" INTEGER`, []string{`"Order"`, "INTEGER"}},
		{"nested parens", "v INTEGER CHECK (v IN (1, 2, 3))", []string{"v", "INTEGER", "CHECK", "(v IN (1, 2, 3))"}},
		{"empty", "", nil},
		{"leading whitespace collapses", "  a  b  ", []string{"a", "b"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := tokenizeColumnDef(c.in); !reflect.DeepEqual(got, c.want) {
				t.Errorf("tokenizeColumnDef(%q) = %#v, want %#v", c.in, got, c.want)
			}
		})
	}
}
