//ff:func feature=orchestrator type=loader control=sequence
//ff:what yongol built-in pkg/ 의 Func 스펙을 로드 (사용자 스펙 오류 아님 — ParseDiagnostics 제외, 운영자 가시성은 slog.Warn 으로 노출)
package yongol

import (
	"log/slog"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// parseYongolPkgSpecs loads built-in pkg/ Func specs. Built-in specs are not
// user spec errors, so they do not flow into ParseDiagnostics. However, a
// silent drop would make it hard for operators to diagnose root causes
// (missing ssac repo, stale version, filesystem issues), so any loading
// problem is surfaced as a structured slog.Warn. When only some files fail,
// the remaining specs are still loaded to ensure graceful degradation.
func parseYongolPkgSpecs(fs *Fullstack) {
	pkgRoot := findYongolPkgRoot()
	if pkgRoot == "" {
		slog.Debug("built-in pkg root not found — skipping built-in specs")
		return
	}
	specs, diags := funcspec.ParseDir(pkgRoot)
	if len(diags) > 0 {
		logBuiltinLoadIssues(pkgRoot, diags)
	}
	if len(specs) > 0 {
		fs.YongolPkgSpecs = specs
	}
}
