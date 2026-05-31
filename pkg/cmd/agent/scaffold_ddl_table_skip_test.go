//ff:func feature=agent type=test control=sequence
//ff:what TestScaffoldDDLTable — 기존파일 skip / 미지원 backend LLM 에러 분기 검증
package agent

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

func TestScaffoldDDLTableSkip(t *testing.T) {
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
	err := scaffoldDDLTable(dir, dbDir, "users", ff, nil, "sysprompt", Config{}, &out)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !strings.Contains(out.String(), "skipped") {
		t.Errorf("expected skip message, got: %q", out.String())
	}
}
