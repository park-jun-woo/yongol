//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what checkInternalPathSecurity — 단일 path의 operation들에서 불필요한 security 선언 검사
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkInternalPathSecurity checks operations on an internal path for unnecessary security.
func checkInternalPathSecurity(path string, item *openapi3.PathItem, openapiFile string) []diagnostic.Diagnostic {
	ops := []*openapi3.Operation{
		item.Get, item.Post, item.Put, item.Delete, item.Patch,
	}
	var diags []diagnostic.Diagnostic
	for _, op := range ops {
		if op == nil || op.OperationID == "" {
			continue
		}
		if hasNonEmptySecurity(op) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    openapiFile,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XDS-81] internal endpoint %q (%s) has security declaration (may be unnecessary for service-to-service)", op.OperationID, path),
				Advice:  "Internal endpoints typically rely on network-level security; consider removing the security declaration if this is a service-to-service call",
			})
		}
	}
	return diags
}
