//ff:func feature=rule type=loader control=iteration dimension=1 topic=catalog
//ff:what NewCatalog — RuleMeta 슬라이스를 받아 byID 인덱스를 포함한 Catalog 반환
package catalog

// NewCatalog builds a Catalog from a pre-parsed rule slice.
func NewCatalog(rules []RuleMeta) *Catalog {
	c := &Catalog{
		Rules: rules,
		byID:  make(map[string]int, len(rules)),
	}
	for i, r := range rules {
		c.byID[r.ID] = i
	}
	return c
}
