//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q11SqlPackagePgxV5 — Fullstack 단위 sql_package 검증 (nil fs/빈 dir/pass/fire/bad yaml) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ11SqlPackagePgxV5_Unit(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := q11SqlPackagePgxV5(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty SpecsDir returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := q11SqlPackagePgxV5(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("sqlc.yaml not found returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{SpecsDir: t.TempDir()}
		diags := q11SqlPackagePgxV5(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("bad yaml returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(":::bad"), 0o644)

		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := q11SqlPackagePgxV5(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("pgx/v5 returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(`sql:
  - gen:
      go:
        sql_package: pgx/v5
`), 0o644)

		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := q11SqlPackagePgxV5(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("database/sql fires", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(`sql:
  - gen:
      go:
        sql_package: database/sql
`), 0o644)

		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := q11SqlPackagePgxV5(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "[Q-11]") {
			t.Errorf("expected Q-11, got %s", diags[0].Message)
		}
	})

	t.Run("absent sql_package fires", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(`sql:
  - gen:
      go:
        out: db
`), 0o644)

		fs := &yongol.Fullstack{SpecsDir: tmp}
		diags := q11SqlPackagePgxV5(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "absent") {
			t.Errorf("expected absent mention, got %s", diags[0].Message)
		}
	})
}

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
