//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06WalkSchemaRef — 스키마 ref 를 visited 중복 없이 수집하고 nested 스키마로 재귀

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06WalkSchemaRef resolves ref to its inline value, records it once (keyed by
// pointer in visited), and recurses into nested schemas via
// o06WalkSchemaChildren. Pure $ref nodes (Value == nil) carry no own
// required/properties and are skipped; their target is enumerated separately via
// components.schemas. The visited set guarantees each schema is appended at most
// once even under cyclic or shared references.
func o06WalkSchemaRef(ref *openapi3.SchemaRef, schemaName string, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if ref == nil || ref.Value == nil {
		return acc
	}
	s := ref.Value
	if visited[s] {
		return acc
	}
	visited[s] = true
	acc = append(acc, o06SchemaEntry{schema: s, schemaName: schemaName})
	for _, propRef := range s.Properties {
		acc = o06WalkSchemaRef(propRef, "", visited, acc)
	}
	return o06WalkSchemaChildren(s, visited, acc)
}
