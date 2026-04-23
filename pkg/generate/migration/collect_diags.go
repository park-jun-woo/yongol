//ff:func feature=migration type=util control=sequence
//ff:what collectDiags — Generate 파이프라인 전체 진단 (errDiags + MIG-001 + CheckSafety + 누락) 병합
package migration

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// collectDiags concatenates previously-gathered error diagnostics with
// MIG-001 rename mismatches and MIG-002/004/005 safety issues.
func collectDiags(errDiags []diagnostic.Diagnostic, prev, curr *Schema, hints *Hints, issues []SafetyIssue, missing []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, errDiags...)
	diags = append(diags, mig001From(prev, curr, hints)...)
	diags = append(diags, issuesToDiags(issues, missing)...)
	return diags
}
