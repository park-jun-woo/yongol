//ff:func feature=orchestrator type=loader control=sequence
//ff:what 경로의 STML 디렉토리를 stml.ParseDir 로 파싱 — 단일/도메인 공용 래퍼, 진단은 호출자가 수집
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// parseSTMLDir parses all .html page files under dir via stml.ParseDir. It is a
// thin, path-based wrapper so the single-site loader (parseSTMLIfPresent) and the
// per-domain loop (parseDomainsIfPresent) share one entry point; diagnostics are
// returned to the caller for collection.
func parseSTMLDir(dir string) ([]stml.PageSpec, []diagnostic.Diagnostic) {
	return stml.ParseDir(dir)
}
