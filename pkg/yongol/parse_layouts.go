//ff:func feature=orchestrator type=loader control=sequence
//ff:what 경로의 layouts/ 디렉토리를 stml.ParseLayoutDir 로 파싱 — 단일/도메인 공용 래퍼, 진단은 호출자가 수집
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// parseLayouts parses all layout .html files under dir via stml.ParseLayoutDir.
// It is a thin, path-based wrapper shared by the single-site loader
// (parseLayoutIfPresent) and the per-domain loop (parseDomainsIfPresent).
func parseLayouts(dir string) ([]stml.LayoutSpec, []diagnostic.Diagnostic) {
	return stml.ParseLayoutDir(dir)
}
