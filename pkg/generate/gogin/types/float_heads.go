// floatHeads enumerates float-family PG type tokens that sqlc pgx/v5
// natively maps to float64. The single-word "DOUBLE" entry was removed
// when parser/ddl gained multi-word PG type support: the parser now
// preserves "DOUBLE PRECISION" verbatim and parseRawType normalises it
// to the canonical alias "FLOAT8" via ddl.NormalizePGTypeHead. The
// lone token "DOUBLE" is not a valid PostgreSQL type name and would
// previously match here by accident, hiding parser truncation bugs.
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package types

var floatHeads = map[string]bool{
	"REAL":   true,
	"FLOAT":  true,
	"FLOAT4": true,
	"FLOAT8": true,
}
