//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what o06CollectAllSchemas — components + 모든 operation 스키마를 (중복 없이) 수집하는 진입점

package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o06CollectAllSchemas walks every schema reachable from components.schemas and
// from each operation's requestBody / response bodies, returning a flat list of
// unique schema values (deduplicated by pointer to avoid reporting the same
// dangling required twice). Nested schemas (properties, array items,
// additionalProperties) are followed so deeply nested required arrays are also
// covered. components entries carry their schema name for precise line lookup;
// inline schemas carry an empty name.
func o06CollectAllSchemas(fs *yongol.Fullstack) []o06SchemaEntry {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	visited := make(map[*openapi3.Schema]bool)
	acc := o06CollectComponentSchemas(fs.OpenAPIDoc, visited, nil)
	acc = o06CollectPathSchemas(fs.OpenAPIDoc, visited, acc)
	return acc
}
