//ff:func feature=validate type=test-helper control=sequence topic=query-structural
//ff:what writeSqlcYaml — 테스트용 최소 sqlc.yaml 생성 헬퍼 (sql_package 가변)

package query

import (
	"os"
	"path/filepath"
	"testing"
)

// writeSqlcYaml drops a minimal sqlc.yaml with the given sql_package into
// <dir>/db/sqlc.yaml. An empty pkg writes the field absent entirely.
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
