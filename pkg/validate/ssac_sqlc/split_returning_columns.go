//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what splitReturningColumns — RETURNING 컬럼 리스트 문자열을 lowercased name set 으로 분할

package ssac_sqlc

import "strings"

// splitReturningColumns splits a RETURNING column-list string on commas,
// trims whitespace, strips a leading `tablealias.` prefix, peels an
// `AS alias` suffix, lowercases the remainder, and returns the resulting
// set. Examples (all → {"id": true, "email": true}):
//
//	"id, email"
//	"u.id, u.email"
//	"id AS user_id, email"
//
// Asterisk is not handled here — callers short-circuit on `*` upstream.
// Quoted identifiers and expressions like `count(*)` are out of scope: the
// parser treats them as literal tokens, which simply forces a "partial"
// classification (a safe outcome for the XQS-20 check).
func splitReturningColumns(clause string) map[string]bool {
	out := make(map[string]bool)
	for _, raw := range strings.Split(clause, ",") {
		col := strings.TrimSpace(raw)
		if col == "" {
			continue
		}
		// Strip leading `alias.` — keep the right-hand side.
		if dot := strings.LastIndex(col, "."); dot >= 0 && dot < len(col)-1 {
			col = col[dot+1:]
		}
		// Strip ` AS alias` — keep the alias as the effective column name.
		if asIdx := indexCaseFold(col, " AS "); asIdx >= 0 {
			col = strings.TrimSpace(col[asIdx+len(" AS "):])
		}
		col = strings.Trim(col, "\" `")
		if col == "" {
			continue
		}
		out[strings.ToLower(col)] = true
	}
	return out
}
