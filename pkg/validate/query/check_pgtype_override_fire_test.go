//ff:func feature=validate type=test control=sequence topic=query-structural
//ff:what checkPgtypeOverrideFire — pass/gap/only-not-null override 검증

package query

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCheckPgtypeOverride_Fire(t *testing.T) {
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

	uuidTable := []ddl.Table{
		{
			Name: "users",
			Columns: map[string]ddl.Column{
				"id": {Name: "id", RawType: "UUID"},
			},
		},
	}

	t.Run("both overrides present no diagnostic", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		writeSqlcYAML(t, dbDir, sqlcBothOverrides)

		fs := &yongol.Fullstack{SpecsDir: tmp, DDLTables: uuidTable}
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 0 {
			t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("missing overrides raises diagnostic", func(t *testing.T) {
		tmp := t.TempDir()
		dbDir := filepath.Join(tmp, "db")
		os.MkdirAll(dbDir, 0o755)
		writeSqlcYAML(t, dbDir, sqlcEmptyOverrides)

		fs := &yongol.Fullstack{SpecsDir: tmp, DDLTables: uuidTable}
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
		writeSqlcYAML(t, dbDir, sqlcNotNullOnly)

		fs := &yongol.Fullstack{SpecsDir: tmp, DDLTables: uuidTable}
		diags := checkPgtypeOverride(fs, uuidRule)
		if len(diags) != 1 {
			t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
		}
		if !strings.Contains(diags[0].Message, "nullable=true") {
			t.Errorf("expected nullable=true in message, got %s", diags[0].Message)
		}
	})
}
