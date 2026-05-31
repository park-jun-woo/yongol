//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q11SqlPackagePgxV5 — Fullstack 단위 sql_package 검증 (nil fs/빈 dir/pass/fire/bad yaml) 검증
package query

import (
	"os"
	"path/filepath"
	"testing"
)

func writeSqlcYaml(t *testing.T, dir string, pkg string) {
	t.Helper()
	dbDir := filepath.Join(dir, "db")
	if err := os.MkdirAll(dbDir, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}
	var body string
	if pkg == "" {
		body = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
`
	} else {
		body = `version: "2"
sql:
  - engine: "postgresql"
    schema: "."
    queries: "queries/"
    gen:
      go:
        package: "db"
        out: "../../arts/backend/internal/db"
        sql_package: "` + pkg + `"
`
	}
	if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(body), 0o644); err != nil {
		t.Fatalf("write sqlc.yaml: %v", err)
	}
}
