//ff:func feature=validate type=util control=iteration dimension=3 topic=stml-openapi
//ff:what collectPageComponentNames — 단일 페이지의 fetch/action/children에서 컴포넌트 이름 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectPageComponentNames walks a single page's fetches, actions, and
// top-level children to gather component names.
func collectPageComponentNames(page stml.PageSpec, out map[string]struct{}) {
	for _, f := range page.Fetches {
		collectFetchComponentNames(f, out)
	}
	for _, a := range page.Actions {
		collectActionComponentNames(a, out)
	}
	for _, ch := range page.Children {
		collectChildComponentNames(ch, out)
	}
}
