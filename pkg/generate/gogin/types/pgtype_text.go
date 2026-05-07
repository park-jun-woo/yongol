//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeText — nullable TEXT 컬럼의 pgtype.Text 매핑 (pgtypex bridge)

package types

// pgtypeText returns the binding for a nullable TEXT / VARCHAR column
// where sqlc pgx/v5 emits pgtype.Text instead of *string.
func pgtypeText(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Text",
		NeedsOverride: false,
		ApiField:      "*string",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgTextPtr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgTextPtr({var})",
		ResponseExpr:   "pgtypex.FromPgTextPtr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgText({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
