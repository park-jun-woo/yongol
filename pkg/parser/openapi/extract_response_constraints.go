//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ExtractResponseConstraints — extracts response field constraints per operationId
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractResponseConstraints returns field constraints for the 2xx response of
// each operationId. When lines is non-nil, each FieldConstraint.Line is set to
// the line number where the response schema property is declared.
func ExtractResponseConstraints(doc *openapi3.T, lines *LineIndex) map[string]map[string]FieldConstraint {
	result := make(map[string]map[string]FieldConstraint)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		extractResponseConstraintsOps(result, item, lines)
	}
	return result
}
