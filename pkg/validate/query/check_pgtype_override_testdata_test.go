//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what checkPgtypeOverrideTestdata — sqlc YAML 테스트 데이터 및 헬퍼

package query

import (
	"os"
	"path/filepath"
	"testing"
)

const sqlcBothOverrides = `sql:
  - gen:
      go:
        overrides:
          - db_type: "uuid"
            go_type:
              import: "github.com/jackc/pgx/v5/pgtype"
              package: "pgtype"
              type: "UUID"
            nullable: false
          - db_type: "uuid"
            go_type:
              import: "github.com/jackc/pgx/v5/pgtype"
              package: "pgtype"
              type: "UUID"
            nullable: true
`

const sqlcEmptyOverrides = `sql:
  - gen:
      go:
        overrides: []
`

const sqlcNotNullOnly = `sql:
  - gen:
      go:
        overrides:
          - db_type: "uuid"
            go_type:
              import: "github.com/jackc/pgx/v5/pgtype"
              package: "pgtype"
              type: "UUID"
            nullable: false
`

func writeSqlcYAML(t *testing.T, dbDir, content string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(content), 0o644); err != nil {
		t.Fatal(err)
	}
}
