//ff:func feature=agent type=test control=sequence
//ff:what TestScaffold — 테이블 없음 skip / 테이블 존재+미지원 backend → DDL phase 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldSQLcError(t *testing.T) {
	// Pre-create the DDL file so the DDL phase skips, then the sqlc phase's LLM
	// call fails for the unsupported backend, yielding a "scaffold sqlc" error.
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
	err := scaffold(dir, ff, nil, cfg, &out)
	if err == nil || !strings.Contains(err.Error(), "scaffold sqlc") {
		t.Fatalf("expected scaffold sqlc error, got: %v", err)
	}
}
