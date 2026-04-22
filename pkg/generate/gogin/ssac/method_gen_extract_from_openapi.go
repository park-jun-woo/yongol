//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.extractFromOpenAPI — operationId에 맞는 op 탐색 후 메타데이터 적재

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// extractFromOpenAPI finds the operation by operationId and populates
// PathParams, QueryParams, BodyFormats, and RespFields with the metadata
// needed to generate correctly typed Go code against oapi-codegen output.
func (g *methodGen) extractFromOpenAPI(doc *openapi3.T, operationId string) {
	if doc.Paths == nil {
		return
	}
	for _, pathItem := range doc.Paths.Map() {
		if g.tryExtractFromPathItem(pathItem, operationId) {
			return
		}
	}
}
