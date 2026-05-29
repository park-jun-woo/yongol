//ff:func feature=manifest type=test control=sequence
//ff:what extractTableName — CREATE TABLE 라인 파싱하며 tables 맵에 File/Line 기록

package ddl

import "testing"

func TestExtractTableName_WritesTable(t *testing.T) {
	tables := map[string]*Table{}
	name := extractTableName("CREATE TABLE users (", tables, "x.sql", 5)
	if name != "users" {
		t.Errorf("name = %q", name)
	}
	tb, ok := tables["users"]
	if !ok {
		t.Fatalf("users not registered")
	}
	if tb.File != "x.sql" || tb.Line != 5 {
		t.Errorf("File/Line = %q/%d", tb.File, tb.Line)
	}
}
