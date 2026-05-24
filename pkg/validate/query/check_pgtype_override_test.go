//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what checkPgtypeOverride — nil fs/empty specsdir/no matching col/no file/bad yaml/pass/gap 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCheckPgtypeOverride(t *testing.T) {
	uuidFilter := func(col ddl.Column) bool {
		return headTokenEquals(col.RawType, "UUID")
	}
	uuidRule := pgtypeOverrideRule{
		RuleID:    "Q-12",
		DBType:    "uuid",
		PgPackage: "pgtype",
		PgType:    "UUID",
		Filter:    uuidFilter,
		Advice:    "Add uuid override",
	}

	t.Run("nil fullstack returns nil", func(t *testing.T) {
		diags := checkPgtypeOverride(nil, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("empty SpecsDir returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{}
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("no matching DDL column returns nil", func(t *testing.T) {
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
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("sqlc.yaml not found returns nil", func(t *testing.T) {
		tmp := t.TempDir()
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
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("invalid yaml returns nil", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(":::bad yaml"), 0o644)

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
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected nil, got %+v", diags)
		}
	})

	t.Run("both overrides present no diagnostic", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		sqlcYAML := `sql:
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
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(sqlcYAML), 0o644)

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
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing overrides raises diagnostic", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		sqlcYAML := `sql:
  - gen:
      go:
        overrides: []
`
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(sqlcYAML), 0o644)

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
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "Q-12") {
			t.Errorf("Message missing Q-12: %s", diags[0].Message)
		}
	})

	t.Run("only not-null override raises diagnostic for nullable", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		sqlcYAML := `sql:
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
		os.WriteFile(filepath.Join(dbDir, "sqlc.yaml"), []byte(sqlcYAML), 0o644)

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
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "nullable=true") {
			t.Errorf("expected nullable=true in message, got %s", diags[0].Message)
		}
	})
}
