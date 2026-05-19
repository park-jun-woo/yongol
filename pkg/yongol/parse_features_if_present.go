//ff:func feature=orchestrator type=loader control=sequence
//ff:what features.yaml 탐지 시 로드 — 진단은 수집, 성공 시 Fullstack.Features 설정
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/features"
)

// parseFeaturesIfPresent loads features.yaml when the KindFeatures SSOT is present.
func parseFeaturesIfPresent(fs *Fullstack, root string, has map[SSOTKind]DetectedSSOT) {
	if _, ok := has[KindFeatures]; !ok {
		return
	}
	feats, diags := features.Load(root)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.Features = feats
	}
}
