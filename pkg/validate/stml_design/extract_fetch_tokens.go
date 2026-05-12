//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what extractFetchTokens — FetchBlock에서 토큰 참조 추출
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// extractFetchTokens processes a FetchBlock.
func extractFetchTokens(fb stml.FetchBlock, file string, out *pageTokenRefs) {
	classifyTokens(fb.ClassName, file, out)
	for _, b := range fb.Binds {
		classifyTokens(b.ClassName, file, out)
	}
	for _, e := range fb.Eaches {
		extractEachTokens(e, file, out)
	}
	for _, comp := range fb.Components {
		recordComponent(comp, file, out)
		classifyTokens(comp.ClassName, file, out)
	}
	for _, c := range fb.Children {
		extractChildTokens(c, file, out)
	}
	for _, nf := range fb.NestedFetches {
		extractFetchTokens(nf, file, out)
	}
}
