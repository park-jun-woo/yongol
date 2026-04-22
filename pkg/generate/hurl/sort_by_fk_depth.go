//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what sortByFKDepth — DDL FK 의존 깊이 순으로 Create 경로 정렬 (부모 먼저)
package hurl

import (
	"sort"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
)

// sortByFKDepth returns OpenAPI paths sorted by FK dependency depth (parents first).
func sortByFKDepth(paths *openapi3.Paths, tables []ddl.Table) []sortedPath {
	if paths == nil {
		return nil
	}
	depthMap := buildFKDepthMap(tables)
	var result []sortedPath
	for path := range paths.Map() {
		result = append(result, sortedPath{Path: path, Depth: depthMap[resourceFromPath(path)]})
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Depth != result[j].Depth {
			return result[i].Depth < result[j].Depth
		}
		return result[i].Path < result[j].Path
	})
	return result
}
