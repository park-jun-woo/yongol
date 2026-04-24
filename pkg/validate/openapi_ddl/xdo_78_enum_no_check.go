//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-ddl
//ff:what XDO-78 — ERROR when an OpenAPI enum request field has no matching DDL CHECK IN constraint

package openapi_ddl

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdo78EnumNoCheck validates XDO-78: reverse of XDO-68. Whenever an OpenAPI
// request property declares an `enum:`, the backing DDL column must carry a
// matching `CHECK (<col> IN (...))` constraint so the value set is enforced
// at the database layer too. Asymmetric declarations let an arbitrary
// string slip past OpenAPI (bypass via non-HTTP ingress, or a later SSaC
// @put that reuses the column) and land in the column.
func xdo78EnumNoCheck(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for opID, fields := range fs.RequestConstraints {
		for fieldName, fc := range fields {
			if len(fc.Enum) == 0 {
				continue
			}
			col := caseconv.PascalToSnake(fieldName)
			_, checkEnums, found := findDDLColumnConstraints(fs, col)
			if !found {
				continue
			}
			if len(checkEnums) > 0 {
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
				Message: fmt.Sprintf("[XDO-78] %s.%s — OpenAPI has enum [%s] but DDL column %s has no CHECK IN constraint", opID, fieldName, strings.Join(fc.Enum, ", "), col),
				Advice:  fmt.Sprintf("Add CHECK (%s IN ('%s')) to the DDL column so the value set is enforced at both layers", col, strings.Join(fc.Enum, "', '")),
			})
		}
	}
	return diags
}
