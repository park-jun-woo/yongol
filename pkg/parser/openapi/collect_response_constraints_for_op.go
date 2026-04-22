//ff:func feature=manifest type=parser control=sequence
//ff:what collectResponseConstraintsForOp — 단일 operation 의 response 제약조건을 수집해 result 에 반영
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectResponseConstraintsForOp extracts response field constraints for a
// single operation and merges them into result keyed by operationId.
func collectResponseConstraintsForOp(result map[string]map[string]FieldConstraint, op *openapi3.Operation, lines *LineIndex) {
	if op.OperationID == "" || op.Responses == nil {
		return
	}
	fields := extractResponseFields(op)
	if len(fields) == 0 {
		return
	}
	if lines != nil {
		fillResponseFieldLines(fields, op.OperationID, lines)
	}
	result[op.OperationID] = fields
}
