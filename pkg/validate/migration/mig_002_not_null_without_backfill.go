//ff:func feature=validate type=rule control=iteration dimension=1 topic=migration-safety
//ff:what MIG-002 — NOT NULL 추가인데 @backfill 힌트도 DEFAULT 도 없음 → ERROR (emit 차단)
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// Mig002NotNullWithoutBackfill promotes safety issues carrying RuleID
// MIG-002 to diagnostics. CheckSafety does the classification.
func Mig002NotNullWithoutBackfill(issues []migration.SafetyIssue) []diagnostic.Diagnostic {
	return emitByRule(issues, "MIG-002", diagnostic.LevelError)
}

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
