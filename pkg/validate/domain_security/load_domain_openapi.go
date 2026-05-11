//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what loadDomainOpenAPIDocs — manifest domains 키로부터 도메인별 OpenAPI 문서를 로드
package domain_security

import (
	"path/filepath"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// domainDoc holds a loaded OpenAPI document and its domain name.
type domainDoc struct {
	Name string
	Doc  *openapi3.T
	Cfg  manifest.DomainConfig
}

// loadDomainOpenAPIDocs loads each domain's OpenAPI file relative to specsDir.
// Domains whose OpenAPI file cannot be loaded are silently skipped (parse errors
// are surfaced elsewhere).
func loadDomainOpenAPIDocs(fs *yongol.Fullstack) []domainDoc {
	var result []domainDoc
	for name, cfg := range fs.Manifest.Domains {
		if cfg.OpenAPI == "" {
			continue
		}
		path := filepath.Join(fs.SpecsDir, cfg.OpenAPI)
		doc, err := openapi3.NewLoader().LoadFromFile(path)
		if err != nil {
			continue
		}
		result = append(result, domainDoc{Name: name, Doc: doc, Cfg: cfg})
	}
	return result
}
