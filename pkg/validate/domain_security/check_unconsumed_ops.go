//ff:func feature=validate type=util control=iteration dimension=2 topic=domain-security
//ff:what checkUnconsumedOps — OpenAPI document에서 미소비 operationId 진단 생성
package domain_security

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// checkUnconsumedOps checks all operations in a document against consumed set,
// returning diagnostics for unconsumed ones. Auth endpoints (empty security) are excluded.
func checkUnconsumedOps(dd domainDoc, consumed map[string]struct{}, rulePrefix, domainLabel string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, item := range dd.Doc.Paths.Map() {
		ops := []*openapi3.Operation{
			item.Get, item.Post, item.Put, item.Delete, item.Patch,
		}
		for _, op := range ops {
			if op == nil || op.OperationID == "" || hasEmptySecurity(op) {
				continue
			}
			if _, ok := consumed[op.OperationID]; !ok {
				diags = append(diags, diagnostic.Diagnostic{
					File:    dd.Cfg.OpenAPI,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[%s] %s operationId %q is not consumed by any STML page in the %s frontend", rulePrefix, domainLabel, op.OperationID, domainLabel),
					Advice:  fmt.Sprintf("Add a data-fetch or data-action referencing %q in a %s STML page, or remove the endpoint from the %s OpenAPI", op.OperationID, domainLabel, domainLabel),
				})
			}
		}
	}
	return diags
}
