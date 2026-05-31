//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what validate/ddl 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD02NullableColumn_ZeroCov(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	sql := "CREATE TABLE users (\n  id BIGINT PRIMARY KEY,\n  email VARCHAR(255)\n);\n"
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d02NullableColumn(fs)
	if len(diags) == 0 {
		t.Errorf("expected D-2 diag for nullable email column")
	}
	// missing db dir → nil
	if got := d02NullableColumn(&yongol.Fullstack{SpecsDir: t.TempDir()}); got != nil {
		t.Errorf("missing db dir should give nil")
	}
}
