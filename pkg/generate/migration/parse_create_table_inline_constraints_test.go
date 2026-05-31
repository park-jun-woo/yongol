//ff:func feature=migration type=test control=iteration dimension=1
//ff:what parse_statements_unit_test — parseCreateTable/parseCreateIndex/parseAlterTable + 하위 parse 헬퍼 통합 단위 테스트
package migration

import (
	"testing"
)

func TestParseCreateTable_InlineConstraints(t *testing.T) {
	s := NewSchema()
	stmt := `CREATE TABLE orders (
		id INTEGER NOT NULL,
		user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
		qty INTEGER CHECK (qty > 0),
		code TEXT UNIQUE,
		PRIMARY KEY (id),
		CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id),
		CHECK (qty <= 100)
	)`
	if err := parseCreateTable(s, stmt); err != nil {
		t.Fatalf("error: %v", err)
	}
	tbl := s.Tables["orders"]
	// inline REFERENCES + named FK -> 2 FKs
	if len(tbl.ForeignKeys) != 2 {
		t.Errorf("got %d FKs, want 2: %+v", len(tbl.ForeignKeys), tbl.ForeignKeys)
	}
	// find the inline one with ON DELETE CASCADE
	var foundCascade bool
	for _, fk := range tbl.ForeignKeys {
		if fk.RefTable == "users" && fk.OnDelete == "CASCADE" {
			foundCascade = true
		}
	}
	if !foundCascade {
		t.Errorf("inline FK ON DELETE CASCADE not parsed: %+v", tbl.ForeignKeys)
	}
	// inline CHECK + table-level CHECK -> 2 checks
	if len(tbl.Checks) != 2 {
		t.Errorf("got %d checks, want 2: %+v", len(tbl.Checks), tbl.Checks)
	}
	// UNIQUE code -> 1 index
	if len(tbl.Indexes) != 1 || !tbl.Indexes[0].Unique {
		t.Errorf("expected 1 unique index for code: %+v", tbl.Indexes)
	}
}
