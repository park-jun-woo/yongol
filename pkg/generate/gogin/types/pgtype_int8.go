//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeInt8 — nullable BIGINT 컬럼의 pgtype.Int8 매핑 (pgtypex bridge)

package types

// pgtypeInt8 returns the binding for a nullable BIGINT / INT8 column
// where sqlc pgx/v5 emits pgtype.Int8 instead of *int64.
func pgtypeInt8(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Int8",
		NeedsOverride: false,
		ApiField:      "*int64",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgInt8Ptr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgInt8Ptr({var})",
		ResponseExpr:   "pgtypex.FromPgInt8Ptr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgInt8({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
