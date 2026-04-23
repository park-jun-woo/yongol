//ff:func feature=migration type=util control=iteration dimension=1
//ff:what issuesToDiags — SafetyIssue 리스트 + 누락 sidecar 경로를 Diagnostic 으로 변환
package migration

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// issuesToDiags converts SafetyIssue records (produced by CheckSafety)
// and missing sidecar file paths into diagnostics. ERROR / WARNING
// level derives from SafetyIssue.Level.
func issuesToDiags(issues []SafetyIssue, missing []string) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for _, is := range issues {
		lvl := diagnostic.LevelWarning
		if is.Level == SafetyError {
			lvl = diagnostic.LevelError
		}
		out = append(out, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   lvl,
			Message: "[" + is.RuleID + "] " + is.Message,
			Advice:  is.Advice,
		})
	}
	for _, p := range missing {
		out = append(out, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[MIG-003] @data_migration file not found: " + p,
			Advice:  "create the sidecar file or fix the path in the hint",
		})
	}
	return out
}
