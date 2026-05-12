//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what classifyTokens — class 문자열을 분할하여 각 토큰 분류
package stml_design

import (
	"strings"
)

// classifyTokens splits a class string and classifies each part.
// Only non-numeric, non-builtin token names are recorded.
func classifyTokens(class, file string, out *pageTokenRefs) {
	if class == "" {
		return
	}
	for _, part := range strings.Fields(class) {
		classifySingleToken(part, class, file, out)
	}
}
