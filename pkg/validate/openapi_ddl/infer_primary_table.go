//ff:func feature=validate type=util control=iteration dimension=2 topic=openapi-ddl
//ff:what inferPrimaryTable — operationId/path에서 주 DDL 테이블 추론

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// inferPrimaryTable guesses the primary DDL table for an operation by
// stripping CRUD prefixes from operationId and pluralising. Falls back to the
// last meaningful path segment. Returns "" when nothing matches DDL.
func inferPrimaryTable(fs *yongol.Fullstack, op *openapi3.Operation, path string) string {
	tables := make(map[string]bool, len(fs.DDLTables))
	for _, t := range fs.DDLTables {
		tables[t.Name] = true
	}
	if op != nil && op.OperationID != "" {
		stripped := stripCRUDPrefix(op.OperationID)
		if stripped != "" {
			cand := modelToTable(stripped)
			if tables[cand] {
				return cand
			}
		}
	}
	for _, seg := range pathSegments(path) {
		cand := modelToTable(seg)
		if tables[cand] {
			return cand
		}
		if tables[seg] {
			return seg
		}
	}
	return ""
}
