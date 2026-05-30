//ff:func feature=gen-gogin type=test control=branch topic=sqlc-post
//ff:what TestGenerate — 빈 DDL/민감컬럼 유무/mkdir 에러 분기 + 파일 산출 검증

package sqlcpost

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func sensitiveTable() ddl.Table {
	return ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":            {Name: "id", RawType: "UUID"},
			"password_hash": {Name: "password_hash", RawType: "TEXT", Sensitive: true},
		},
		ColumnOrder: []string{"id", "password_hash"},
	}
}

func plainTable() ddl.Table {
	return ddl.Table{
		Name: "logs",
		Columns: map[string]ddl.Column{
			"id": {Name: "id", RawType: "UUID"},
		},
		ColumnOrder: []string{"id"},
	}
}

func TestGenerate_NoTables(t *testing.T) {
	if err := Generate(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Errorf("empty DDLTables should return nil, got: %v", err)
	}
}

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
