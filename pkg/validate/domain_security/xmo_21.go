//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XMO-21 — admin OpenAPI operationId가 admin STML에서 미소비되면 ERROR (admin frontend 존재 시)
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
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
	// Skip if admin domain has no frontend configured.
	if adminDoc.Cfg.Frontend == "" {
		return nil
	}
	if len(fs.STMLPages) == 0 {
		return nil
	}

	// Filter STML pages belonging to the admin frontend directory.
	adminPages := filterPagesByDomain(fs.STMLPages, adminDoc.Cfg.Frontend)
	consumed := collectConsumedOpsFromPages(adminPages)

	var diags []diagnostic.Diagnostic
	for _, item := range adminDoc.Doc.Paths.Map() {
		ops := []*openapi3.Operation{
			item.Get, item.Post, item.Put, item.Delete, item.Patch,
		}
		for _, op := range ops {
			if op == nil || op.OperationID == "" {
				continue
			}
			// Skip auth endpoints.
			if hasEmptySecurity(op) {
				continue
			}
			if _, ok := consumed[op.OperationID]; !ok {
				diags = append(diags, diagnostic.Diagnostic{
					File:    adminDoc.Cfg.OpenAPI,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[XMO-21] admin operationId %q is not consumed by any STML page in the admin frontend", op.OperationID),
					Advice:  fmt.Sprintf("Add a data-fetch or data-action referencing %q in an admin STML page, or remove the endpoint from the admin OpenAPI", op.OperationID),
				})
			}
		}
	}
	return diags
}
