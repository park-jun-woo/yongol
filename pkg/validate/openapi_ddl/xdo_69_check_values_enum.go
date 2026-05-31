//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-69 — ERROR when the DDL CHECK IN value set differs from the OpenAPI enum value set

package openapi_ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo69CheckValuesEnum validates XDO-69: when both DDL CHECK IN and OpenAPI
// enum exist for a field, their value sets must match.
func xdo69CheckValuesEnum(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for opID, fields := range fs.RequestConstraints {
		for fieldName, fc := range fields {
			col := caseconv.PascalToSnake(fieldName)
			_, checkEnums, found := findDDLColumnConstraints(fs, col)
			if !found || len(checkEnums) == 0 || len(fc.Enum) == 0 {
				continue
			}
			if enumsMatch(checkEnums, fc.Enum) {
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
				Message: fmt.Sprintf("[XDO-69] %s.%s — DDL CHECK IN (%s) ≠ OpenAPI enum [%s]", opID, fieldName, strings.Join(checkEnums, ", "), strings.Join(fc.Enum, ", ")),
				Advice:  "Align the OpenAPI enum values with the DDL CHECK IN values",
			})
		}
	}
	return diags
}
