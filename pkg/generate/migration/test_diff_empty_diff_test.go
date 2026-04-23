//ff:func feature=migration type=test control=sequence
//ff:what TestDiff_EmptyDiff — prev == curr 이면 ops 없음
package migration

import "testing"

func TestDiff_EmptyDiff(t *testing.T) {
	sql := `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL);`
	ops := Diff(mustAST(t, sql), mustAST(t, sql), nil)
	if len(ops) != 0 {
		t.Errorf("expected no ops, got %d: %+v", len(ops), ops)
	}
}
