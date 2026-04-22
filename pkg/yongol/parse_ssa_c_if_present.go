//ff:func feature=orchestrator type=loader control=sequence
//ff:what SSaC 탐지 시 ParseDir 실행 — 진단은 수집, 성공 시 ServiceFuncs 설정
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// parseSSaCIfPresent parses the SSaC directory when the KindSSaC SSOT is present.
func parseSSaCIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindSSaC]
	if !ok {
		return
	}
	funcs, diags := ssac.ParseDir(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.ServiceFuncs = funcs
	}
}
