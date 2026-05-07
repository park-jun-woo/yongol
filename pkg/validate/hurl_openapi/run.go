//ff:func feature=validate type=rule control=sequence topic=hurl-openapi
//ff:what Run — Hurl↔OpenAPI 교차 검증 실행 (XOH-01/02/03/04/08/09/10/11)
package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes every Hurl↔OpenAPI cross-validation rule. Each rule is
// skipped cheaply when the prerequisite SSOT is missing; callers rely on
// the step-kind gating in pkg/validate/all_steps.go to avoid running
// this at all when either Hurl or OpenAPI is absent.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xoh01URLMethod(fs)...)
	diags = append(diags, xoh02StatusDeclared(fs)...)
	diags = append(diags, xoh03RequestFieldInSchema(fs)...)
	diags = append(diags, xoh04AssertPathInSchema(fs)...)
	diags = append(diags, xoh08CapturePathInSchema(fs)...)
	diags = append(diags, xoh09UnusedCapture(fs)...)
	xoh10 := xoh10SmokeRequired(fs)
	diags = append(diags, xoh10...)
	if len(xoh10) == 0 {
		diags = append(diags, xoh11SmokeCoverage(fs)...)
	}
	return diags
}
