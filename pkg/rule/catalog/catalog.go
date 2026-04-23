//ff:type feature=rule type=model topic=catalog
//ff:what Catalog — rulebook.md 파싱 결과를 Rule ID 기준으로 조회하는 lookup table
package catalog

// Catalog is an ordered slice of rules plus an id→index map for O(1) lookup.
// The slice order follows the rulebook.md order, which the SARIF emitter
// relies on for stable `ruleIndex` values within a run.
type Catalog struct {
	Rules []RuleMeta
	byID  map[string]int
}
