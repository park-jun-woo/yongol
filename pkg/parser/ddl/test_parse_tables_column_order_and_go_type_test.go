//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — ColumnOrder 선언 순서 보존 + Columns RawType 보존

package ddl

import (
	"os"
	"path/filepath"
	"strings"
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
	// Phase002 — parser preserves RawType verbatim. Phase001 ships GoTypeOf
	// in pkg/generate/gogin/types; the parser-level test only asserts the
	// raw token survives the round-trip.
	wantRaw := map[string]string{
		"id":         "BIGSERIAL",
		"name":       "VARCHAR(100)",
		"amount":     "NUMERIC",
		"active":     "BOOLEAN",
		"payload":    "JSONB",
		"created_at": "TIMESTAMPTZ",
	}
	for col, wt := range wantRaw {
		if got := strings.ToUpper(strings.TrimSpace(tb.Columns[col].RawType)); got != wt {
			t.Errorf("Columns[%s].RawType = %q, want %q", col, got, wt)
		}
	}
}
