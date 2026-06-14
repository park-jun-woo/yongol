//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what inferInlineResponseEntity — 비 SSaC inline 응답의 property 집합 ↔ DDL 컬럼 매칭으로 엔티티 역추정 (B-2 fallback)

package openapi_ddl

// inferInlineResponseEntity is the non-SSaC fallback (strategy B-2): an inline
// response with ≥2 top-level keys is attributed to the unique DDL table whose
// columns contain every key (and which has a component schema). Ambiguous
// matches (≥2 candidate tables) yield "" to avoid misattribution.
func inferInlineResponseEntity(idx *entityIndex, keys []string) string {
	if len(keys) < 2 {
		return ""
	}
	candidate := ""
	for _, t := range idx.tables {
		allPresent := true
		for _, k := range keys {
			if _, ok := t.Columns[k]; !ok {
				allPresent = false
				break
			}
		}
		if !allPresent {
			continue
		}
		comp := idx.schemaForTable[t.Name]
		if comp == "" {
			continue
		}
		if candidate != "" && candidate != comp {
			return ""
		}
		candidate = comp
	}
	return candidate
}
