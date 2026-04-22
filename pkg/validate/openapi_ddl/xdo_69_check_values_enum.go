//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-69 — DDL CHECK IN 값 집합과 OpenAPI enum 값 집합이 서로 다르면 ERROR

package openapi_ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
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
				Advice:  "OpenAPI enum 값을 DDL CHECK IN 값과 동일하게 맞추세요",
			})
		}
	}
	return diags
}
