//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-70 — OpenAPI maxLength 가 DDL VARCHAR(n) 보다 크면 WARNING

package openapi_ddl

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// xdo70MaxLengthExceedsVarchar validates XDO-70: maxLength must not exceed the
// underlying DDL VARCHAR length.
func xdo70MaxLengthExceedsVarchar(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for opID, fields := range fs.RequestConstraints {
		for fieldName, fc := range fields {
			if fc.MaxLength == nil {
				continue
			}
			col := caseconv.PascalToSnake(fieldName)
			varcharLen, _, found := findDDLColumnConstraints(fs, col)
			if !found || varcharLen == 0 {
				continue
			}
			if *fc.MaxLength <= varcharLen {
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
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XDO-70] %s.%s — OpenAPI maxLength %d > DDL VARCHAR(%d) for %s", opID, fieldName, *fc.MaxLength, varcharLen, col),
				Advice:  fmt.Sprintf("OpenAPI maxLength 를 %d 이하로 줄이거나 DDL VARCHAR 길이를 늘리세요", varcharLen),
			})
		}
	}
	return diags
}
