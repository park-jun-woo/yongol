//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-openapi
//ff:what STML 페이지들에서 auth.refresh 캡처 존재 여부를 판정한다

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

func pagesHaveRefreshCapture(pages []stml.PageSpec) bool {
	for _, p := range pages {
		if actionsHaveRefreshCapture(p.Actions) {
			return true
		}
	}
	return false
}
