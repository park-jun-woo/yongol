//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldSQLcTable — 기존파일 skip / DDL 부재 skip / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestScaffoldSQLcTableLLMError(t *testing.T) {
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	queriesDir := filepath.Join(dbDir, "queries")
	if err := os.MkdirAll(queriesDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte("CREATE TABLE users();\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffoldSQLcTable(dir, queriesDir, "users", "sys", nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error")
	}
}
