//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what extractReturningClause — sqlc 쿼리 raw body 에서 RETURNING 절 추출 (light-weight)

package ssac_sqlc

import (
	"regexp"
	"strings"
)

// returningRe captures the column list following a top-level RETURNING keyword.
// The body may span multiple lines; we strip inline `-- ...` comments and
// collapse whitespace before matching. The capture extends up to the
// terminating `;` or end-of-body, whichever comes first.
var returningRe = regexp.MustCompile(`(?i)\bRETURNING\s+([^;]+)`)

// extractReturningClause returns the raw column list captured after the
// RETURNING keyword in body, or "" when the body has no RETURNING clause.
// Inline `-- ...` comments are stripped before matching so a trailing
// `RETURNING id, email; -- partial` is parsed cleanly.
//
// Examples:
//
//	"INSERT INTO users (...) VALUES (...) RETURNING *;"        → "*"
//	"INSERT INTO users (...) VALUES (...) RETURNING id, email" → "id, email"
//	"SELECT * FROM users WHERE id = @id;"                      → ""
func extractReturningClause(body string) string {
	if body == "" {
		return ""
	}
	stripped := stripSQLLineComments(body)
	m := returningRe.FindStringSubmatch(stripped)
	if m == nil {
		return ""
	}
	return strings.TrimSpace(m[1])
}
