//ff:func feature=migration type=test control=iteration dimension=1
//ff:what diff_helpers_unit_test — checkDiffForOne/findPrevViaRenameHint/checkDropOps/checkAlterOrAddOps/allCreateTableOps/createTableWithDeps/parseTableFKRef 단위 테스트
package migration

import (
	"testing"
)

func TestParseTableFKRef(t *testing.T) {
	cases := []struct {
		name         string
		toks         []string
		wantTable    string
		wantCols     []string
		wantConsumed int
	}{
		{"target with parens", []string{"(a)", "REFERENCES", "users(id)"}, "users", []string{"id"}, 3},
		{"target then separate parens", []string{"(a)", "REFERENCES", "users", "(id, tenant)"}, "users", []string{"id", "tenant"}, 4},
		{"target without cols", []string{"(a)", "REFERENCES", "users"}, "users", nil, 3},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertParseTableFKRef(t, c.toks, c.wantTable, c.wantCols, c.wantConsumed)
		})
	}
}
