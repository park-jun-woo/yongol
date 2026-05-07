//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeInt4 — nullable INTEGER 컬럼의 pgtype.Int4 매핑 (pgtypex bridge)

package types

// pgtypeInt4 returns the binding for a nullable INTEGER / INT4 column
// where sqlc pgx/v5 emits pgtype.Int4 instead of *int32.
func pgtypeInt4(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Int4",
		NeedsOverride: false,
		ApiField:      "*int64",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgInt4Ptr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgInt4Ptr({var})",
		ResponseExpr:   "pgtypex.FromPgInt4Ptr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgInt4({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
