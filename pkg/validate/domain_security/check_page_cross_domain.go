//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what checkPageCrossDomain — 단일 STML 페이지에서 도메인 경계 위반 검출
package domain_security

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// checkPageCrossDomain checks a single page's fetch and action blocks for cross-domain calls.
func checkPageCrossDomain(page stml.PageSpec, pageDomain string, opDomain map[string]string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, f := range page.Fetches {
		diags = append(diags, checkFetchCrossDomain(f, page.FileName, pageDomain, opDomain)...)
	}
	for _, a := range page.Actions {
		if ownerDomain, ok := opDomain[a.OperationID]; ok && ownerDomain != pageDomain {
			diags = append(diags, diagnostic.Diagnostic{
				File:    page.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XMO-22] STML page calls operationId %q which belongs to domain %q (page is in domain %q)", a.OperationID, ownerDomain, pageDomain),
				Advice:  fmt.Sprintf("Move the endpoint to the %q domain OpenAPI, or move the page to the %q frontend", pageDomain, ownerDomain),
			})
		}
	}
	return diags
}
