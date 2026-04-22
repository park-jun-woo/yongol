//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ExtractRequestConstraints — extracts requestBody field constraints per operationId
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractRequestConstraints returns field constraints for the request body of
// each operationId. When lines is non-nil, each FieldConstraint.Line is set to
// the line number where the requestBody property is declared.
func ExtractRequestConstraints(doc *openapi3.T, lines *LineIndex) map[string]map[string]FieldConstraint {
	result := make(map[string]map[string]FieldConstraint)
	if doc == nil || doc.Paths == nil {
		return result
	}
	for _, item := range doc.Paths.Map() {
		extractRequestConstraintsOps(result, item, lines)
	}
	return result
}
