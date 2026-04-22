//ff:func feature=validate type=rule control=sequence topic=rego-structural
//ff:what Run — Rego 검증 전체 실행 (P-*, XPP-*)
package rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Rego validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, p01Parse(fs)...)
	diags = append(diags, xpp30OwnershipNoAnnotation(fs)...)
	return diags
}
