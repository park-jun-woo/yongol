// floatHeads enumerates float-family PG type tokens that sqlc pgx/v5
// natively maps to float64. DOUBLE PRECISION is intentionally absent —
// the parser preserves it as a multi-word token (parseRawType.MultiToken)
// and dispatch routes it to KindUnsupported. Single-word "DOUBLE" is
// rare in practice but tolerated here for the few users who write it.
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package types

var floatHeads = map[string]bool{
	"REAL":   true,
	"FLOAT":  true,
	"FLOAT4": true,
	"FLOAT8": true,
	"DOUBLE": true,
}
