//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what o06WalkSchemaChildren — 스키마의 array items 와 additionalProperties 스키마로 재귀

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkSchemaChildren recurses into the non-property nested schemas of s: the
// array items schema and the additionalProperties schema. Property recursion is
// handled by the caller (o06WalkSchemaRef). nil children are no-ops inside
// o06WalkSchemaRef, so no explicit nil guard is needed for items.
func o06WalkSchemaChildren(s *openapi3.Schema, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	acc = o06WalkSchemaRef(s.Items, "", visited, acc)
	acc = o06WalkSchemaRef(s.AdditionalProperties.Schema, "", visited, acc)
	return acc
}
