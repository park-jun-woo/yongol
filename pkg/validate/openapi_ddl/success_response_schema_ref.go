//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what successResponseSchemaRef — operation 의 관례적 2xx JSON 응답 schema ref 반환

package openapi_ddl

import (
	"strconv"

	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// successResponseSchemaRef returns the application/json schema of op's
// conventional 2xx response (via DeriveSuccessStatus). Returns nil when the
// operation has no JSON-bodied 2xx response (e.g. 204) so callers skip it.
func successResponseSchemaRef(op *openapi3.Operation, method string) *openapi3.SchemaRef {
	if op == nil || op.Responses == nil {
		return nil
	}
	status := oapiparser.DeriveSuccessStatus(op, method)
	if status == 0 {
		return nil
	}
	resp := op.Responses.Map()[strconv.Itoa(status)]
	if resp == nil || resp.Value == nil || resp.Value.Content == nil {
		return nil
	}
	ct := resp.Value.Content.Get("application/json")
	if ct == nil || ct.Schema == nil || ct.Schema.Value == nil {
		return nil
	}
	return ct.Schema
}
