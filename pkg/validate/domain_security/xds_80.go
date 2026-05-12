//ff:func feature=validate type=rule control=iteration dimension=1 topic=domain-security
//ff:what XDS-80 — admin OpenAPI endpoint가 security: [] (public 접근 허용) 이면 ERROR
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xds80AdminPublicAccess detects admin domain endpoints that allow public
// access via explicit empty security (security: []).
func xds80AdminPublicAccess(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var diags []diagnostic.Diagnostic
	for _, dd := range docs {
		if dd.Name != "admin" || dd.Doc.Paths == nil {
			continue
		}
		for path, item := range dd.Doc.Paths.Map() {
			diags = append(diags, checkAdminPathSecurity(path, item, dd.Cfg.OpenAPI)...)
		}
	}
	return diags
}
