//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06WalkResponse — 단일 response 의 모든 media-type 스키마를 walk 에 추가

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkResponse appends every media-type schema declared by a single response
// to acc. nil refs and nil values are skipped.
func o06WalkResponse(resp *openapi3.ResponseRef, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if resp == nil || resp.Value == nil {
		return acc
	}
	for _, mt := range resp.Value.Content {
		acc = o06WalkMediaType(mt, visited, acc)
	}
	return acc
}
