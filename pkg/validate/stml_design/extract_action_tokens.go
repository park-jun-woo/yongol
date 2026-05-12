//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what extractActionTokens — ActionBlock에서 토큰 참조 추출
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// extractActionTokens processes an ActionBlock.
func extractActionTokens(ab stml.ActionBlock, file string, out *pageTokenRefs) {
	classifyTokens(ab.ClassName, file, out)
	for _, f := range ab.Fields {
		classifyTokens(f.ClassName, file, out)
	}
	for _, c := range ab.Children {
		extractChildTokens(c, file, out)
	}
}
