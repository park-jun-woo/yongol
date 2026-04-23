//ff:func feature=migration type=test control=iteration dimension=1
//ff:what TestBuildASTFromSQL_CreateIndex — plain INDEX + UNIQUE INDEX + WHERE partial 인덱스 파싱
package migration

import (
	"strings"
	"testing"
)

func TestBuildASTFromSQL_CreateIndex(t *testing.T) {
	sql := `
CREATE TABLE t (id BIGSERIAL PRIMARY KEY, name TEXT NOT NULL);
CREATE INDEX idx_t_name ON t (name);
CREATE UNIQUE INDEX uq_t_name ON t (name) WHERE name <> '';
`
	s := NewSchema()
	if err := BuildASTFromSQL(s, sql); err != nil {
		t.Fatal(err)
	}
	tbl := s.Tables["t"]
	if len(tbl.Indexes) != 2 {
		t.Fatalf("expected 2 indexes, got %d: %+v", len(tbl.Indexes), tbl.Indexes)
	}
	var plain, uq *Index
	for _, idx := range tbl.Indexes {
		if idx.Name == "idx_t_name" {
			plain = idx
		}
		if idx.Name == "uq_t_name" {
			uq = idx
		}
	}
	if plain == nil || plain.Unique {
		t.Errorf("plain idx: %+v", plain)
	}
	if uq == nil || !uq.Unique || !strings.Contains(uq.Where, "<>") {
		t.Errorf("unique idx: %+v", uq)
	}
}
