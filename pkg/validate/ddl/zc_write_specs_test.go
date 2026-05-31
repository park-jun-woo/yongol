//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestZeroCov — 0% DDL rule 함수 (d01/d03/d04/d05/d06/d07/d15/xdd61/parseSqlcYaml/Run) 회귀
package ddl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func zcWriteSpecs(t *testing.T, sqlcYaml, schemaSQL string) *yongol.Fullstack {
	t.Helper()
	tmp := t.TempDir()
	dbDir := filepath.Join(tmp, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	if sqlcYaml != "" {
		if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(sqlcYaml), 0o644); err != nil {
			t.Fatalf("write sqlc.yaml: %v", err)
		}
	}
	if schemaSQL != "" {
		if err := os.WriteFile(filepath.Join(dbDir, "schema.sql"), []byte(schemaSQL), 0o644); err != nil {
			t.Fatalf("write schema.sql: %v", err)
		}
	}
	return &yongol.Fullstack{
		SpecsDir:  tmp,
		Presences: map[yongol.SSOTKind]yongol.SSOTPresence{yongol.KindDDL: yongol.SSOTPopulated},
	}
}
