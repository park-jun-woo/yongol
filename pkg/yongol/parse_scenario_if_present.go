//ff:func feature=orchestrator type=loader control=sequence
//ff:what Hurl 시나리오 탐지 시 파일 수집 + 엔트리 파싱
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// parseScenarioIfPresent collects .hurl files and parses all entries.
func parseScenarioIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindScenario]
	if !ok {
		return
	}
	fs.HurlFiles = hurl.CollectFiles(d.Path)
	appendHurlEntries(fs)
}
