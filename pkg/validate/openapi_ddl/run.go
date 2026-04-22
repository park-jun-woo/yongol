//ff:func feature=validate type=rule control=sequence topic=openapi-ddl
//ff:what Run — OpenAPI↔DDL 교차 검증 실행 (XDO-*, XOD-*)
package openapi_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all OpenAPI↔DDL cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xdo09GhostProperty(fs)...)
	diags = append(diags, xdo67MaxLengthVarchar(fs)...)
	diags = append(diags, xdo68CheckInEnum(fs)...)
	diags = append(diags, xdo69CheckValuesEnum(fs)...)
	diags = append(diags, xdo70MaxLengthExceedsVarchar(fs)...)
	diags = append(diags, xdo75OptionalNotNullNoDefault(fs)...)
	diags = append(diags, xdo76RequiredNullable(fs)...)
	diags = append(diags, xdo77ColumnTypeMismatch(fs)...)
	diags = append(diags, xod10DDLToResponse(fs)...)
	return diags
}
