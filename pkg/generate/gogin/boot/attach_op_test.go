//ff:func feature=gen-gogin type=test-helper control=selection
//ff:what attachOp — opSpec 하나를 OpenAPI 문서의 pathItem에 method별로 연결
package boot

import "github.com/getkin/kin-openapi/openapi3"

// attachOp finds or creates the PathItem for o.path and assigns the
// operation to the correct verb slot based on o.method. Only GET/POST/PUT/
// DELETE are recognised — other verbs are silently ignored, matching the
// test fixture's scope.
func attachOp(doc *openapi3.T, o opSpec) {
	pi := doc.Paths.Find(o.path)
	if pi == nil {
		pi = &openapi3.PathItem{}
		doc.Paths.Set(o.path, pi)
	}
	operation := &openapi3.Operation{OperationID: o.opID, Security: o.sec}
	switch o.method {
	case "GET":
		pi.Get = operation
	case "POST":
		pi.Post = operation
	case "PUT":
		pi.Put = operation
	case "DELETE":
		pi.Delete = operation
	}
}
