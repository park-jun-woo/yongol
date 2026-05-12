//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what extractEachTokens — EachBlock에서 토큰 참조 추출
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// extractEachTokens processes an EachBlock.
func extractEachTokens(eb stml.EachBlock, file string, out *pageTokenRefs) {
	classifyTokens(eb.ClassName, file, out)
	classifyTokens(eb.ItemClassName, file, out)
	for _, b := range eb.Binds {
		classifyTokens(b.ClassName, file, out)
	}
	for _, comp := range eb.Components {
		recordComponent(comp, file, out)
		classifyTokens(comp.ClassName, file, out)
	}
	for _, c := range eb.Children {
		extractChildTokens(c, file, out)
	}
}
