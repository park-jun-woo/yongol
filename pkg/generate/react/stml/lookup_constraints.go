//ff:func feature=stml-gen type=util control=sequence
//ff:what operationId로 OpenAPI 필드 제약조건을 조회한다
package stml

import oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"

// lookupConstraints retrieves the field constraints for the given operationId.
// Returns nil when no constraints exist.
func lookupConstraints(operationID string, constraints map[string]map[string]oapiparser.FieldConstraint) map[string]oapiparser.FieldConstraint {
	if constraints == nil {
		return nil
	}
	fields, ok := constraints[operationID]
	if !ok || len(fields) == 0 {
		return nil
	}
	return fields
}
