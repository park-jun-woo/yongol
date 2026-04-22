//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-ddl
//ff:what xdo09ScanSchemaProps — 단일 schema 의 property 를 DDL 컬럼과 대조하여 ghost property 진단 생성

package openapi_ddl

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo09ScanSchemaProps inspects a single components/schemas entry and emits
// XDO-9 diagnostics for every property that has no matching DDL column.
func xdo09ScanSchemaProps(fs *yongol.Fullstack, schemaName string, schemaRef *openapi3.SchemaRef, tables map[string]map[string]bool) []diagnostic.Diagnostic {
	if schemaRef == nil || schemaRef.Value == nil {
		return nil
	}
	tableName := modelToTable(schemaName)
	cols, ok := tables[tableName]
	if !ok {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for propName := range schemaRef.Value.Properties {
		if cols[propName] {
			continue
		}
		line := fs.OpenAPILines.SchemaPropertyLine(schemaName, propName)
		if line == 0 {
			line = fs.OpenAPILines.SchemaLine(schemaName)
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "api/openapi.yaml",
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[XDO-9] schema %s property %q has no DDL column %s.%s", schemaName, propName, tableName, propName),
			Advice:  fmt.Sprintf("OpenAPI schema 에서 컬럼 %q 를 제거하거나 DDL %s 테이블에 추가하세요", propName, tableName),
		})
	}
	return diags
}
