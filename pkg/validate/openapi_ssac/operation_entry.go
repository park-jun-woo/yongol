//ff:type feature=validate type=model topic=ssac-openapi
//ff:what OperationEntry — OpenAPI operationId → (HTTP method, Operation) 매핑 엔트리

package openapi_ssac

import "github.com/getkin/kin-openapi/openapi3"

// OperationEntry pairs an OpenAPI operation with the HTTP method it is
// served under. Required by XOS-80/82 which need the method to choose
// the conventional 2xx status (see openapi.DeriveSuccessStatus).
type OperationEntry struct {
	Method string
	Op     *openapi3.Operation
}
