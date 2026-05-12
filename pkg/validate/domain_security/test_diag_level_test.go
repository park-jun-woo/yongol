//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=domain-security
//ff:what diagLevel — 진단 목록에서 prefix 일치하는 첫 진단의 레벨 반환
package domain_security

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// diagLevel returns the level of the first diagnostic matching the prefix.
func diagLevel(diags []diagnostic.Diagnostic, prefix string) diagnostic.Level {
	for _, d := range diags {
		if len(d.Message) >= len(prefix) && d.Message[:len(prefix)] == prefix {
			return d.Level
		}
	}
	return ""
}
