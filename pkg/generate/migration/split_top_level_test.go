//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

import (
	"reflect"
	"testing"
)

func TestSplitTopLevel(t *testing.T) {
	cases := []struct {
		name string
		in   string
		sep  byte
		want []string
	}{
		{"simple comma", "a, b, c", ',', []string{"a", " b", " c"}},
		{"comma inside parens ignored", "a, b(1, 2), c", ',', []string{"a", " b(1, 2)", " c"}},
		{"comma inside single quote ignored", "a, 'x,y', c", ',', []string{"a", " 'x,y'", " c"}},
		{"comma inside double quote ignored", `a, "x,y", c`, ',', []string{"a", ` "x,y"`, " c"}},
		{"no sep", "abc", ',', []string{"abc"}},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := splitTopLevel(c.in, c.sep); !reflect.DeepEqual(got, c.want) {
				t.Errorf("splitTopLevel(%q,%q) = %#v, want %#v", c.in, c.sep, got, c.want)
			}
		})
	}
}
