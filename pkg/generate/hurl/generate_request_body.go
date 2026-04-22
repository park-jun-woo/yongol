//ff:func feature=gen-hurl type=util control=sequence
//ff:what generateRequestBody — OpenAPI requestBody schema에서 dummy JSON body 생성
package hurl

import (
	"encoding/json"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// generateRequestBody builds a JSON request body from the operation's requestBody schema.
func generateRequestBody(op *openapi3.Operation, fs *yongol.Fullstack, role string) string {
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return ""
	}
	mt := op.RequestBody.Value.Content.Get("application/json")
	if mt == nil {
		mt = firstMediaType(op.RequestBody.Value.Content)
	}
	if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
		return ""
	}
	body := buildBodyFromSchema(mt.Schema.Value, fs, role)
	if body == nil {
		return ""
	}
	data, err := json.Marshal(body)
	if err != nil {
		return "{}"
	}
	return string(data)
}
