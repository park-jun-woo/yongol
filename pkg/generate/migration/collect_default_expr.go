//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what collectDefaultExpr — DEFAULT 표현식 토큰 수집 (stop 키워드까지)
package migration

import "strings"

// collectDefaultExpr consumes tokens forming a DEFAULT expression.
// It stops at the next known column-level keyword (NOT, NULL, UNIQUE,
// PRIMARY, REFERENCES, CHECK, DEFAULT) and returns the joined text plus
// the number of tokens consumed.
func collectDefaultExpr(toks []string) (string, int) {
	stop := defaultExprStopKeywords()
	var parts []string
	i := 0
	for ; i < len(toks); i++ {
		if stop[strings.ToUpper(toks[i])] {
			break
		}
		parts = append(parts, toks[i])
	}
	return strings.Join(parts, " "), i
}
