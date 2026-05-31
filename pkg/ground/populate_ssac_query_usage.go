//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateSSaCQueryUsage — 전역 query 필드 사용 집합 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateSSaCQueryUsage registers the aggregate set of query fields used
// across all SSaC functions. S-53 (pagination query used in SSaC) consumes this.
func populateSSaCQueryUsage(g *rule.Ground, fs *yongol.Fullstack) {
	used := make(rule.StringSet)
	for _, fn := range fs.ServiceFuncs {
		if fn.Subscribe != nil {
			continue
		}
		for _, seq := range fn.Sequences {
			collectQueryFields(seq.Args, seq.Inputs, used)
		}
	}
	g.Lookup["SSaC.queryUsage"] = used
}
