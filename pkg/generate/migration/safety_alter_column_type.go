//ff:func feature=migration type=util control=sequence
//ff:what safetyAlterColumnType — MIG-005 (타입 변경 + risky + @cast 없음)
package migration

import "fmt"

// safetyAlterColumnType emits MIG-005 when a risky type cast is done
// without an explicit @cast hint.
func safetyAlterColumnType(v AlterColumnType) []SafetyIssue {
	if v.Using != "" || !riskyCast(v.From, v.To) {
		return nil
	}
	return []SafetyIssue{{
		Level:   SafetyWarning,
		RuleID:  "MIG-005",
		Message: fmt.Sprintf("type change %s.%s %s → %s without @cast USING", v.Table, v.Column, v.From.SQL(), v.To.SQL()),
		Advice:  fmt.Sprintf("add `-- @cast using=<expr>` on %s line if default USING cast is unsafe", v.Column),
	}}
}
