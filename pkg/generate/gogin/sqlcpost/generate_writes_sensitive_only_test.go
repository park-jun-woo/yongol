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

func TestGenerate_WritesSensitiveOnly(t *testing.T) {
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{sensitiveTable(), plainTable()}}
	arts := t.TempDir()
	if err := Generate(fs, arts); err != nil {
		t.Fatalf("Generate: %v", err)
	}
	dbDir := filepath.Join(arts, "backend", "internal", "db")
	if _, err := os.Stat(filepath.Join(dbDir, "users_log.go")); err != nil {
		t.Errorf("expected users_log.go for sensitive table: %v", err)
	}
	if _, err := os.Stat(filepath.Join(dbDir, "logs_log.go")); !os.IsNotExist(err) {
		t.Errorf("did not expect logs_log.go for non-sensitive table, stat err = %v", err)
	}
}
