//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what getSuccessResponseSchema — operation의 2xx response JSON schema 추출
package hurl

import (
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// getSuccessResponseSchema extracts the JSON schema from the first 2xx response.
func getSuccessResponseSchema(op *openapi3.Operation) *openapi3.Schema {
	if op == nil || op.Responses == nil {
		return nil
	}
	for code, ref := range op.Responses.Map() {
		if !strings.HasPrefix(code, "2") || ref == nil || ref.Value == nil || ref.Value.Content == nil {
			continue
		}
		ct := ref.Value.Content.Get("application/json")
		if ct != nil && ct.Schema != nil && ct.Schema.Value != nil {
			return ct.Schema.Value
		}
	}
	return nil
}
