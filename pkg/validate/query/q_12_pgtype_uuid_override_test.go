//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what q12PgtypeUuidOverride — UUID override 검증 (nil fs/pass/누락) 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestQ12PgtypeUuidOverride_Unit(t *testing.T) {
	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := q12PgtypeUuidOverride(nil)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no UUID columns returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			SpecsDir: "/tmp/fake",
			DDLTables: []ddl.Table{
				{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id": {Name: "id", RawType: "BIGINT"},
					},
				},
			},
		}
		diags := q12PgtypeUuidOverride(fs)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("UUID column with both overrides pass", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(`sql:
  - gen:
      go:
        overrides:
          - db_type: "uuid"
            go_type:
              import: "github.com/jackc/pgx/v5/pgtype"
              type: "UUID"
            nullable: false
          - db_type: "uuid"
            go_type:
              import: "github.com/jackc/pgx/v5/pgtype"
              type: "UUID"
            nullable: true
`), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir: tmp,
			DDLTables: []ddl.Table{
				{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id": {Name: "id", RawType: "UUID"},
					},
				},
			},
		}
		diags := q12PgtypeUuidOverride(fs)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("UUID column without overrides fires", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(`sql:
  - gen:
      go:
        overrides: []
`), 0o644)
		fs := &yongol.Fullstack{
			SpecsDir: tmp,
			DDLTables: []ddl.Table{
				{
					Name: "users",
					Columns: map[string]ddl.Column{
						"id": {Name: "id", RawType: "UUID"},
					},
				},
			},
		}
		diags := q12PgtypeUuidOverride(fs)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "Q-12") {
			t.Errorf("expected Q-12, got %s", diags[0].Message)
		}
	})
}
