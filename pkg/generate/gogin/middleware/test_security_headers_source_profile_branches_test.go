//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=security-headers
//ff:what TestSecurityHeadersSource_ProfileBranches — profile 분기 문자열 잔존 확인

package middleware

import (
	"strings"
	"testing"
)

// TestSecurityHeadersSource_ProfileBranches — the profile branches are
// expressed as string comparisons in source; verify each branch literal is
// present so a refactor cannot silently remove a profile.
func TestSecurityHeadersSource_ProfileBranches(t *testing.T) {
	for _, must := range []string{
		`profile != "dev"`,
		`profile == "api"`,
		`profile == "dev"`,
	} {
		if !strings.Contains(securityHeadersSource, must) {
			t.Errorf("security_headers source missing profile branch %q", must)
		}
	}
}
