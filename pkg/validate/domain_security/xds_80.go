//ff:func feature=validate type=rule control=iteration dimension=2 topic=domain-security
//ff:what XDS-80 — admin OpenAPI endpoint가 security: [] (public 접근 허용) 이면 ERROR
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xds80AdminPublicAccess detects admin domain endpoints that allow public
// access via explicit empty security (security: []).
func xds80AdminPublicAccess(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	docs := loadDomainOpenAPIDocs(fs)

	var diags []diagnostic.Diagnostic
	for _, dd := range docs {
		if dd.Name != "admin" {
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
				if hasEmptySecurity(op) {
					diags = append(diags, diagnostic.Diagnostic{
						File:    dd.Cfg.OpenAPI,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelError,
						Message: fmt.Sprintf("[XDS-80] admin endpoint %q (%s) has security: [] allowing public access", op.OperationID, path),
						Advice:  "Remove the empty security override or add an appropriate security requirement to the admin endpoint",
					})
				}
			}
		}
	}
	return diags
}

// hasEmptySecurity returns true when the operation has explicit `security: []`.
func hasEmptySecurity(op *openapi3.Operation) bool {
	return op.Security != nil && len(*op.Security) == 0
}
