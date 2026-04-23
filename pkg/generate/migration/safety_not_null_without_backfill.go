//ff:func feature=migration type=util control=sequence
//ff:what safetyNotNullWithoutBackfill — MIG-002 (AlterColumnNullable, To=false, backfill 없음)
package migration

import "fmt"

// safetyNotNullWithoutBackfill emits MIG-002 when a NOT NULL is set
// without a backfill clause.
func safetyNotNullWithoutBackfill(v AlterColumnNullable) []SafetyIssue {
	if v.To || v.Backfill != "" {
		return nil
	}
	return []SafetyIssue{{
		Level:   SafetyError,
		RuleID:  "MIG-002",
		Message: fmt.Sprintf("NOT NULL added to %s.%s without @backfill hint", v.Table, v.Column),
		Advice:  fmt.Sprintf("add `-- @backfill default=<value>` on the %s column line in specs/db/*.sql", v.Column),
	}}
}
