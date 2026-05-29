//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — 다중 토큰 PG 타입 8 종 RawType 보존 회귀

package ddl

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestParseTables_MultiTokenPgTypes(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE multi_tokens (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    score DOUBLE PRECISION NOT NULL,
    occurred_at TIMESTAMP WITH TIME ZONE NOT NULL,
    occurred_naive TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    daily TIME WITH TIME ZONE NOT NULL,
    daily_naive TIME WITHOUT TIME ZONE NOT NULL,
    name CHARACTER VARYING(255) NOT NULL,
    fixed CHARACTER(10) NOT NULL,
    flags BIT VARYING(8) NOT NULL
);`
	if err := os.WriteFile(filepath.Join(dir, "multi.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatalf("tables len = %d, want 1", len(tables))
	}
	tb := tables[0]
	wantRaw := map[string]string{
		"score":          "DOUBLE PRECISION",
		"occurred_at":    "TIMESTAMP WITH TIME ZONE",
		"occurred_naive": "TIMESTAMP WITHOUT TIME ZONE",
		"daily":          "TIME WITH TIME ZONE",
		"daily_naive":    "TIME WITHOUT TIME ZONE",
		"name":           "CHARACTER VARYING(255)",
		"fixed":          "CHARACTER(10)",
		"flags":          "BIT VARYING(8)",
	}
	for col, wt := range wantRaw {
		got := strings.ToUpper(strings.TrimSpace(tb.Columns[col].RawType))
		if got != wt {
			t.Errorf("Columns[%s].RawType = %q, want %q", col, got, wt)
		}
	}
}
