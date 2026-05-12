//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what checkAdminPathSecurity — 단일 path의 operation들에서 public 접근 허용 여부 검사
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkAdminPathSecurity checks all operations on a single path for empty security.
func checkAdminPathSecurity(path string, item *openapi3.PathItem, openapiFile string) []diagnostic.Diagnostic {
	ops := []*openapi3.Operation{
		item.Get, item.Post, item.Put, item.Delete, item.Patch,
	}
	var diags []diagnostic.Diagnostic
	for _, op := range ops {
		if op == nil || op.OperationID == "" {
			continue
		}
		if hasEmptySecurity(op) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    openapiFile,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XDS-80] admin endpoint %q (%s) has security: [] allowing public access", op.OperationID, path),
				Advice:  "Remove the empty security override or add an appropriate security requirement to the admin endpoint",
			})
		}
	}
	return diags
}
