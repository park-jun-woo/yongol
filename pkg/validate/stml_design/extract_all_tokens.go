//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what extractAllTokens — 전체 STML 페이지에서 커스텀 토큰 참조 수집
package stml_design

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// extractAllTokens collects custom token references from all STML pages.
func extractAllTokens(fs *yongol.Fullstack) pageTokenRefs {
	var result pageTokenRefs
	for _, page := range fs.STMLPages {
		extractPageTokens(page, &result)
	}
	return result
}
