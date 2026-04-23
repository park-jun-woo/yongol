//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what splitTrailingComment — SQL 라인에서 `-- ...` 주석을 DDL 본문과 분리

package ddl

import "strings"

// splitTrailingComment returns (ddl, comment) where comment is the text
// after `-- ` on the same line (without the `--` prefix), respecting
// single-quoted string literals.
func splitTrailingComment(line string) (string, string) {
	inSQ := false
	for i := 0; i < len(line)-1; i++ {
		c := line[i]
		if c == '\'' {
			if inSQ && i+1 < len(line) && line[i+1] == '\'' {
				i++
				continue
			}
			inSQ = !inSQ
			continue
		}
		if !inSQ && c == '-' && line[i+1] == '-' {
			return line[:i], strings.TrimSpace(line[i+2:])
		}
	}
	return line, ""
}
