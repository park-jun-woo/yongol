//ff:type feature=migration type=model
//ff:what SafetyIssue — CheckSafety 가 생성하는 단일 진단 (Rule ID + Level + 메시지)
package migration

// SafetyIssue carries one diagnostic produced by CheckSafety. The Rule
// ID lets the CLI / validate pipeline map it to MIG-00N entries.
type SafetyIssue struct {
	Level   SafetyLevel
	RuleID  string // "MIG-002" / "MIG-004" / "MIG-005"
	Message string
	Advice  string // suggested fix (hint to add, etc)
}
