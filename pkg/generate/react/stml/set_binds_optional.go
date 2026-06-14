//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what ParamBind 슬라이스에서 required 집합에 없는 route.* 바인드에 Optional=true를 표시한다 (BUG-136)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// setBindsOptional sets Optional=true on every route.<Name> bind whose segment
// is not in the required set (fetch-consumed). Non-route sources and required
// names keep Optional=false.
func setBindsOptional(binds []stmlparser.ParamBind, required map[string]bool) {
	for i := range binds {
		name, ok := routeSegmentName(binds[i].Source)
		if !ok || required[name] {
			continue
		}
		binds[i].Optional = true
	}
}
