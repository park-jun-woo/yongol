//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what ExtractRequestConstraints — operationId별 requestBody 필드 제약조건 추출
package openapi

import "github.com/getkin/kin-openapi/openapi3"

// ExtractRequestConstraints returns field constraints for the request body of
// each operationId. lines 가 nil 이 아니면 각 FieldConstraint 의 Line 을
// requestBody 의 property 가 선언된 줄 번호로 채운다.
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
