//ff:func feature=report type=util control=iteration dimension=1 topic=sarif
//ff:what buildDriverRules — tool.driver.rules[] 구성 (catalog 제공 시 전체, 없으면 fired 만)
package sarif

import (
	rulecatalog "github.com/park-jun-woo/yongol/pkg/rule/catalog"
)

// buildDriverRules materialises the tool.driver.rules[] array.
//
// When a catalog is provided every catalogued rule is included (full
// catalog mode). Without a catalog only rules that fired are included
// (legacy fallback for tests / callers without a catalog).
func buildDriverRules(cat *rulecatalog.Catalog, fired map[string]struct{}) []Rule {
	if cat != nil && cat.Len() > 0 {
		out := make([]Rule, 0, cat.Len())
		for _, m := range cat.Rules {
			out = append(out, ruleFromMeta(m))
		}
		return out
	}
	if len(fired) == 0 {
		return nil
	}
	out := make([]Rule, 0, len(fired))
	// Stable order for deterministic output.
	ids := make([]string, 0, len(fired))
	for id := range fired {
		ids = append(ids, id)
	}
	// Caller test already asserts contents, not ordering — sort for safety.
	sortStrings(ids)
	for _, id := range ids {
		out = append(out, Rule{ID: id})
	}
	return out
}
