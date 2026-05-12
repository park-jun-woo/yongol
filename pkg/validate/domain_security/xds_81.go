//ff:func feature=validate type=rule control=iteration dimension=1 topic=domain-security
//ff:what XDS-81 — internal OpenAPI endpoint에 security 선언됨 (불필요할 수 있음, WARNING)
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xds81InternalSecurity warns when internal domain endpoints have explicit
// security declarations, which are typically unnecessary for service-to-service calls.
func xds81InternalSecurity(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var diags []diagnostic.Diagnostic
	for _, dd := range docs {
		if dd.Name != "internal" || dd.Doc.Paths == nil {
			continue
		}
		for path, item := range dd.Doc.Paths.Map() {
			diags = append(diags, checkInternalPathSecurity(path, item, dd.Cfg.OpenAPI)...)
		}
	}
	return diags
}
