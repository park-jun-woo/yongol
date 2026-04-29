//ff:func feature=manifest type=util control=sequence
//ff:what NormalizePGTypeHead — 다중 토큰 PG 타입 head 를 단일 토큰 별칭으로 환산

package ddl

import "strings"

// NormalizePGTypeHead returns the canonical single-token alias for the
// given head if it is a recognized multi-word PostgreSQL type;
// otherwise returns the head as-is (after upper-casing and trimming).
//
// The head is the portion of Column.RawType with array marker "[]" and
// parameter list "(...)" already stripped. Example mappings:
//
//	"DOUBLE PRECISION"            → "FLOAT8"
//	"TIMESTAMP WITH TIME ZONE"    → "TIMESTAMPTZ"
//	"TIMESTAMP WITHOUT TIME ZONE" → "TIMESTAMP"
//	"CHARACTER VARYING"           → "VARCHAR"
//	"BIGINT"                      → "BIGINT"  (no alias)
//
// The mapping table lives in pg_head_aliases.go (const-only) so it can
// be audited as a single SSOT shared by gen/types and validate/query.
func NormalizePGTypeHead(head string) string {
	upper := strings.ToUpper(strings.TrimSpace(head))
	if alias, ok := pgHeadAliases[upper]; ok {
		return alias
	}
	return upper
}
