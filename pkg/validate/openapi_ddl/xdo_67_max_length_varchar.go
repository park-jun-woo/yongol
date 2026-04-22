//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-67 — DDL VARCHAR(n) 컬럼인데 OpenAPI 요청 필드가 maxLength를 지정하지 않으면 ERROR

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
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
				Advice:  fmt.Sprintf("OpenAPI 필드 %q 에 maxLength: %d 추가하세요 (DDL VARCHAR 길이와 일치)", fieldName, varcharLen),
			})
		}
	}
	return diags
}
