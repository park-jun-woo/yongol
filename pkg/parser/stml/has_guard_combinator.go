//ff:func feature=stml-parse type=util control=sequence dimension=1
//ff:what 조건 문자열에 결합·그룹 토큰(&& || 선행! 괄호)이 있는지 판별한다
package stml

import "strings"

// HasGuardCombinator reports whether the condition uses a guard combining or
// grouping token (&&, ||, a leading !, or parentheses). Only such conditions are
// routed through the guard AST converter (codegen) and validated by TM-17;
// single comparisons and lifecycle suffixes keep the legacy path so existing
// behavior is preserved.
func HasGuardCombinator(condition string) bool {
	if strings.Contains(condition, "&&") || strings.Contains(condition, "||") {
		return true
	}
	if strings.Contains(condition, "(") || strings.Contains(condition, ")") {
		return true
	}
	return strings.HasPrefix(strings.TrimSpace(condition), "!")
}
