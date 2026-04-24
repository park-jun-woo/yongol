//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what writeLogValueFields — REDACTED sensitive + slog.Any 비민감 라인 emit

package sqlcpost

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestWriteLogValueFields_RedactedAndAny exercises the full line emission:
// sensitive columns must stay as slog.String(…, "[REDACTED]") while every
// other column — regardless of Go type — emits slog.Any. This is the
// exact shape required to build cleanly against sqlc pgx/v5 pgtype
// wrappers (BUG-024).
func TestWriteLogValueFields_RedactedAndAny(t *testing.T) {
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
	var b strings.Builder
	writeLogValueFields(&b, table, table.ColumnOrder)
	got := b.String()
	want := "" +
		"\t\tslog.Any(\"id\", r.ID),\n" +
		"\t\tslog.Any(\"org_id\", r.OrgID),\n" +
		"\t\tslog.Any(\"email\", r.Email),\n" +
		"\t\tslog.String(\"password_hash\", \"[REDACTED]\"),\n" +
		"\t\tslog.Any(\"role\", r.Role),\n" +
		"\t\tslog.Any(\"created_at\", r.CreatedAt),\n"
	if got != want {
		t.Fatalf("writeLogValueFields output mismatch:\n--- got ---\n%s--- want ---\n%s", got, want)
	}
}
