//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what buildOpDomainMap — 각 operationId가 속한 도메인 매핑 생성
package domain_security

// buildOpDomainMap maps each operationId to the domain that owns it.
// If an opID appears in multiple domains, it maps to the first one found
// (XDO-90 handles the duplicate separately).
func buildOpDomainMap(docs []domainDoc) map[string]string {
	result := make(map[string]string)
	for _, dd := range docs {
		if dd.Doc.Paths == nil {
			continue
		}
		collectDocOpDomains(dd, result)
	}
	return result
}
