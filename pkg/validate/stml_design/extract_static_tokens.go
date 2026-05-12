//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what extractStaticTokens — StaticElement에서 재귀적 토큰 추출
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// extractStaticTokens processes a StaticElement recursively.
func extractStaticTokens(se stml.StaticElement, file string, out *pageTokenRefs) {
	classifyTokens(se.ClassName, file, out)
	for _, c := range se.Children {
		extractChildTokens(c, file, out)
	}
}
