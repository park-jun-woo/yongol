//ff:func feature=report type=util control=sequence topic=sarif
//ff:what attachRuleIndex — catalog 에 존재하는 ruleID 라면 Result.RuleIndex 설정

package sarif

import (
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

// attachRuleIndex fills Result.RuleIndex from the catalog's stable slice
// position when both a catalog and a non-empty ruleID are available. No-op
// otherwise.
func attachRuleIndex(res *Result, cat *rulecatalog.Catalog, ruleID string) {
	if cat == nil || ruleID == "" {
		return
	}
	idx := cat.Index(ruleID)
	if idx < 0 {
		return
	}
	i := idx
	res.RuleIndex = &i
}
