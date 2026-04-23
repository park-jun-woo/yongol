//ff:func feature=validate type=util control=iteration dimension=1 topic=migration-safety
//ff:what SafetyIssue 집합에서 특정 RuleID 만 골라 Diagnostic 리스트로 변환

package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// emitByRule converts SafetyIssues carrying the given ruleID into
// diagnostics at the specified level. Shared by MIG-002/004/005.
func emitByRule(issues []migration.SafetyIssue, ruleID string, lvl diagnostic.Level) []diagnostic.Diagnostic {
	var out []diagnostic.Diagnostic
	for _, is := range issues {
		if is.RuleID != ruleID {
			continue
		}
		out = append(out, diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   lvl,
			Message: "[" + ruleID + "] " + is.Message,
			Advice:  is.Advice,
		})
	}
	return out
}
