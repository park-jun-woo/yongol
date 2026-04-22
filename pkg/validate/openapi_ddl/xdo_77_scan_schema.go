//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what xdo77ScanSchema — 단일 schema 의 속성별로 DDL 타입 대조

package openapi_ddl

import (
	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo77ScanSchema walks every property of a single schema and emits XDO-77
// diagnostics when the OpenAPI type does not match the mapped DDL Go type.
func xdo77ScanSchema(fs *yongol.Fullstack, schemaName string, schemaRef *openapi3.SchemaRef, tableIndex map[string]map[string]string) []diagnostic.Diagnostic {
	if schemaRef == nil || schemaRef.Value == nil {
		return nil
	}
	tableName := modelToTable(schemaName)
	cols, ok := tableIndex[tableName]
	if !ok {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for propName, propRef := range schemaRef.Value.Properties {
		d, ok := xdo77PropDiag(fs, schemaName, tableName, propName, propRef, cols)
		if ok {
			diags = append(diags, d)
		}
	}
	return diags
}
