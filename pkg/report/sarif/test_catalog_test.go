//ff:func feature=report type=test control=sequence topic=sarif
//ff:what TestAttachRuleIndex — nil cat / 빈 ruleID / 미존재 / 존재 시 RuleIndex 설정 분기 검증
package sarif

import (
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

func testCatalog() *rulecatalog.Catalog {
	return rulecatalog.NewCatalog([]rulecatalog.RuleMeta{
		{ID: "S-1", Level: "ERROR"},
		{ID: "S-2", Level: "WARNING"},
		{ID: "X-3", Level: "ERROR"},
	})
}
