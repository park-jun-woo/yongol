//ff:func feature=generate type=util control=sequence
//ff:what OpenAPI doc에서 operationId로 requestBody 스키마를 찾아 기본 FieldConstraint를 생성한다
package generate

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// resolveFieldsFromOpenAPI looks up the operation in the OpenAPI doc and builds
// default FieldConstraint entries from its requestBody schema. Returns nil when
// the operation or schema is not found.
func resolveFieldsFromOpenAPI(doc *openapi3.T, opID string) map[string]oapiparser.FieldConstraint {
	op, found := findOpenAPIOpByID(doc, opID)
	if !found {
		return nil
	}
	schema := extractRequestBodySchema(op)
	if schema == nil {
		return nil
	}
	return buildDefaultFieldConstraints(schema)
}
