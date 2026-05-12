//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what XMO-10 — OpenAPI operationId가 STML data-fetch/data-action에서 소비되지 않음 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// xmo10Unconsumed detects OpenAPI operationIds that are never referenced
// from any STML data-fetch or data-action block. Auth endpoints (those
// with an empty security requirement, i.e. security: []) are excluded.
func xmo10Unconsumed(pages []stml.PageSpec, doc *openapi3.T) []diagnostic.Diagnostic {
	if doc == nil || doc.Paths == nil || len(pages) == 0 {
		return nil
	}

	consumed := collectConsumedOps(pages)

	var diags []diagnostic.Diagnostic
	for _, item := range doc.Paths.Map() {
		for _, op := range []*openapi3.Operation{item.Get, item.Post, item.Put, item.Delete, item.Patch} {
			if op == nil || op.OperationID == "" || isAuthEndpoint(op) {
				continue
			}
			if _, ok := consumed[op.OperationID]; !ok {
				diags = append(diags, diagnostic.Diagnostic{
					File:    "openapi.yaml",
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelWarning,
					Message: fmt.Sprintf("[XMO-10] operationId %q is defined in OpenAPI but never consumed by any STML data-fetch or data-action", op.OperationID),
					Advice:  fmt.Sprintf("Either add a data-fetch or data-action referencing %q in an STML page, or remove the endpoint from OpenAPI if it is unused", op.OperationID),
				})
			}
		}
	}
	return diags
}
