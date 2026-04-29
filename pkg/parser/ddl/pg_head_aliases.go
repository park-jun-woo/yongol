// pgHeadAliases maps multi-word PostgreSQL type names (upper-case,
// joined by single space) to their canonical single-token alias. Both
// forms are idiomatic in PostgreSQL; downstream matrices (types/ family
// heads, validate/query head comparisons) are keyed by the single-token
// form, so consumers normalize via NormalizePGTypeHead.
//
// Const-only file — filefunc skips //ff annotations on const/var-only
// files.

package ddl

var pgHeadAliases = map[string]string{
	"DOUBLE PRECISION":            "FLOAT8",
	"TIMESTAMP WITH TIME ZONE":    "TIMESTAMPTZ",
	"TIMESTAMP WITHOUT TIME ZONE": "TIMESTAMP",
	"TIME WITH TIME ZONE":         "TIMETZ",
	"TIME WITHOUT TIME ZONE":      "TIME",
	"CHARACTER VARYING":           "VARCHAR",
	"CHARACTER":                   "CHAR", // single-word, registered for symmetry
	"BIT VARYING":                 "VARBIT",
}
