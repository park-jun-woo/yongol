//ff:func feature=rule type=accessor control=sequence topic=catalog
//ff:what Catalog.Len — Catalog 내 규칙 수
package catalog

// Len returns the number of rules in the catalog.
func (c *Catalog) Len() int {
	if c == nil {
		return 0
	}
	return len(c.Rules)
}
