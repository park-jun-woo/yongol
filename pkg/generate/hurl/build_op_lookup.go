//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what buildOpLookup — OpenAPI operationID에서 method+path 룩업 맵 구축
package hurl

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildOpLookup builds a map from operationID to method+path.
func buildOpLookup(fs *yongol.Fullstack) map[string]opInfo {
	lookup := map[string]opInfo{}
	for path, pathItem := range fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pathItem.Operations() {
			if op.OperationID != "" {
				lookup[op.OperationID] = opInfo{Method: strings.ToUpper(method), Path: path}
			}
		}
	}
	return lookup
}
