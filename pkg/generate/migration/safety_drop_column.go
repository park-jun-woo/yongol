//ff:func feature=migration type=util control=sequence
//ff:what safetyDropColumn — MIG-004 (DropColumn + @allow_destructive 없음)
package migration

import "fmt"

// safetyDropColumn emits MIG-004 when DROP COLUMN lacks @allow_destructive.
func safetyDropColumn(v DropColumn) []SafetyIssue {
	if v.AllowDestructive {
		return nil
	}
	return []SafetyIssue{{
		Level:   SafetyWarning,
		RuleID:  "MIG-004",
		Message: fmt.Sprintf("DROP COLUMN %s.%s without @allow_destructive", v.Table, v.Column),
		Advice:  fmt.Sprintf("add `-- @allow_destructive` above CREATE TABLE %s", v.Table),
	}}
}
