//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeInt2 — nullable SMALLINT 컬럼의 pgtype.Int2 매핑 (pgtypex bridge)

package types

// pgtypeInt2 returns the binding for a nullable SMALLINT / INT2 column
// where sqlc pgx/v5 emits pgtype.Int2 instead of *int16.
func pgtypeInt2(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Int2",
		NeedsOverride: false,
		ApiField:      "*int64",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgInt2Ptr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgInt2Ptr({var})",
		ResponseExpr:   "pgtypex.FromPgInt2Ptr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgInt2({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
