//ff:type feature=validate type=model topic=stml-openapi
//ff:what operationEntry — OpenAPI operationId → (HTTP method, Operation) 매핑 엔트리

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// operationEntry pairs an OpenAPI operation with the HTTP method it is
// served under.
type operationEntry struct {
	method string
	op     *openapi3.Operation
}
