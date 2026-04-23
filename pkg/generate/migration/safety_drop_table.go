//ff:func feature=migration type=util control=sequence
//ff:what safetyDropTable — MIG-004 (DropTable + @allow_destructive 없음)
package migration

import "fmt"

// safetyDropTable emits MIG-004 when DROP TABLE lacks @allow_destructive.
func safetyDropTable(v DropTable) []SafetyIssue {
	if v.AllowDestructive {
		return nil
	}
	return []SafetyIssue{{
		Level:   SafetyWarning,
		RuleID:  "MIG-004",
		Message: fmt.Sprintf("DROP TABLE %s without @allow_destructive", v.Name),
		Advice:  "add `-- @allow_destructive` above the (removed) CREATE TABLE or keep the table",
	}}
}
