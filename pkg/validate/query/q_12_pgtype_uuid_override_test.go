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

// TestQ12PgtypeUuidOverride covers the four Q-12 scenarios via a table:
//
//   1. DDL has no UUID column                                      → 0 diags (rule skipped)
//   2. DDL has UUID + sqlc.yaml has both NULL/NOT NULL overrides   → 0 diags
//   3. DDL has UUID + sqlc.yaml only has nullable=true             → 1 diag (missing nullable=false)
//   4. DDL has UUID + sqlc.yaml has no overrides at all            → 1 diag, both sides reported
//
// Per Phase001 spec: case 4 collapses the two missing entries into a
// single diagnostic ("Q-12 = one rule, one message"). The advice block
// contains both YAML stanzas so the user pastes once. Each row is run by
// runQ12PgtypeUuidOverrideCase so this func stays within the Q4 PURE
// line budget.
func TestQ12PgtypeUuidOverride(t *testing.T) {
	cases := []q12UuidTestCase{
		{
			name:      "no-uuid skips",
			ddl:       q12DDLNoUUID,
			sqlc:      q12SqlcNoOverrides,
			wantDiags: 0,
		},
		{
			name:      "both overrides pass",
			ddl:       q12DDLWithUUID,
			sqlc:      q12SqlcBothOverrides,
			wantDiags: 0,
		},
		{
			name:           "missing nullable=false",
			ddl:            q12DDLWithUUID,
			sqlc:           q12SqlcOnlyNullable,
			wantDiags:      1,
			wantMsgSubstrs: []string{"[Q-12]", "nullable=false"},
		},
		{
			name:           "both entries absent collapse to one diag",
			ddl:            q12DDLWithUUID,
			sqlc:           q12SqlcNoOverrides,
			wantDiags:      1,
			wantMsgSubstrs: []string{"[Q-12]", "nullable=false and nullable=true"},
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) { runQ12PgtypeUuidOverrideCase(t, tc) })
	}
}
