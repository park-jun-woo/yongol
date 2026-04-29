// stringHeads enumerates string-family PG type tokens that sqlc pgx/v5
// natively maps to Go string in NOT NULL mode.
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package types

var stringHeads = map[string]bool{
	"VARCHAR": true,
	"TEXT":    true,
	"CHAR":    true,
	"BPCHAR":  true,
}
