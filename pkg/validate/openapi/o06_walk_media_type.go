//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what o06WalkMediaType — media-type 의 inline 스키마를 walk 에 추가

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkMediaType appends the inline schema of a single media-type entry to acc.
// nil media types are skipped. The empty schemaName marks the schema as inline
// (no LineIndex entry).
func o06WalkMediaType(mt *openapi3.MediaType, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if mt == nil {
		return acc
	}
	return o06WalkSchemaRef(mt.Schema, "", visited, acc)
}
