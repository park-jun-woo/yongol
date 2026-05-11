//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XMO-22 — STML에서 호출한 operationId가 다른 도메인 OpenAPI에 속하면 WARNING (도메인 경계 위반)
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo22CrossDomainCall detects STML pages that call operationIds belonging
// to a different domain's OpenAPI specification (domain boundary violation).
func xmo22CrossDomainCall(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)
	if len(docs) < 2 || len(fs.STMLPages) == 0 {
		return nil
	}

	// Build opID → owning domain map.
	opDomain := buildOpDomainMap(docs)

	// Build domain → frontend dir map for page classification.
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

		// Check fetches.
		for _, f := range page.Fetches {
			diags = append(diags, checkFetchCrossDomain(f, page.FileName, pageDomain, opDomain)...)
		}
		// Check actions.
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
	}
	return diags
}

// buildOpDomainMap maps each operationId to the domain that owns it.
// If an opID appears in multiple domains, it maps to the first one found
// (XDO-90 handles the duplicate separately).
func buildOpDomainMap(docs []domainDoc) map[string]string {
	result := make(map[string]string)
	for _, dd := range docs {
		if dd.Doc.Paths == nil {
			continue
		}
		for _, item := range dd.Doc.Paths.Map() {
			ops := []*openapi3.Operation{
				item.Get, item.Post, item.Put, item.Delete, item.Patch,
			}
			for _, op := range ops {
				if op == nil || op.OperationID == "" {
					continue
				}
				if _, exists := result[op.OperationID]; !exists {
					result[op.OperationID] = dd.Name
				}
			}
		}
	}
	return result
}

// classifyPageDomain determines which domain a page belongs to based on its
// file path and the domain frontend directory configuration.
func classifyPageDomain(page stml.PageSpec, domainDirs map[string]string) string {
	for domain, dir := range domainDirs {
		pages := filterPagesByDomain([]stml.PageSpec{page}, dir)
		if len(pages) > 0 {
			return domain
		}
	}
	return ""
}

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
