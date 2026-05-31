//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldSQLc — 테이블 없음 nil / 기존파일 skip / 미지원 backend LLM 에러 / mkdir 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSQLcLLMError(t *testing.T) {
	// DDL present but queries missing → scaffoldSQLcTable's LLM call fails.
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dbDir, "users.sql"), []byte("CREATE TABLE users();\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	if err := scaffoldSQLc(dir, ff, nil, cfg, &out); err == nil {
		t.Fatal("expected LLM error from scaffoldSQLcTable")
	}
}
