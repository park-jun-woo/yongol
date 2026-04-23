//ff:func feature=migration type=util control=sequence
//ff:what safetyAddColumnNotNull — MIG-002 (AddColumn, NOT NULL, default/backfill 없음)
package migration

import "fmt"

// safetyAddColumnNotNull emits MIG-002 when a NOT NULL column is added
// without a default or backfill.
func safetyAddColumnNotNull(v AddColumn) []SafetyIssue {
	if v.Column.Nullable || v.Column.Default != "" || v.Backfill != "" {
		return nil
	}
	return []SafetyIssue{{
		Level:   SafetyError,
		RuleID:  "MIG-002",
		Message: fmt.Sprintf("NOT NULL column %s.%s added without DEFAULT or @backfill", v.Table, v.Column.Name),
		Advice:  fmt.Sprintf("set DEFAULT in DDL or add `-- @backfill default=<value>` on %s line", v.Column.Name),
	}}
}
