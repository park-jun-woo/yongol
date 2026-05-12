//ff:func feature=validate type=rule control=iteration dimension=1 topic=domain-security
//ff:what XMO-20 — public OpenAPI operationId가 public STML에서 미소비되면 ERROR
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo20PublicUnconsumed detects public domain OpenAPI operationIds that are
// not consumed by any STML page in the public frontend directory.
func xmo20PublicUnconsumed(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var publicDoc *domainDoc
	for i, dd := range docs {
		if dd.Name == "public" {
			publicDoc = &docs[i]
			break
		}
	}
	if publicDoc == nil || publicDoc.Doc.Paths == nil {
		return nil
	}
	if len(fs.STMLPages) == 0 {
		return nil
	}

	publicPages := filterPagesByDomain(fs.STMLPages, publicDoc.Cfg.Frontend)
	consumed := collectConsumedOpsFromPages(publicPages)

	return checkUnconsumedOps(*publicDoc, consumed, "XMO-20", "public")
}
