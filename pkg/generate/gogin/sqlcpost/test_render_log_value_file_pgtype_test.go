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
		Columns: map[string]string{
			"id":            "pgtype.UUID",
			"org_id":        "pgtype.UUID",
			"email":         "string",
			"password_hash": "string",
			"role":          "string",
			"created_at":    "pgtype.Timestamp",
		},
		ColumnOrder: []string{"id", "org_id", "email", "password_hash", "role", "created_at"},
		SensitiveColumns: map[string]bool{
			"password_hash": true,
		},
	}
	src, err := renderLogValueFile(table)
	if err != nil {
		t.Fatalf("renderLogValueFile: %v", err)
	}
	assertRenderLogValueFileOutput(t, src)
}
