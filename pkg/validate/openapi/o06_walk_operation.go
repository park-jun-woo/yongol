//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what o06WalkOperation — 단일 operation 의 requestBody/response inline 스키마를 walk 에 추가

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkOperation appends every inline schema reachable from op's requestBody
// and responses to acc, deduplicating via visited. Request and response
// traversal are delegated to keep loop nesting within filefunc's depth limit.
func o06WalkOperation(op *openapi3.Operation, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if op == nil {
		return acc
	}
	acc = o06WalkRequestBody(op.RequestBody, visited, acc)
	acc = o06WalkResponses(op.Responses, visited, acc)
	return acc
}
