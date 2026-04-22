//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFromPathItem — 하나의 PathItem에서 모든 verb의 200 응답 $ref 스키마 이름 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// collectFromPathItem scans every verb on one PathItem and records the
// $ref schema names referenced by properties of its 200 response body.
// Extracted from collectResponseSchemas to keep that loop at depth 1.
func collectFromPathItem(pathItem *openapi3.PathItem, out map[string]bool) {
	ops := []*openapi3.Operation{
		pathItem.Get, pathItem.Post, pathItem.Put, pathItem.Delete, pathItem.Patch,
	}
	for _, op := range ops {
		collectFrom200Response(op, out)
	}
}
