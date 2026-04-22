//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.tryExtractFromPathItem — PathItem의 verb 중 매칭되는 op 찾아 메타데이터 적재

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// tryExtractFromPathItem scans each verb on pathItem. When it finds the
// operation whose OperationID matches, it populates g's metadata maps from
// that op plus the path-level parameters and returns true. Returns false
// when no verb matches — extractFromOpenAPI keeps scanning.
func (g *methodGen) tryExtractFromPathItem(pathItem *openapi3.PathItem, operationId string) bool {
	ops := []*openapi3.Operation{
		pathItem.Get, pathItem.Post, pathItem.Put, pathItem.Delete, pathItem.Patch,
	}
	for _, op := range ops {
		if op == nil || op.OperationID != operationId {
			continue
		}
		// path-level + operation-level parameters
		for _, p := range pathItem.Parameters {
			g.addParam(p, operationId)
		}
		for _, p := range op.Parameters {
			g.addParam(p, operationId)
		}
		// request body schema (for format-based casts)
		g.extractBodyFormats(op)
		// 200 response schema
		g.extractRespFields(op)
		return true
	}
	return false
}
