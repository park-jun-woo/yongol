//ff:func feature=validate type=rule control=sequence topic=ddl-rego
//ff:what Run — DDL↔Rego 교차 검증 실행 (XDP-*)
package ddl_rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all DDL↔Rego cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xdp31OwnershipTable(fs)...)
	diags = append(diags, xdp32OwnershipColumn(fs)...)
	diags = append(diags, xdp33OwnershipJoinTable(fs)...)
	diags = append(diags, xdp34OwnershipJoinColumn(fs)...)
	diags = append(diags, xdp65RoleDDLCheck(fs)...)
	return diags
}
