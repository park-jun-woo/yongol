//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeBool — nullable BOOLEAN 컬럼의 pgtype.Bool 매핑 (pgtypex bridge)

package types

// pgtypeBool returns the binding for a nullable BOOLEAN column where sqlc
// pgx/v5 emits pgtype.Bool instead of *bool.
func pgtypeBool(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:    "pgtype.Bool",
		NeedsOverride: false,
		ApiField:      "*bool",
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
			"github.com/park-jun-woo/ssac/pkg/pgtypex",
		},
		ConvertExpr:    "pgtypex.FromPgBoolPtr({row}.{field})",
		InsertExpr:     "pgtypex.ToPgBoolPtr({var})",
		ResponseExpr:   "pgtypex.FromPgBoolPtr({var}.{field})",
		NilCheckExpr:   "pgtypex.IsNilPgBool({var})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
