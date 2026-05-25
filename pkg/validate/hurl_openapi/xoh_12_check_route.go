//ff:func feature=validate type=rule control=sequence topic=hurl-openapi
//ff:what xoh12CheckRoute — 단일 route 의 미커버 status code 진단 생성

package hurl_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func xoh12CheckRoute(route apiRoute, covered map[string]map[string]bool) []diagnostic.Diagnostic {
	if route.Op == nil || route.Op.OperationID == "" {
		return nil
	}
	opID := route.Op.OperationID

	declared := collectNon5xxCodes(route.Responses)
	if len(declared) == 0 {
		return nil
	}

	coveredSet := covered[opID]
	missing := filterUncovered(declared, coveredSet)
	if len(missing) == 0 {
		return nil
	}

	coveredStr := formatCoveredCodes(coveredSet)

	return []diagnostic.Diagnostic{{
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: fmt.Sprintf("[XOH-12] %s — OpenAPI declares [%s] but hurl only covers [%s]", opID, strings.Join(missing, ", "), coveredStr),
		Advice:  "Add scenario/invariant hurl tests for the uncovered status codes",
	}}
}
