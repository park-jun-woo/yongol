//ff:func feature=validate type=rule control=sequence topic=ssac-ddl
//ff:what Run — SSaC↔DDL 교차 검증 실행 (XDS-*, XSD-*)
package ssac_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔DDL cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xds12ResultNoDDLTable(fs)...)
	diags = append(diags, xsd55DDLToModelRef(fs)...)
	return diags
}
