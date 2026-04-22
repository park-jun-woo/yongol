//ff:func feature=orchestrator type=logger control=sequence
//ff:what logBuiltinLoadIssues — built-in funcspec 로딩 실패를 stderr 구조화 경고로 노출

package yongol

import (
	"log/slog"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// logBuiltinLoadIssues emits a RuntimeWarning via slog.Warn when built-in
// funcspec loading from sibling repo `github.com/park-jun-woo/ssac` runs into
// parse diagnostics. These are environment/install issues — not user SSOT
// errors — so they do not flow into ParseDiagnostics. Surfacing them via
// slog.Warn lets operators and CI pick up the root cause without polluting the
// user-facing validate flow. Only the first diagnostic message is included to
// avoid log floods; diag_count lets callers gauge severity.
func logBuiltinLoadIssues(pkgRoot string, diags []diagnostic.Diagnostic) {
	first := ""
	if len(diags) > 0 {
		first = diags[0].Message
	}
	slog.Warn("built-in funcspec load issues",
		"pkg_root", pkgRoot,
		"diag_count", len(diags),
		"first_message", first,
	)
}
