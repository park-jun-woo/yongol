//ff:func feature=validate type=util control=sequence topic=openapi-ddl
//ff:what xdo77PropDiag — compares the OpenAPI↔DDL type of a single property and emits a diagnostic on mismatch

package openapi_ddl

import (
	"fmt"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo77PropDiag compares a single OpenAPI property type/format against the
// corresponding DDL Go type and returns (diag, true) when they are
// incompatible per `ddlGoTypeToOpenAPI`.
func xdo77PropDiag(fs *yongol.Fullstack, schemaName, tableName, propName string, propRef *openapi3.SchemaRef, cols map[string]string) (diagnostic.Diagnostic, bool) {
	if propRef == nil || propRef.Value == nil {
		return diagnostic.Diagnostic{}, false
	}
	colName := propName
	ddlGoType, colExists := cols[colName]
	if !colExists {
		// Column not in DDL — handled by XDO-9 (ghost property).
		return diagnostic.Diagnostic{}, false
	}
	compat, known := ddlGoTypeToOpenAPI[ddlGoType]
	if !known {
		return diagnostic.Diagnostic{}, false
	}
	oaType := propRef.Value.Type.Slice()
	if len(oaType) == 0 {
		return diagnostic.Diagnostic{}, false
	}
	actualType := oaType[0]
	actualFormat := propRef.Value.Format
	typeMismatch := actualType != compat.oType
	formatMismatch := compat.oFormat != "" && actualFormat != compat.oFormat
	if !typeMismatch && !formatMismatch {
		return diagnostic.Diagnostic{}, false
	}
	oaDisplay := actualType
	if actualFormat != "" {
		oaDisplay = actualType + "/" + actualFormat
	}
	ddlDisplay := compat.oType
	if compat.oFormat != "" {
		ddlDisplay = compat.oType + "/" + compat.oFormat
	}
	line := fs.OpenAPILines.SchemaPropertyLine(schemaName, propName)
	if line == 0 {
		line = fs.OpenAPILines.SchemaLine(schemaName)
	}
	return diagnostic.Diagnostic{
		File:  "api/openapi.yaml",
		Line:  line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XDO-77] schema %s field %s — OpenAPI type %q ↔ DDL column %s.%s type %q mismatch",
			schemaName, propName,
			oaDisplay,
			tableName, colName,
			ddlDisplay,
		),
		Advice: "Align the OpenAPI type to match the DDL column type",
	}, true
}
