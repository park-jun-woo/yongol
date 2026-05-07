//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeFloat8 — nullable DOUBLE PRECISION 컬럼의 pgtype.Float8 매핑 (pgtypex bridge)

package types

// pgtypeFloat8 returns the binding for a nullable DOUBLE PRECISION /
// FLOAT8 column where sqlc pgx/v5 emits pgtype.Float8 instead of *float64.
func pgtypeFloat8(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Float8",
		NeedsOverride: false,
		ApiField:      "*float64",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgFloat8Ptr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgFloat8Ptr({var})",
		ResponseExpr:   "pgtypex.FromPgFloat8Ptr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgFloat8({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
