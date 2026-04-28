//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — ColumnOrder 선언 순서 보존 + Columns Go 타입 매핑

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_ColumnOrderAndGoType(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE mixed (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    amount NUMERIC NOT NULL,
    active BOOLEAN NOT NULL DEFAULT 'false',
    payload JSONB,
    created_at TIMESTAMPTZ NOT NULL
);`
	if err := os.WriteFile(filepath.Join(dir, "mixed.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	tb := tables[0]
	wantOrder := []string{"id", "name", "amount", "active", "payload", "created_at"}
	if len(tb.ColumnOrder) != len(wantOrder) {
		t.Fatalf("ColumnOrder len = %d, want %d: %v", len(tb.ColumnOrder), len(wantOrder), tb.ColumnOrder)
	}
	for i, c := range wantOrder {
		if tb.ColumnOrder[i] != c {
			t.Errorf("ColumnOrder[%d] = %q, want %q", i, tb.ColumnOrder[i], c)
		}
	}
	wantTypes := map[string]string{
		"id":         "int64",
		"name":       "string",
		"amount":     "float64",
		"active":     "bool",
		"payload":    "json.RawMessage",
		"created_at": "time.Time",
	}
	for col, wt := range wantTypes {
		if got := GoTypeOf(tb.Columns[col]); got != wt {
			t.Errorf("GoTypeOf(Columns[%s]) = %q, want %q", col, got, wt)
		}
	}
}
