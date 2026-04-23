//ff:func feature=rule type=accessor control=sequence topic=catalog
//ff:what Catalog.Index — Rule ID 의 0-based index 반환 (없으면 -1)
package catalog

// Index returns the 0-based index of the given rule ID in the Rules slice.
// Returns -1 when the rule is not present.
func (c *Catalog) Index(ruleID string) int {
	if c == nil {
		return -1
	}
	if i, ok := c.byID[ruleID]; ok {
		return i
	}
	return -1
}
