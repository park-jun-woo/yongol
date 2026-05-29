//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=domain-security
//ff:what countDiag — 진단 목록에서 prefix 일치 개수 반환
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// countDiag returns the count of diagnostics whose message starts with the prefix.
func countDiag(diags []diagnostic.Diagnostic, prefix string) int {
	n := 0
	for _, d := range diags {
		if len(d.Message) >= len(prefix) && d.Message[:len(prefix)] == prefix {
			n++
		}
	}
	return n
}
