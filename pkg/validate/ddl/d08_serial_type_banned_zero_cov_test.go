//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what validate/ddl 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestD08SerialTypeBanned_ZeroCov(t *testing.T) {
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	sql := "CREATE TABLE t (\n  id BIGSERIAL PRIMARY KEY\n);\n"
	if err := os.WriteFile(filepath.Join(dbDir, "t.sql"), []byte(sql), 0o644); err != nil {
		t.Fatal(err)
	}
	fs := &yongol.Fullstack{SpecsDir: tmp}
	diags := d08SerialTypeBanned(fs)
	if len(diags) == 0 {
		t.Errorf("expected D-8 diag for BIGSERIAL column")
	}
	// missing dir → nil
	if got := d08SerialTypeBanned(&yongol.Fullstack{SpecsDir: t.TempDir()}); got != nil {
		t.Errorf("missing db dir should give nil")
	}
}
