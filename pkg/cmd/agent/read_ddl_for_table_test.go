//ff:func feature=agent type=test control=sequence
//ff:what TestReadDDLForTable — 테이블명으로 db/<table>.sql 읽기 및 부재/빈 입력 처리 검증

package agent

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadDDLForTable(t *testing.T) {
	dir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(dir, "db"), 0o755); err != nil {
		t.Fatal(err)
	}
	ddl := "CREATE TABLE users (id UUID);"
	if err := os.WriteFile(filepath.Join(dir, "db", "users.sql"), []byte(ddl), 0o644); err != nil {
		t.Fatal(err)
	}

	if got := readDDLForTable(dir, "users"); got != ddl {
		t.Errorf("readDDLForTable = %q, want %q", got, ddl)
	}
	if got := readDDLForTable(dir, ""); got != "" {
		t.Errorf("empty table = %q, want empty", got)
	}
	if got := readDDLForTable(dir, "missing"); got != "" {
		t.Errorf("missing file = %q, want empty", got)
	}
}
