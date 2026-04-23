//ff:func feature=validate type=rule control=sequence topic=migration-safety
//ff:what MIG-004 — DROP TABLE/COLUMN 인데 @allow_destructive 없음 → WARNING
package migration

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/generate/migration"
)

// Mig004DestructiveWithoutAllow surfaces MIG-004 issues at WARNING level.
func Mig004DestructiveWithoutAllow(issues []migration.SafetyIssue) []diagnostic.Diagnostic {
	return emitByRule(issues, "MIG-004", diagnostic.LevelWarning)
}
