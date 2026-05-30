//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06WalkRequestBody — requestBody 의 모든 media-type 스키마를 walk 에 추가

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkRequestBody appends every media-type schema declared by a requestBody
// to acc. nil bodies and nil values are skipped. $ref-only media schemas resolve
// to their components value (already enumerated elsewhere and deduplicated via
// visited).
func o06WalkRequestBody(body *openapi3.RequestBodyRef, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if body == nil || body.Value == nil {
		return acc
	}
	for _, mt := range body.Value.Content {
		acc = o06WalkMediaType(mt, visited, acc)
	}
	return acc
}
