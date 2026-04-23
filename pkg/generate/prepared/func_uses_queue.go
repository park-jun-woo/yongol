//ff:func feature=generate type=util control=iteration dimension=1
//ff:what funcUsesQueue — 개별 SSaC 함수가 @subscribe 또는 publish를 쓰는지

package prepared

import "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// funcUsesQueue reports whether a single ServiceFunc participates in
// the queue subsystem (either via @subscribe or a publish sequence).
func funcUsesQueue(fn ssac.ServiceFunc) bool {
	if fn.Subscribe != nil {
		return true
	}
	for _, seq := range fn.Sequences {
		if seq.Type == "publish" {
			return true
		}
	}
	return false
}
