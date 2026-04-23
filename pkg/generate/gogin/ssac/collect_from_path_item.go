//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFromPathItem — 하나의 PathItem에서 모든 verb의 200 응답 $ref 스키마 이름 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// collectFromPathItem scans every verb on one PathItem and records the
// $ref schema names referenced by the success 2xx response body (both
// direct body refs and property-level refs). Each verb is paired with
// its HTTP method so collectFrom200Response can pick the correct 2xx
// status via DeriveSuccessStatus — POST→201, DELETE→204, etc.
func collectFromPathItem(pathItem *openapi3.PathItem, out map[string]bool) {
	verbs := []struct {
		method string
		op     *openapi3.Operation
	}{
		{"GET", pathItem.Get},
		{"POST", pathItem.Post},
		{"PUT", pathItem.Put},
		{"DELETE", pathItem.Delete},
		{"PATCH", pathItem.Patch},
	}
	for _, v := range verbs {
		collectFrom200Response(v.op, v.method, out)
	}
}
