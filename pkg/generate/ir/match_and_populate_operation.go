//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what matchAndPopulateOperation -- pathItem 의 method 별 operation 중 operationID 매칭 시 plan 채우고 true 반환

package ir

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// matchAndPopulateOperation scans the HTTP methods of pathItem for an operation
// whose OperationID equals operationID. On a match it fills plan (method, path,
// status, params, body) and returns true. Returns false when no method matches.
func matchAndPopulateOperation(operationID, path string, pathItem *openapi3.PathItem, plan *ServicePlan, pathParams, queryParams map[string]bool) bool {
	ops := map[string]*openapi3.Operation{
		"GET": pathItem.Get, "POST": pathItem.Post,
		"PUT": pathItem.Put, "DELETE": pathItem.Delete,
		"PATCH": pathItem.Patch,
	}
	for method, op := range ops {
		if op == nil || op.OperationID != operationID {
			continue
		}
		plan.HTTPMethod = method
		plan.URLPath = path
		if s := oapiparser.DeriveSuccessStatus(op, method); s != 0 {
			plan.SuccessStatus = s
		}
		populateOperationParams(plan, pathItem, op, pathParams, queryParams)
		return true
	}
	return false
}
