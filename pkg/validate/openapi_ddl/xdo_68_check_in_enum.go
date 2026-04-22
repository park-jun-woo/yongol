//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-68 — DDL CHECK IN(...) 컬럼인데 OpenAPI 요청 필드에 enum이 없으면 ERROR

package openapi_ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// xdo68CheckInEnum validates XDO-68: request fields backed by a CHECK IN
// constrained column must declare an OpenAPI enum.
func xdo68CheckInEnum(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for opID, fields := range fs.RequestConstraints {
		for fieldName, fc := range fields {
			col := caseconv.PascalToSnake(fieldName)
			_, checkEnums, found := findDDLColumnConstraints(fs, col)
			if !found || len(checkEnums) == 0 {
				continue
			}
			if len(fc.Enum) > 0 {
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
				Message: fmt.Sprintf("[XDO-68] %s.%s — DDL column %s has CHECK IN (%s) but OpenAPI has no enum", opID, fieldName, col, strings.Join(checkEnums, ", ")),
				Advice:  fmt.Sprintf("OpenAPI 에 enum: [%s] 를 추가해 DDL CHECK IN 값과 일치시키세요", strings.Join(checkEnums, ", ")),
			})
		}
	}
	return diags
}
