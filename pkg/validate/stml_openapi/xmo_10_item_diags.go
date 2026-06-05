//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what xmo10ItemDiags — PathItem 내 미소비·non-no-front operation마다 XMO-10 ERROR 생성

package stml_openapi

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// xmo10ItemDiags emits an ERROR for each operation in item that is neither
// consumed nor marked no-front.
func xmo10ItemDiags(item *openapi3.PathItem, consumed map[string]struct{}) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, op := range operationsOf(item) {
		if op == nil || op.OperationID == "" || isNoFront(op) {
			continue
		}
		if _, ok := consumed[op.OperationID]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "openapi.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XMO-10] operationId %q is defined in OpenAPI but never consumed by any STML page or component", op.OperationID),
			Advice:  fmt.Sprintf("Consume %q from an STML data-fetch/data-action or component, mark it with tags: [\"no-front\"], or remove it from OpenAPI", op.OperationID),
		})
	}
	return diags
}
