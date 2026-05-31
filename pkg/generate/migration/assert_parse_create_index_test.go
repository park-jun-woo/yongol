//ff:func feature=migration type=test-helper control=sequence
//ff:what assertParseCreateIndex — parseCreateIndex 후 생성된 인덱스 속성/컬럼 검증 헬퍼
package migration

import "testing"

// assertParseCreateIndex parses stmt into s and asserts the resulting index's
// unique/method/where flags and column count.
func assertParseCreateIndex(t *testing.T, s *Schema, stmt, idxName string, unique bool, method, where string, cols []string) {
	t.Helper()
	if err := parseCreateIndex(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	idx := findIndexByName(s.Tables["users"].Indexes, idxName)
	if idx == nil {
		t.Fatalf("index %s not found", idxName)
	}
	if idx.Unique != unique || idx.Method != method || idx.Where != where {
		t.Errorf("idx %+v mismatch (want unique=%v method=%q where=%q)", idx, unique, method, where)
	}
	if len(idx.Columns) != len(cols) {
		t.Errorf("cols %v, want %v", idx.Columns, cols)
	}
}
