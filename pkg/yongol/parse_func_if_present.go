//ff:func feature=orchestrator type=loader control=sequence
//ff:what Func 스펙 탐지 시 ParseDir 실행 — 성공 시 ProjectFuncSpecs 설정
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// parseFuncIfPresent parses project Func specs when KindFunc is present.
func parseFuncIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindFunc]
	if !ok {
		return
	}
	specs, diags := funcspec.ParseDir(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.ProjectFuncSpecs = specs
	}
}
