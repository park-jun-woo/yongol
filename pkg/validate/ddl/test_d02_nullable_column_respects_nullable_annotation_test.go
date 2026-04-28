//ff:func feature=validate type=test control=sequence topic=ddl-structural
//ff:what TestD02NullableColumn_RespectsNullableAnnotation — `-- @nullable` 면제 확인 (BUG-028)

package ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// TestD02NullableColumn_RespectsNullableAnnotation — BUG-028 regression.
// When the parser has recorded `-- @nullable` on a column, D-2 must not
// fire for that column.
func TestD02NullableColumn_RespectsNullableAnnotation(t *testing.T) {
	sql := `CREATE TABLE refresh_tokens (
    token_hash  TEXT        PRIMARY KEY,
    claims      JSONB       NOT NULL,
    revoked_at  TIMESTAMPTZ -- @nullable
);`
	tables := []ddl.Table{{
		Name: "refresh_tokens",
		Columns: map[string]ddl.Column{
			"revoked_at": {Name: "revoked_at", RawType: "TIMESTAMPTZ", NullableAnnot: true},
		},
	}}
	msgs := runD02InTmpDir(t, "refresh_tokens.sql", sql, tables)
	if len(msgs) != 0 {
		t.Fatalf("expected 0 diagnostics (column is @nullable-annotated), got %d: %+v", len(msgs), msgs)
	}
}
