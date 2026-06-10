//ff:func feature=stml-gen type=util control=sequence topic=import-collect
//ff:what 단일 LinkRef의 임포트 플래그(Link, route.* 소스 시 useParams)를 설정한다
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// markLinkImports flags the imports one LinkRef needs.
func markLinkImports(lr stmlparser.LinkRef, is *importSet) {
	is.useLink = true
	if linkUsesRouteParams(lr) {
		is.useParams = true
	}
}
