//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestIndexMethod_USING — USING <method> 절이 파싱/emit/diff 전 과정에서 보존되는지 확인 (BUG-032 / Phase009)
package migration

import (
	"strings"
	"testing"
)

// TestBuildASTFromSQL_IndexUsingGIN verifies the parser captures the
// `USING <method>` clause into Index.Method. Regression guard for BUG-032.
func TestBuildASTFromSQL_IndexUsingGIN(t *testing.T) {
	sql := `
CREATE TABLE refresh_tokens (id BIGSERIAL PRIMARY KEY, claims JSONB);
CREATE INDEX refresh_tokens_claims_idx ON refresh_tokens USING GIN (claims);
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	tbl := s.Tables["refresh_tokens"]
	if tbl == nil {
		t.Fatalf("refresh_tokens table not found")
	}
	var gin *Index
	for _, idx := range tbl.Indexes {
		if idx.Name == "refresh_tokens_claims_idx" {
			gin = idx
			break
		}
	}
	if gin == nil {
		t.Fatalf("refresh_tokens_claims_idx not found")
	}
	if gin.Method != "gin" {
		t.Errorf("Method = %q, want %q", gin.Method, "gin")
	}
}

// TestBuildASTFromSQL_IndexUsingBTREE verifies an explicit btree is
// parsed as "btree" (the emitter later normalises it away).
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

// TestCreateIndex_SQL_EmitsUsing verifies emit round-trips the method for
// non-btree indexes and omits it for btree (postgres default).
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
			name:   "btree is omitted",
			idx:    &Index{Name: "i", Columns: []string{"c"}, Method: "btree"},
			notHas: "USING",
		},
		{
			name:   "empty method is omitted",
			idx:    &Index{Name: "i", Columns: []string{"c"}, Method: ""},
			notHas: "USING",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			op := CreateIndex{Table: "t", Index: tc.idx}
			got := op.SQL()
			if tc.substr != "" && !strings.Contains(got, tc.substr) {
				t.Errorf("SQL = %q, want substring %q", got, tc.substr)
			}
			if tc.notHas != "" && strings.Contains(got, tc.notHas) {
				t.Errorf("SQL = %q, unexpected substring %q", got, tc.notHas)
			}
		})
	}
}

// TestDiff_IndexMethodChange_BtreeToGIN verifies that changing an index
// method from btree → gin emits DROP + CREATE.
func TestDiff_IndexMethodChange_BtreeToGIN(t *testing.T) {
	prev := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, c JSONB);
CREATE INDEX idx_t_c ON t (c);`)
	curr := mustAST(t, `CREATE TABLE t (id BIGSERIAL PRIMARY KEY, c JSONB);
CREATE INDEX idx_t_c ON t USING GIN (c);`)
	ops := Diff(prev, curr, nil)
	var sawDrop, sawCreate bool
	for _, op := range ops {
		if _, ok := op.(DropIndex); ok {
			sawDrop = true
		}
		if ci, ok := op.(CreateIndex); ok {
			if ci.Index.Method != "gin" {
				t.Errorf("CreateIndex.Method = %q, want %q", ci.Index.Method, "gin")
			}
			sawCreate = true
		}
	}
	if !sawDrop || !sawCreate {
		t.Errorf("expected DROP + CREATE for method change, got: %+v", ops)
	}
}
