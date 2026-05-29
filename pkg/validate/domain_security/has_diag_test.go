//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=domain-security
//ff:what hasDiag — 진단 목록에서 prefix 일치 여부 확인
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// hasDiag returns true if any diagnostic message contains the given rule prefix.
func hasDiag(diags []diagnostic.Diagnostic, prefix string) bool {
	for _, d := range diags {
		if len(d.Message) >= len(prefix) && d.Message[:len(prefix)] == prefix {
			return true
		}
	}
	return false
}
