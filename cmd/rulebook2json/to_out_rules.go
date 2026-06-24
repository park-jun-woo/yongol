//ff:func feature=rule type=command control=iteration dimension=1 topic=catalog
//ff:what catalog.Catalog 의 RuleMeta 목록을 outRule 슬라이스로 변환한다 (rulebook 순서 보존).
package main

import "github.com/park-jun-woo/yongol/pkg/rule/catalog"

// toOutRules projects the parsed catalog's RuleMeta entries onto the lowercase
// outRule JSON shape, preserving rulebook order.
func toOutRules(cat *catalog.Catalog) []outRule {
	rules := make([]outRule, 0, len(cat.Rules))
	for _, r := range cat.Rules {
		rules = append(rules, outRule{
			ID:      r.ID,
			Level:   r.Level,
			Desc:    r.Description,
			Source:  r.Source,
			Section: r.SectionTitle,
		})
	}
	return rules
}
