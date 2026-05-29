//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD02NullableColumn_PrimaryKeyExempt — PRIMARY KEY 는 implicit NOT NULL

package ddl

import "testing"

// TestD02NullableColumn_PrimaryKeyExempt — PRIMARY KEY columns are
// implicitly NOT NULL and must not trigger [D-2] even without an
// explicit NOT NULL constraint.
func TestD02NullableColumn_PrimaryKeyExempt(t *testing.T) {
	sql := `CREATE TABLE refresh_tokens (
    token_hash TEXT PRIMARY KEY,
    claims     JSONB NOT NULL
);`
	msgs := runD02InTmpDir(t, "refresh_tokens.sql", sql, nil)
	if len(msgs) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(msgs), msgs)
	}
}
