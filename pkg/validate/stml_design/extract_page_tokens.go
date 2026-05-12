//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what extractPageTokens — 단일 STML 페이지에서 토큰 참조 추출
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// extractPageTokens processes a single page.
func extractPageTokens(page stml.PageSpec, out *pageTokenRefs) {
	for _, f := range page.Fetches {
		extractFetchTokens(f, page.FileName, out)
	}
	for _, a := range page.Actions {
		extractActionTokens(a, page.FileName, out)
	}
	for _, c := range page.Children {
		extractChildTokens(c, page.FileName, out)
	}
}
