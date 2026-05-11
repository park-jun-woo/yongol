//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XDS-81 — internal OpenAPI endpoint에 security 선언됨 (불필요할 수 있음, WARNING)
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xds81InternalSecurity warns when internal domain endpoints have explicit
// security declarations, which are typically unnecessary for service-to-service calls.
func xds81InternalSecurity(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var diags []diagnostic.Diagnostic
	for _, dd := range docs {
		if dd.Name != "internal" {
			continue
		}
		if dd.Doc.Paths == nil {
			continue
		}
		for path, item := range dd.Doc.Paths.Map() {
			ops := []*openapi3.Operation{
				item.Get, item.Post, item.Put, item.Delete, item.Patch,
			}
			for _, op := range ops {
				if op == nil || op.OperationID == "" {
					continue
				}
				if hasNonEmptySecurity(op) {
					diags = append(diags, diagnostic.Diagnostic{
						File:    dd.Cfg.OpenAPI,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelWarning,
						Message: fmt.Sprintf("[XDS-81] internal endpoint %q (%s) has security declaration (may be unnecessary for service-to-service)", op.OperationID, path),
						Advice:  "Internal endpoints typically rely on network-level security; consider removing the security declaration if this is a service-to-service call",
					})
				}
			}
		}
	}
	return diags
}

// hasNonEmptySecurity returns true when the operation has an explicit non-empty
// security requirement (at least one entry).
func hasNonEmptySecurity(op *openapi3.Operation) bool {
	return op.Security != nil && len(*op.Security) > 0
}
