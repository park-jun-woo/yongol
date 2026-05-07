//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeFloat4 — nullable REAL 컬럼의 pgtype.Float4 매핑 (pgtypex bridge)

package types

// pgtypeFloat4 returns the binding for a nullable REAL / FLOAT4 column
// where sqlc pgx/v5 emits pgtype.Float4 instead of *float32.
func pgtypeFloat4(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Float4",
		NeedsOverride: false,
		ApiField:      "*float64",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgFloat4Ptr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgFloat4Ptr({var})",
		ResponseExpr:   "pgtypex.FromPgFloat4Ptr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgFloat4({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
