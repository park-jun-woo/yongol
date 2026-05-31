//ff:func feature=migration type=test control=iteration dimension=1
//ff:what tokenize_split_unit_test — tokenizeColumnDef/splitStatements/splitTopLevel/parseColumnList/collectDefaultExpr 단위 테스트
package migration

import (
	"testing"
)

func TestSplitStatements(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want int // number of non-empty trimmed statements expected
	}{
		{"two statements", "CREATE TABLE a (id int); CREATE TABLE b (id int);", 2},
		{"semicolon in string is ignored", "INSERT INTO t VALUES ('a;b');", 1},
		{"semicolon in parens at top", "CREATE TABLE a (id int, c CHECK (x > 0));", 1},
		{"line comment stripped", "-- comment with ; inside\nSELECT 1;", 1},
		{"block comment with semicolon", "/* a; b */ SELECT 1;", 1},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertSplitStatementsCount(t, c.in, c.want)
		})
	}
}
