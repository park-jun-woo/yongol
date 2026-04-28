//ff:func feature=validate type=test-helper control=iteration dimension=2 topic=hurl-openapi
//ff:what newDoc — 테스트용 최소 *openapi3.T 빌더 (paths→methods→operation 트리)

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// newDoc builds a minimal *openapi3.T whose paths match the given
// table. Operations are attached by HTTP method verb. Used by every
// hurl_openapi rule test to keep fixtures inline and readable.
func newDoc(paths map[string]map[string]*openapi3.Operation) *openapi3.T {
	doc := &openapi3.T{Paths: &openapi3.Paths{}}
	for p, methods := range paths {
		pi := &openapi3.PathItem{}
		for m, op := range methods {
			if op.Responses == nil {
				op.Responses = openapi3.NewResponses()
			}
			setOp(pi, m, op)
		}
		doc.Paths.Set(p, pi)
	}
	return doc
}
