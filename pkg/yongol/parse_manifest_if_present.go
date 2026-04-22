//ff:func feature=orchestrator type=loader control=sequence
//ff:what manifest.yaml 탐지 시 로드 — 진단은 수집, 성공 시 Fullstack.Manifest 설정
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

// parseManifestIfPresent loads manifest.yaml when the KindConfig SSOT is present.
func parseManifestIfPresent(fs *Fullstack, root string, has map[SSOTKind]DetectedSSOT) {
	if _, ok := has[KindConfig]; !ok {
		return
	}
	cfg, diags := manifest.Load(root)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.Manifest = cfg
	}
}
