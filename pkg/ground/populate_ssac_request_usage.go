//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSSaCRequestUsage — 함수별 request 필드 사용 집합 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateSSaCRequestUsage registers per-function request field usage.
// S-51 (OpenAPI request field used in SSaC) uses this set.
func populateSSaCRequestUsage(g *rule.Ground, fs *yongol.Fullstack) {
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		used := make(rule.StringSet)
		for _, seq := range fn.Sequences {
			collectRequestFields(seq.Args, seq.Inputs, used)
		}
		g.Lookup["SSaC.requestUsage."+fn.Name] = used
	}
}
