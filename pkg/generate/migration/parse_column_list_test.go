//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestParseColumnList(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{"parenthesised", "(a, b, c)", []string{"a", "b", "c"}},
		{"bare", "a, b", []string{"a", "b"}},
		{"uppercase lowered", "(ID, Name)", []string{"id", "name"}},
		{"quoted preserved", `("Order", id)`, []string{"Order", "id"}},
		{"empty entries skipped", "(a, , b)", []string{"a", "b"}},
		{"single", "id", []string{"id"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := parseColumnList(c.in); !reflect.DeepEqual(got, c.want) {
				t.Errorf("parseColumnList(%q) = %#v, want %#v", c.in, got, c.want)
			}
		})
	}
}
