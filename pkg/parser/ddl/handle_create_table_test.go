//ff:func feature=manifest type=test control=sequence
//ff:what handleCreateTable — table 등록 + pendingArchived/pendingFuncManaged 반영
package ddl

import (
	"testing"
)

func TestHandleCreateTable(t *testing.T) {
	t.Run("archived and func-managed flags", func(t *testing.T) {
		tables := map[string]*Table{}
		name := handleCreateTable("CREATE TABLE users (", tables, true, true, "/x.sql", 1)
		if name != "users" {
			t.Fatalf("name = %q, want users", name)
		}
		tb := tables["users"]
		if tb == nil {
			t.Fatal("table not registered")
		}
		if !tb.Archived || !tb.FuncManaged {
			t.Errorf("flags = archived:%v func:%v, want both true", tb.Archived, tb.FuncManaged)
		}
	})
	t.Run("no flags leaves false", func(t *testing.T) {
		tables := map[string]*Table{}
		handleCreateTable("CREATE TABLE plain (", tables, false, false, "/x.sql", 1)
		tb := tables["plain"]
		if tb.Archived || tb.FuncManaged {
			t.Errorf("flags = archived:%v func:%v, want both false", tb.Archived, tb.FuncManaged)
		}
	})
}
