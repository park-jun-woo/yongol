//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestNestifyPath_ZeroCov — {param} → :param 변환
package nestjs

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestWritePrismaSchema_ZeroCov(t *testing.T) {
	// empty → no-op
	if err := writePrismaSchema(&yongol.Fullstack{}, t.TempDir()); err != nil {
		t.Fatalf("empty writePrismaSchema error: %v", err)
	}
	// with tables → schema file written
	backend := t.TempDir()
	fs := &yongol.Fullstack{DDLTables: []ddl.Table{{
		Name:        "users",
		ColumnOrder: []string{"id"},
		Columns:     map[string]ddl.Column{"id": {RawType: "BIGINT", NotNull: true}},
		PrimaryKey:  []string{"id"},
	}}}
	if err := writePrismaSchema(fs, backend); err != nil {
		t.Fatalf("writePrismaSchema error: %v", err)
	}
	if _, err := os.Stat(filepath.Join(backend, "prisma", "schema.prisma")); err != nil {
		t.Errorf("expected schema.prisma: %v", err)
	}
}
