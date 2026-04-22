//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what collectDeleteOps — OpenAPI에서 DELETE operation 목록 수집
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectDeleteOps collects all DELETE operations from OpenAPI.
func collectDeleteOps(fs *yongol.Fullstack) []deleteOp {
	var ops []deleteOp
	for path, pathItem := range fs.OpenAPIDoc.Paths.Map() {
		if pathItem.Delete == nil || pathItem.Delete.OperationID == "" {
			continue
		}
		if isAuthOpID(pathItem.Delete.OperationID) {
			continue
		}
		ops = append(ops, deleteOp{path: path, opID: pathItem.Delete.OperationID, resource: resourceFromPath(path)})
	}
	return ops
}
