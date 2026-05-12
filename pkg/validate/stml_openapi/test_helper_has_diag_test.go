//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-openapi
//ff:what hasDiag — 진단 목록에서 주어진 규칙 prefix 존재 여부 확인

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// hasDiag returns true if any diagnostic has the given rule prefix in its message.
func hasDiag(diags []diagnostic.Diagnostic, rulePrefix string) bool {
	for _, d := range diags {
		if len(d.Message) >= len(rulePrefix) && d.Message[:len(rulePrefix)] == rulePrefix {
			return true
		}
	}
	return false
}
