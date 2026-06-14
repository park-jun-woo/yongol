//ff:type feature=validate type=model topic=openapi-ddl
//ff:what canonOp — canonical 응답 테스트용 operation 표현 (method/opID/Responses)

package openapi_ddl

import "github.com/getkin/kin-openapi/openapi3"

// canonOp describes one operation for buildCanonicalFS: its HTTP method,
// operationId and the 2xx responses object.
type canonOp struct {
	method string
	opID   string
	resp   *openapi3.Responses
}
