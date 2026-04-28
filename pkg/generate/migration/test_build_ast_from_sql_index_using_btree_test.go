//ff:func feature=migration type=test control=sequence
//ff:what TestBuildASTFromSQL_IndexUsingBTREE — 명시적 btree 보존

package migration

import "testing"

// TestBuildASTFromSQL_IndexUsingBTREE verifies an explicit btree is
// parsed as "btree" and preserved verbatim by the emitter.
func TestBuildASTFromSQL_IndexUsingBTREE(t *testing.T) {
	sql := `
CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT);
CREATE INDEX idx_t_name ON t USING BTREE (name);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	tbl := s.Tables["t"]
	if len(tbl.Indexes) != 1 {
		t.Fatalf("expected 1 index, got %d", len(tbl.Indexes))
	}
	if got := tbl.Indexes[0].Method; got != "btree" {
		t.Errorf("Method = %q, want %q", got, "btree")
	}
}
