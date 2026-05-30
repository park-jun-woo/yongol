//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06CollectPathSchemas — 모든 path/operation 의 requestBody·response inline 스키마를 walk 에 추가

package openapi

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
)

// o06CollectPathSchemas appends every inline schema reachable from each
// operation's requestBody and responses to acc, iterating paths in sorted order
// for stable diagnostics. Per-operation traversal is delegated to
// o06WalkOperation to keep loop nesting within filefunc's depth limit.
func o06CollectPathSchemas(doc *openapi3.T, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if doc.Paths == nil {
		return acc
	}
	paths := doc.Paths.Map()
	keys := make([]string, 0, len(paths))
	for p := range paths {
		keys = append(keys, p)
	}
	sort.Strings(keys)
	for _, p := range keys {
		acc = o06CollectItemSchemas(paths[p], visited, acc)
	}
	return acc
}
