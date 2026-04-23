//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what splitTrailingComment — SQL 라인에서 `-- ...` 주석을 DDL 본문과 분리

package ddl

import "strings"

// splitTrailingComment returns (ddl, comment) where comment is the text
// after `-- ` on the same line (without the `--` prefix), respecting
// single-quoted string literals.
func splitTrailingComment(line string) (string, string) {
	inSQ := false
	i := 0
	for i < len(line)-1 {
		c := line[i]
		if c == '\'' {
			inSQ, i = advanceSingleQuote(line, i, inSQ)
			continue
		}
		if !inSQ && c == '-' && line[i+1] == '-' {
			return line[:i], strings.TrimSpace(line[i+2:])
		}
		i++
	}
	return line, ""
}
