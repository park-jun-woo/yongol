//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectResponseSchemas — OpenAPI 200 응답에서 참조된 $ref schema 이름 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// collectResponseSchemas finds all $ref schema names used in 200 responses.
func collectResponseSchemas(doc *openapi3.T) map[string]bool {
	result := make(map[string]bool)
	if doc.Paths == nil {
		return result
	}
	for _, pathItem := range doc.Paths.Map() {
		collectFromPathItem(pathItem, result)
	}
	return result
}
