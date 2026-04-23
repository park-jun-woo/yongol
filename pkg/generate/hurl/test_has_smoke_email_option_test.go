//ff:func feature=gen-hurl type=test-helper control=iteration dimension=1
//ff:what hasSmokeEmailOption — Options 슬라이스에서 smoke_email={{newUuid}} 시드 옵션 존재 검사

package hurl

import "strings"

// hasSmokeEmailOption returns true when opts contains a "smoke_email=..."
// entry that includes `{{newUuid}}`. Extracted so the email-uniqueness
// test stays within the depth=2 nesting budget.
func hasSmokeEmailOption(opts []string) bool {
	for _, opt := range opts {
		if strings.HasPrefix(opt, "smoke_email=") && strings.Contains(opt, "{{newUuid}}") {
			return true
		}
	}
	return false
}
