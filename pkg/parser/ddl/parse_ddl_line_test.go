//ff:func feature=manifest type=test control=sequence
//ff:what parseDDLLine — CREATE TABLE/INDEX 분기, 닫는 괄호 리셋, 컬럼 위임
package ddl

import (
	"testing"
)

func TestParseDDLLine(t *testing.T) {
	t.Run("create table sets current table", func(t *testing.T) {
		tables := map[string]*Table{}
		cur := parseDDLLine("CREATE TABLE users (", "", tables, false, false, "/x.sql", 1)
		if cur != "users" || tables["users"] == nil {
			t.Errorf("cur = %q, tables = %v", cur, tables)
		}
	})
	t.Run("closing paren resets current table", func(t *testing.T) {
		tables := map[string]*Table{"users": newTable()}
		cur := parseDDLLine(");", "users", tables, false, false, "/x.sql", 5)
		if cur != "" {
			t.Errorf("cur = %q, want empty", cur)
		}
	})
	t.Run("create index does not change current table", func(t *testing.T) {
		tables := map[string]*Table{"users": {Name: "users"}}
		cur := parseDDLLine("CREATE INDEX i ON users (email)", "users", tables, false, false, "/x.sql", 2)
		if cur != "users" || len(tables["users"].Indexes) != 1 {
			t.Errorf("cur = %q, indexes = %+v", cur, tables["users"].Indexes)
		}
	})
	t.Run("column line outside table is ignored", func(t *testing.T) {
		tables := map[string]*Table{}
		cur := parseDDLLine("id BIGINT", "", tables, false, false, "/x.sql", 3)
		if cur != "" || len(tables) != 0 {
			t.Errorf("expected no-op, cur=%q tables=%v", cur, tables)
		}
	})
	t.Run("column line inside table parsed", func(t *testing.T) {
		tables := map[string]*Table{"users": newTable()}
		parseDDLLine("id BIGINT NOT NULL", "users", tables, false, false, "/x.sql", 4)
		if _, ok := tables["users"].Columns["id"]; !ok {
			t.Errorf("expected column id")
		}
	})
}
