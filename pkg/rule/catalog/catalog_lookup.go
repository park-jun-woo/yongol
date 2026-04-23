//ff:func feature=rule type=accessor control=sequence topic=catalog
//ff:what Catalog.Lookup — Rule ID 로 RuleMeta 조회 (없으면 second return false)
package catalog

// Lookup returns the RuleMeta for the given rule ID and whether it was found.
func (c *Catalog) Lookup(ruleID string) (RuleMeta, bool) {
	if c == nil {
		return RuleMeta{}, false
	}
	i, ok := c.byID[ruleID]
	if !ok {
		return RuleMeta{}, false
	}
	return c.Rules[i], true
}
