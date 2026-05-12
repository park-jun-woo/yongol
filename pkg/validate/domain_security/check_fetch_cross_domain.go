//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what checkFetchCrossDomain — FetchBlock에서 재귀적으로 도메인 경계 위반 검출
package domain_security

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkFetchCrossDomain recursively checks fetch blocks for cross-domain calls.
func checkFetchCrossDomain(f stml.FetchBlock, fileName, pageDomain string, opDomain map[string]string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	if ownerDomain, ok := opDomain[f.OperationID]; ok && ownerDomain != pageDomain {
		diags = append(diags, diagnostic.Diagnostic{
			File:    fileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[XMO-22] STML page calls operationId %q which belongs to domain %q (page is in domain %q)", f.OperationID, ownerDomain, pageDomain),
			Advice:  fmt.Sprintf("Move the endpoint to the %q domain OpenAPI, or move the page to the %q frontend", pageDomain, ownerDomain),
		})
	}
	for _, child := range f.NestedFetches {
		diags = append(diags, checkFetchCrossDomain(child, fileName, pageDomain, opDomain)...)
	}
	return diags
}
