//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what findPageByName — 페이지명(확장자 없는 STML 파일명)으로 PageSpec 포인터 조회 (없으면 nil)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// findPageByName returns the page whose Name matches, or nil.
func findPageByName(pages []stml.PageSpec, name string) *stml.PageSpec {
	for i := range pages {
		if pages[i].Name == name {
			return &pages[i]
		}
	}
	return nil
}
