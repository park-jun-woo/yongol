//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06CollectItemSchemas — 단일 path item 의 모든 operation 스키마를 walk 에 추가

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// o06CollectItemSchemas appends every inline schema reachable from each
// operation of a single path item to acc. nil items are skipped. Operation-level
// traversal is delegated to o06WalkOperation.
func o06CollectItemSchemas(item *openapi3.PathItem, visited map[*openapi3.Schema]bool, acc []o06SchemaEntry) []o06SchemaEntry {
	if item == nil {
		return acc
	}
	for _, op := range item.Operations() {
		acc = o06WalkOperation(op, visited, acc)
	}
	return acc
}
