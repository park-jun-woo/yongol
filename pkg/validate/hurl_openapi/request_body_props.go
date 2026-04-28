//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what requestBodyProps — operation 의 첫 JSON requestBody 의 top-level property 추출

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// requestBodyProps returns the top-level properties of the first JSON
// content type declared on an operation's request body. Second return
// signals whether a usable schema exists.
func requestBodyProps(op *openapi3.Operation) (map[string]struct{}, bool) {
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil, false
	}
	for _, mt := range op.RequestBody.Value.Content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		return schemaPropertyNames(mt.Schema.Value), true
	}
	return nil, false
}
