//ff:func feature=validate type=rule control=iteration dimension=1 topic=domain-security
//ff:what XMO-21 — admin OpenAPI operationId가 admin STML에서 미소비되면 ERROR (admin frontend 존재 시)
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo21AdminUnconsumed detects admin domain OpenAPI operationIds that are
// not consumed by any STML page in the admin frontend directory.
// Only fires when the admin domain has a frontend directory configured.
func xmo21AdminUnconsumed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var adminDoc *domainDoc
	for i, dd := range docs {
		if dd.Name == "admin" {
			adminDoc = &docs[i]
			break
		}
	}
	if adminDoc == nil || adminDoc.Doc.Paths == nil {
		return nil
	}
	if adminDoc.Cfg.Frontend == "" {
		return nil
	}
	if len(fs.STMLPages) == 0 {
		return nil
	}

	adminPages := filterPagesByDomain(fs.STMLPages, adminDoc.Cfg.Frontend)
	consumed := collectConsumedOpsFromPages(adminPages)

	return checkUnconsumedOps(*adminDoc, consumed, "XMO-21", "admin")
}
