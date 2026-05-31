//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-67 — ERROR when a DDL VARCHAR(n) column's OpenAPI request field does not specify maxLength

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo67MaxLengthVarchar validates XDO-67: request fields backed by a VARCHAR
// column must declare maxLength.
func xdo67MaxLengthVarchar(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for opID, fields := range fs.RequestConstraints {
		for fieldName, fc := range fields {
			col := caseconv.PascalToSnake(fieldName)
			varcharLen, _, found := findDDLColumnConstraints(fs, col)
			if !found || varcharLen == 0 {
				continue
			}
			if fc.MaxLength != nil {
				continue
			}
			line := fs.OpenAPILines.RequestFieldLine(opID, fieldName)
			if line == 0 {
				line = fc.Line
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    "api/openapi.yaml",
				Line:    line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[XDO-67] %s.%s — DDL column %s is VARCHAR(%d) but OpenAPI has no maxLength", opID, fieldName, col, varcharLen),
				Advice:  fmt.Sprintf("Add maxLength: %d to OpenAPI field %q to match the DDL VARCHAR length", varcharLen, fieldName),
			})
		}
	}
	return diags
}
