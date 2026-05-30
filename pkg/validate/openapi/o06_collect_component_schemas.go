//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06CollectComponentSchemas — components.schemas 항목을 이름 순으로 walk 에 추가

package openapi

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// o06CollectComponentSchemas appends every schema reachable from
// doc.Components.Schemas to acc, iterating in deterministic name order so
// diagnostics are stable. Each component entry carries its schema name for
// precise LineIndex lookups.
func o06CollectComponentSchemas(doc *openapi3.T, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if doc.Components == nil || doc.Components.Schemas == nil {
		return acc
	}
	names := make([]string, 0, len(doc.Components.Schemas))
	for name := range doc.Components.Schemas {
		names = append(names, name)
	}
	sort.Strings(names)
	for _, name := range names {
		acc = o06WalkSchemaRef(doc.Components.Schemas[name], name, visited, acc)
	}
	return acc
}
