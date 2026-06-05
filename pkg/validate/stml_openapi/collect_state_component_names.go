//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectStateComponentNames — StateBind children에서 컴포넌트 이름 재귀 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectStateComponentNames gathers component names from a StateBind's
// children (recursively).
func collectStateComponentNames(s stml.StateBind, out map[string]struct{}) {
	for _, ch := range s.Children {
		collectChildComponentNames(ch, out)
	}
}
