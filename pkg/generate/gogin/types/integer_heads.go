// integerHeads enumerates the PG type tokens that map to Go int64 (sqlc
// pgx/v5 emits int64 for all of these uniformly under default settings).
// Smaller-width INT2 / INT4 still surface as int64 on the api side because
// oapi-codegen `type: integer` defaults to int64; the generator does not
// narrow types per-column today.
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package types

var integerHeads = map[string]bool{
	"BIGINT":      true,
	"BIGSERIAL":   true,
	"INTEGER":     true,
	"INT":         true,
	"INT4":        true,
	"INT8":        true,
	"INT2":        true,
	"SMALLINT":    true,
	"SERIAL":      true,
	"SMALLSERIAL": true,
}
