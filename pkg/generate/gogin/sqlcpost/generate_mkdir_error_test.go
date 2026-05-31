//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestGenerate — 빈 DDL/민감컬럼 유무/mkdir 에러 분기 + 파일 산출 검증
package sqlcpost

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerate_MkdirError(t *testing.T) {
	arts := t.TempDir()
	// Place a regular file where the db directory path needs to be created.
	internal := filepath.Join(arts, "backend", "internal")
	if err := os.MkdirAll(internal, 0o755); err != nil {
		t.Fatalf("setup: %v", err)
	}
	if err := os.WriteFile(filepath.Join(internal, "db"), []byte("x"), 0o644); err != nil {
		t.Fatalf("setup: %v", err)
	}
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{sensitiveTable()}}
	if err := Generate(fs, arts); err == nil {
		t.Errorf("expected mkdir error, got nil")
	}
}
