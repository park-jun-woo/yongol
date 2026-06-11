//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what requestBodyAllRequired — operation의 requestBody top-level 필드가 전부 required인지 판정

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// requestBodyAllRequired reports whether every top-level property of an
// operation's requestBody schema is listed as required (allOf members resolved
// one level, mirroring requestBodyFields). Returns false when there is no
// requestBody or it has no properties — there is nothing to relax.
func requestBodyAllRequired(op *openapi3.Operation) bool {
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return false
	}
	for _, mt := range op.RequestBody.Value.Content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		return allPropsRequired(mt.Schema.Value)
	}
	return false
}
