//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what topLevelKeys — 응답 schema 의 top-level property 키를 정렬해 반환

package openapi_ddl

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// topLevelKeys returns the sorted top-level property names of a response schema.
// For a $ref the kin-openapi loader populates Value with the resolved schema, so
// this yields the component's own top-level fields.
func topLevelKeys(schemaRef *openapi3.SchemaRef) []string {
	if schemaRef == nil || schemaRef.Value == nil {
		return nil
	}
	keys := make([]string, 0, len(schemaRef.Value.Properties))
	for k := range schemaRef.Value.Properties {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
