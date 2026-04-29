//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what stripSQLLineComments — SQL 본문에서 `--` 라인 주석 제거 (XQS-20 RETURNING 추출 전처리)

package ssac_sqlc

import "strings"

// stripSQLLineComments removes everything from `--` to end-of-line on every
// line of body. Block comments `/* ... */` are not handled (out of scope —
// sqlc query bodies in this project consistently use `--` line comments).
// Used by extractReturningClause to keep `RETURNING id; -- partial` parseable.
func stripSQLLineComments(body string) string {
	lines := strings.Split(body, "\n")
	for i, ln := range lines {
		if idx := strings.Index(ln, "--"); idx >= 0 {
			lines[i] = ln[:idx]
		}
	}
	return strings.Join(lines, "\n")
}
