//ff:func feature=validate type=rule control=sequence topic=features-ddl
//ff:what Run — Features↔DDL 교차 검증 실행 (XFD-*)
package features_ddl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Features↔DDL cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xfd01TableExists(fs)...)
	diags = append(diags, xfd02FKExists(fs)...)
	return diags
}
