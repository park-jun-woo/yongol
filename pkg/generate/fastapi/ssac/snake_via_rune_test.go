//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestFastapiSsacHelpers_ZeroCov — addExtPkgRef/appendSnakeRune 커버
package ssac

import (
	"strings"
	"unicode"
)

func snakeViaRune(s string) string {
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		prevUpper := i > 0 && unicode.IsUpper(runes[i-1])
		nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
		appendSnakeRune(&b, i, r, prevUpper, nextLower)
	}
	return b.String()
}
