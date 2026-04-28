//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestCreateIndex_SQL_EmitsUsing — emit 이 USING 절을 보존/생략하는 조건 확인

package migration

import "testing"

// TestCreateIndex_SQL_EmitsUsing verifies emit preserves the user-declared
// method verbatim (including explicit btree) and omits USING only when no
// method was declared in the DDL.
func TestCreateIndex_SQL_EmitsUsing(t *testing.T) {
	tests := []struct {
		name   string
		idx    *Index
		substr string
		notHas string
	}{
		{
			name:   "gin emits USING gin",
			idx:    &Index{Name: "i", Columns: []string{"c"}, Method: "gin"},
			substr: "USING gin",
		},
		{
			name:   "hash emits USING hash",
			idx:    &Index{Name: "i", Columns: []string{"c"}, Method: "hash"},
			substr: "USING hash",
		},
		{
			name:   "explicit btree is preserved",
			idx:    &Index{Name: "i", Columns: []string{"c"}, Method: "btree"},
			substr: "USING btree",
		},
		{
			name:   "empty method is omitted",
			idx:    &Index{Name: "i", Columns: []string{"c"}, Method: ""},
			notHas: "USING",
		},
	}
	for _, tc := range tests {
		runCreateIndexSQLCase(t, tc.name, tc.idx, tc.substr, tc.notHas)
	}
}
