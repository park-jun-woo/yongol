//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldSQLcTable — 기존파일 skip / DDL 부재 skip / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestScaffoldSQLcTableDDLNotFound(t *testing.T) {
	dir := t.TempDir()
	queriesDir := filepath.Join(dir, "db", "queries")
	if err := os.MkdirAll(queriesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	// No db/users.sql DDL file → skipped (DDL not found).
	var out bytes.Buffer
	if err := scaffoldSQLcTable(dir, queriesDir, "users", "sys", nil, Config{}, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "DDL not found") {
		t.Errorf("expected DDL-not-found message, got: %q", out.String())
	}
}
