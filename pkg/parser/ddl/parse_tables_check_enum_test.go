//ff:func feature=manifest type=test control=iteration dimension=1
//ff:what ParseTables — CHECK (x IN (...)) enum 값 수집

package ddl

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseTables_CheckEnum(t *testing.T) {
	dir := t.TempDir()
	sql := `CREATE TABLE orders (
    id BIGSERIAL PRIMARY KEY,
    status VARCHAR(32) NOT NULL CHECK (status IN ('pending','paid','cancelled'))
);`
	if err := os.WriteFile(filepath.Join(dir, "orders.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	tables, diags := ParseTables(dir)
	if len(diags) > 0 {
		t.Fatalf("unexpected diags: %v", diags)
	}
	if len(tables) != 1 {
		t.Fatalf("tables count = %d", len(tables))
	}
	vals := tables[0].Columns["status"].CheckEnum
	if len(vals) != 3 {
		t.Fatalf("Columns[status].CheckEnum = %v, want 3 values", vals)
	}
	want := map[string]bool{"pending": true, "paid": true, "cancelled": true}
	for _, v := range vals {
		if !want[v] {
			t.Errorf("unexpected enum val %q", v)
		}
	}
}
