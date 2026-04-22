//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what inferCaptureField — response schema의 id 필드 탐색 → capture 생성 (snake 정규화)
package hurl

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"
)

// inferCaptureField finds id fields in the response schema and returns captures.
// Variable names go through snakeHurlName so they can be referenced later in
// paths that also use snake-normalized identifiers.
func inferCaptureField(op *openapi3.Operation, resource string) []capture {
	schema := getSuccessResponseSchema(op)
	if schema == nil {
		return nil
	}
	res := snakeHurlName(resource)
	if _, ok := schema.Properties["id"]; ok {
		return []capture{{VarName: res + "_id", JSONPath: `$.id`}}
	}
	for name, propRef := range schema.Properties {
		if propRef == nil || propRef.Value == nil {
			continue
		}
		if _, hasID := propRef.Value.Properties["id"]; hasID {
			return []capture{{VarName: snakeHurlName(name) + "_id", JSONPath: fmt.Sprintf("$.%s.id", name)}}
		}
	}
	return nil
}
