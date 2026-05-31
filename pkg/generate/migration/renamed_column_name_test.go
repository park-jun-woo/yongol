//ff:func feature=migration type=test control=iteration dimension=1
//ff:what helpers_lookup_unit_test — checkMap/columnMap/fkMap/indexMap/rename/setFKAction/newEmptyHints/NewSchema/collectTypeTokens 단위 테스트
package migration

import (
	"testing"
)

func TestRenamedColumnName(t *testing.T) {
	rules := []RenameColumnHint{
		{Table: "users", From: "old", To: "new"},
	}
	cases := []struct {
		name      string
		prev, cur string
		col       string
		want      string
	}{
		{"match on prev table", "users", "members", "old", "new"},
		{"match on new table", "people", "users", "old", "new"},
		{"no col match", "users", "users", "other", "other"},
		{"no table match", "orders", "orders", "old", "old"},
	}
	for _, c := range cases {
		c := c
		t.Run(c.name, func(t *testing.T) {
			if got := renamedColumnName(c.prev, c.cur, c.col, rules); got != c.want {
				t.Errorf("renamedColumnName = %q, want %q", got, c.want)
			}
		})
	}
	if got := renamedColumnName("t", "t", "c", nil); got != "c" {
		t.Errorf("nil rules -> %q, want c", got)
	}
}
