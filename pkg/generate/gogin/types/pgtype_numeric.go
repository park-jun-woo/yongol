//ff:func feature=gen-gogin type=util control=sequence
//ff:what pgtypeNumeric — NUMERIC/DECIMAL 컬럼의 pgtype.Numeric 매핑 (NeedsOverride=true)

package types

// pgtypeNumeric returns the binding for a NUMERIC / DECIMAL column. sqlc
// pgx/v5 emits pgtype.Numeric for both NULL and NOT NULL by default; the
// override is still required because some users pin sqlc to legacy modes
// where pgtype.Numeric is not the default. The convert site routes
// through pgNumericToString which calls `.Value()` and stringifies.
//
// The api surface uses string for NUMERIC because oapi-codegen does not
// have a built-in arbitrary-precision decimal type. Callers needing
// numeric arithmetic should parse on the consumer side.
func pgtypeNumeric(notNull bool, defaultLiteral string) GoTypeBinding {
	apiField := "string"
	if !notNull {
		apiField = "*string"
	}
	return GoTypeBinding{
		SqlcGoType:     "pgtype.Numeric",
		NeedsOverride:  true,
		ApiField:       apiField,
		Imports: []string{
			"github.com/jackc/pgx/v5/pgtype",
		},
		ConvertExpr:    "pgNumericToString({row}.{field})",
		InsertExpr:     "{var}",
		ResponseExpr:   "pgNumericToString({var}.{field})",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPgtype,
		Supported:      true,
	}
}
