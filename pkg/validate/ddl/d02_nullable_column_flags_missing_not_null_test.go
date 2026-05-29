//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD02NullableColumn_FlagsMissingNotNull — NOT NULL 누락 시 [D-2] ERROR

package ddl

import (
	"strings"
	"testing"
)

// TestD02NullableColumn_FlagsMissingNotNull — column without NOT NULL
// and without `-- @nullable` annotation must trigger [D-2].
func TestD02NullableColumn_FlagsMissingNotNull(t *testing.T) {
	sql := `CREATE TABLE refresh_tokens (
    token_hash  TEXT        PRIMARY KEY,
    claims      JSONB       NOT NULL,
    revoked_at  TIMESTAMPTZ
);`
	msgs := runD02InTmpDir(t, "refresh_tokens.sql", sql, nil)
	if len(msgs) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(msgs), msgs)
	}
	if !strings.Contains(msgs[0], "[D-2]") || !strings.Contains(msgs[0], "revoked_at") {
		t.Errorf("diag should mention [D-2] and revoked_at, got %q", msgs[0])
	}
}
