//ff:func feature=manifest type=parser control=sequence
//ff:what collectRequestConstraintsForOp — 단일 operation 의 requestBody 제약조건을 수집해 result 에 반영
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectRequestConstraintsForOp extracts requestBody field constraints for a
// single operation and merges them into result keyed by operationId.
func collectRequestConstraintsForOp(result map[string]map[string]FieldConstraint, op *openapi3.Operation, lines *LineIndex) {
	if op.OperationID == "" || op.RequestBody == nil {
		return
	}
	fields := extractBodyConstraints(op.RequestBody, op.OperationID)
	if len(fields) == 0 {
		return
	}
	if lines != nil {
		fillRequestFieldLines(fields, op.OperationID, lines)
	}
	result[op.OperationID] = fields
}
