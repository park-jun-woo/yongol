//ff:func feature=validate type=rule control=iteration dimension=1 topic=migration-safety
//ff:what MIG-005 — 타입 변경 + @cast 힌트 없음 (int↔text 등) → WARNING
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// Mig005CastMissing surfaces MIG-005 issues at WARNING level.
func Mig005CastMissing(issues []migration.SafetyIssue) []diagnostic.Diagnostic {
	return emitByRule(issues, "MIG-005", diagnostic.LevelWarning)
}
