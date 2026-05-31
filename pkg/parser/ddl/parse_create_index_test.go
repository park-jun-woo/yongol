//ff:func feature=manifest type=test control=sequence
//ff:what parseCreateIndex — CREATE INDEX 파싱 (UNIQUE/USING method/no ON)
package ddl

import (
	"testing"
)

func TestParseCreateIndex(t *testing.T) {
	t.Run("basic index", func(t *testing.T) {
		tables := map[string]*Table{"users": {Name: "users"}}
		parseCreateIndex("CREATE INDEX idx_email ON users (email)", tables)
		idx := tables["users"].Indexes
		if len(idx) != 1 || idx[0].Name != "idx_email" || idx[0].IsUnique {
			t.Fatalf("indexes = %+v", idx)
		}
		if len(idx[0].Columns) != 1 || idx[0].Columns[0] != "email" {
			t.Errorf("columns = %v", idx[0].Columns)
		}
	})
	t.Run("unique with using method", func(t *testing.T) {
		tables := map[string]*Table{"docs": {Name: "docs"}}
		parseCreateIndex("CREATE UNIQUE INDEX idx_gin ON docs USING gin (body)", tables)
		idx := tables["docs"].Indexes[0]
		if !idx.IsUnique || idx.Method != "gin" {
			t.Errorf("index = %+v, want unique gin", idx)
		}
	})
	t.Run("no ON clause is no-op", func(t *testing.T) {
		tables := map[string]*Table{"t": {Name: "t"}}
		parseCreateIndex("CREATE INDEX broken", tables)
		if len(tables["t"].Indexes) != 0 {
			t.Errorf("expected no indexes")
		}
	})
	t.Run("no paren after ON is no-op", func(t *testing.T) {
		tables := map[string]*Table{"t": {Name: "t"}}
		parseCreateIndex("CREATE INDEX i ON t", tables)
		if len(tables["t"].Indexes) != 0 {
			t.Errorf("expected no indexes (no column list)")
		}
	})
	t.Run("unknown table is no-op", func(t *testing.T) {
		tables := map[string]*Table{}
		parseCreateIndex("CREATE INDEX i ON ghost (x)", tables)
		if len(tables) != 0 {
			t.Errorf("no table should be created")
		}
	})
}
