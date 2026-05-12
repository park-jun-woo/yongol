//ff:func feature=validate type=rule control=iteration dimension=1 topic=domain-security
//ff:what XMO-22 — STML에서 호출한 operationId가 다른 도메인 OpenAPI에 속하면 WARNING (도메인 경계 위반)
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo22CrossDomainCall detects STML pages that call operationIds belonging
// to a different domain's OpenAPI specification (domain boundary violation).
func xmo22CrossDomainCall(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)
	if len(docs) < 2 || len(fs.STMLPages) == 0 {
		return nil
	}

	opDomain := buildOpDomainMap(docs)

	domainFrontendDir := make(map[string]string)
	for _, dd := range docs {
		if dd.Cfg.Frontend != "" {
			domainFrontendDir[dd.Name] = dd.Cfg.Frontend
		}
	}

	var diags []diagnostic.Diagnostic
	for _, page := range fs.STMLPages {
		pageDomain := classifyPageDomain(page, domainFrontendDir)
		if pageDomain == "" {
			continue
		}
		diags = append(diags, checkPageCrossDomain(page, pageDomain, opDomain)...)
	}
	return diags
}
