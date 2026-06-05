//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what xmo12ItemDiags — PathItem 내 no-front인데 소비된 operation마다 XMO-12 WARNING 생성

package stml_openapi

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// xmo12ItemDiags emits a WARNING for each no-front operation in item that is
// nonetheless present in the consumed set.
func xmo12ItemDiags(item *openapi3.PathItem, consumed map[string]struct{}) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, op := range operationsOf(item) {
		if op == nil || op.OperationID == "" || !isNoFront(op) {
			continue
		}
		if _, ok := consumed[op.OperationID]; !ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "openapi.yaml",
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[XMO-12] operationId %q is tagged \"no-front\" but is actually consumed by an STML page or component", op.OperationID),
			Advice:  fmt.Sprintf("Remove the \"no-front\" tag from %q if the frontend uses it, or stop consuming it if it is meant to be backend-only", op.OperationID),
		})
	}
	return diags
}
