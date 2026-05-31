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

func TestScaffoldDDLTableLLMError(t *testing.T) {
	dir := t.TempDir()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatal(err)
	}
	ff := &features.FeaturesFile{Tables: map[string]features.TableDef{"users": {}}}
	var out bytes.Buffer
	cfg := Config{Backend: "unsupported-backend", Model: "m"}
	err := scaffoldDDLTable(dir, dbDir, "users", ff, nil, "sysprompt", cfg, &out)
	if err == nil || !strings.Contains(err.Error(), "scaffold ddl") {
		t.Fatalf("expected scaffold ddl LLM error, got: %v", err)
	}
}
