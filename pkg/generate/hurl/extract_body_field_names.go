//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what extractBodyFieldNames — operation requestBody schema에서 필드명 목록 추출
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// extractBodyFieldNames extracts field names from the operation's requestBody schema.
func extractBodyFieldNames(op *openapi3.Operation) []string {
	if op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil
	}
	var names []string
	for _, mt := range op.RequestBody.Value.Content {
		if mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		for name := range mt.Schema.Value.Properties {
			names = append(names, name)
		}
	}
	return names
}
