//ff:func feature=orchestrator type=loader control=sequence
//ff:what STML 탐지 시 ParseDir 실행 — 진단은 수집, 성공 시 STMLPages 설정
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// parseSTMLIfPresent parses the STML directory when the KindSTML SSOT is present.
func parseSTMLIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindSTML]
	if !ok {
		return
	}
	pages, diags := stml.ParseDir(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.STMLPages = pages
	}
}
