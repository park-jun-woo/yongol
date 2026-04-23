//ff:func feature=migration type=util control=iteration dimension=1
//ff:what CheckSafety — Operation 리스트 전수 점검 → MIG-002/004/005 SafetyIssue 반환
package migration

import "sort"

// CheckSafety returns the list of issues found across the op list.
// Callers decide whether to abort (any SafetyError) or just warn
// (SafetyWarning only).
func CheckSafety(ops []Operation) []SafetyIssue {
	var issues []SafetyIssue
	for _, op := range ops {
		issues = append(issues, checkSafetyForOp(op)...)
	}
	sort.SliceStable(issues, func(i, j int) bool {
		if issues[i].RuleID != issues[j].RuleID {
			return issues[i].RuleID < issues[j].RuleID
		}
		return issues[i].Message < issues[j].Message
	})
	return issues
}
