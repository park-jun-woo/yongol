//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what loadDomainOpenAPIDocs — Phase004 사전 파싱 결과(fs.DomainOpenAPIDocs)와 manifest 도메인 설정을 재조립해 []domainDoc 반환 (디스크 재파싱 제거)
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// loadDomainOpenAPIDocs assembles each domain's pre-parsed OpenAPI document
// (populated by ParseAll's domain loop in Phase004 into fs.DomainOpenAPIDocs)
// with its manifest DomainConfig. The rules consume Cfg.OpenAPI (diagnostic
// paths) and Cfg.Frontend (filterPagesByDomain), so Cfg must travel alongside
// the doc — the raw fs.DomainOpenAPIDocs map alone would drop it. Domains
// whose doc was not pre-parsed are skipped. No disk re-parse occurs here; the
// seven rule call-sites stay unchanged.
func loadDomainOpenAPIDocs(fs *yongol.Fullstack) []domainDoc {
	var result []domainDoc
	for name, cfg := range fs.Manifest.Domains {
		doc := fs.DomainOpenAPIDocs[name]
		if doc == nil {
			continue
		}
		result = append(result, domainDoc{Name: name, Doc: doc, Cfg: cfg})
	}
	return result
}
