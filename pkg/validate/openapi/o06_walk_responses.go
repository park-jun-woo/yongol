//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06WalkResponses — 모든 response 의 media-type 스키마를 walk 에 추가

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkResponses appends every media-type schema declared by any response to
// acc. nil response sets and nil response values are skipped. Per-response
// media-type traversal is delegated to o06WalkResponse to keep loop nesting
// within filefunc's depth limit.
func o06WalkResponses(resps *openapi3.Responses, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if resps == nil {
		return acc
	}
	for _, resp := range resps.Map() {
		acc = o06WalkResponse(resp, visited, acc)
	}
	return acc
}
