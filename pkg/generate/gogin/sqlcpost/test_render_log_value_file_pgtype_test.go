//ff:func feature=gen-gogin type=test control=sequence
//ff:what renderLogValueFile — pgtype.* 컬럼 섞여도 log/slog 단일 import 로 산출

package sqlcpost

import (
	"strings"
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

	// Import block should be exactly log/slog — nothing else.
	wantImportBlock := "import (\n\t\"log/slog\"\n)\n\n"
	if !strings.Contains(src, wantImportBlock) {
		t.Errorf("expected import block %q in output, got:\n%s", wantImportBlock, src)
	}
	for _, banned := range []string{"\"time\"", "\"encoding/json\"", "time.Time", "json.RawMessage"} {
		if strings.Contains(src, banned) {
			t.Errorf("generated source must not reference %s anymore (BUG-024):\n%s", banned, src)
		}
	}

	// LogValue body must be slog.Any + REDACTED.
	for _, line := range []string{
		"\t\tslog.Any(\"id\", r.ID),\n",
		"\t\tslog.Any(\"org_id\", r.OrgID),\n",
		"\t\tslog.Any(\"email\", r.Email),\n",
		"\t\tslog.String(\"password_hash\", \"[REDACTED]\"),\n",
		"\t\tslog.Any(\"role\", r.Role),\n",
		"\t\tslog.Any(\"created_at\", r.CreatedAt),\n",
	} {
		if !strings.Contains(src, line) {
			t.Errorf("missing line %q in generated source:\n%s", line, src)
		}
	}
}
