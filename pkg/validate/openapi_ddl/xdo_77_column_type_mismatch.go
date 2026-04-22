//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-77 — DDL column 타입 ↔ OpenAPI field 타입 불일치 → ERROR

package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// ddlGoTypeToOpenAPI maps DDL-derived Go type strings to their required OpenAPI type+format.
var ddlGoTypeToOpenAPI = map[string]openAPITypeCompat{
	"int64":           {oType: "integer", oFormat: "int64"},
	"int32":           {oType: "integer", oFormat: "int32"},
	"int":             {oType: "integer", oFormat: ""},
	"string":          {oType: "string", oFormat: ""},
	"bool":            {oType: "boolean", oFormat: ""},
	"time.Time":       {oType: "string", oFormat: "date-time"},
	"float64":         {oType: "number", oFormat: ""},
	"float32":         {oType: "number", oFormat: ""},
	"json.RawMessage": {oType: "object", oFormat: ""},
}

// xdo77ColumnTypeMismatch validates XDO-77: for every components/schemas property
// that corresponds to a DDL column, the OpenAPI type/format must match the DDL
// Go type according to the compatibility table above.
func xdo77ColumnTypeMismatch(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Components == nil || fs.OpenAPIDoc.Components.Schemas == nil {
		return nil
	}

	tableIndex := xdo77BuildTableIndex(fs)

	var diags []diagnostic.Diagnostic
	for schemaName, schemaRef := range fs.OpenAPIDoc.Components.Schemas {
		diags = append(diags, xdo77ScanSchema(fs, schemaName, schemaRef, tableIndex)...)
	}
	return diags
}
