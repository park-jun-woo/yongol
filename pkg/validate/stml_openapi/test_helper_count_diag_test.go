//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what countDiag — 진단 목록에서 주어진 규칙 prefix 매칭 횟수 반환

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// countDiag returns the number of diagnostics matching the given rule prefix.
func countDiag(diags []diagnostic.Diagnostic, rulePrefix string) int {
	n := 0
	for _, d := range diags {
		if len(d.Message) >= len(rulePrefix) && d.Message[:len(rulePrefix)] == rulePrefix {
			n++
		}
	}
	return n
}
