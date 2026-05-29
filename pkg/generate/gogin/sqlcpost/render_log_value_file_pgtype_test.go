//ff:func feature=gen-gogin type=test control=sequence
//ff:what renderLogValueFile — pgtype.* 컬럼 섞여도 log/slog 단일 import 로 산출

package sqlcpost

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestRenderLogValueFile_PgtypeOnlyImportsSlog verifies that the generated
// <table>_log.go file only imports "log/slog" — the previous code used
// slog.Time which forced a "time" import even when sqlc emitted
// pgtype.Timestamp, producing an unused-import build error on top of the
// type-mismatch error described in BUG-024. With slog.Any uniformly we
// never reference time.Time or json.RawMessage at codegen level.
func TestRenderLogValueFile_PgtypeOnlyImportsSlog(t *testing.T) {
	table := ddl.Table{
		Name: "users",
		Columns: map[string]ddl.Column{
			"id":            {Name: "id", RawType: "UUID"},
			"org_id":        {Name: "org_id", RawType: "UUID"},
			"email":         {Name: "email", RawType: "VARCHAR(255)"},
			"password_hash": {Name: "password_hash", RawType: "TEXT", Sensitive: true},
			"role":          {Name: "role", RawType: "VARCHAR(20)"},
			"created_at":    {Name: "created_at", RawType: "TIMESTAMPTZ"},
		},
		ColumnOrder: []string{"id", "org_id", "email", "password_hash", "role", "created_at"},
	}
	src, err := renderLogValueFile(table)
	if err != nil {
		t.Fatalf("renderLogValueFile: %v", err)
	}
	assertRenderLogValueFileOutput(t, src)
}
