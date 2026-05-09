//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what inferSqlcParamIntWidth — infer the integer width of a sqlc param from the query body cast

package ssac_sqlc

import "regexp"

// inferSqlcParamIntWidth inspects the raw SQL body for a named sqlc param and
// returns its integer width ("int32" or "int64") based on the PostgreSQL type
// cast following the param reference.
//
// Patterns recognised:
//
//	sqlc.arg(name)::bigint  or  @name::bigint  → "int64"
//	sqlc.arg(name)::int8                        → "int64"
//	sqlc.arg(name)::int     or  ::int4 / ::integer → "int32"
//	sqlc.arg(name)          (no cast)           → "int32"  (PG LIMIT/OFFSET default)
//
// Whitespace around `::` is tolerated.
func inferSqlcParamIntWidth(body, paramName string) string {
	// Build two patterns: one for sqlc.arg(name), one for @name.
	// Capture optional :: cast after the param reference.
	sqlcArgPat := regexp.MustCompile(
		`sqlc\.arg\(\s*` + regexp.QuoteMeta(paramName) + `\s*\)\s*(?:::\s*(\w+))?`,
	)
	atPat := regexp.MustCompile(
		`@` + regexp.QuoteMeta(paramName) + `\b\s*(?:::\s*(\w+))?`,
	)

	for _, pat := range []*regexp.Regexp{sqlcArgPat, atPat} {
		m := pat.FindStringSubmatch(body)
		if m == nil {
			continue
		}
		return castToIntWidth(m[1])
	}
	return ""
}

// castToIntWidth maps a PostgreSQL type-cast token to "int32" or "int64".
// Empty token (no cast) defaults to "int32".
func castToIntWidth(cast string) string {
	switch cast {
	case "bigint", "int8":
		return "int64"
	case "int", "int4", "integer", "":
		return "int32"
	default:
		// Non-integer cast (e.g. ::text) — not an int param.
		return ""
	}
}
