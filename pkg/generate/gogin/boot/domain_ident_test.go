//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what domainIdent — 하이픈/대문자/숫자선행/빈문자열 정규화 (sanitizeDomainName 동치 검증)
package boot

import "testing"

func TestDomainIdent(t *testing.T) {
	cases := []struct{ in, want string }{
		{"public", "public"},
		{"my-admin", "my_admin"},
		{"Admin", "admin"},
		{"2nd", "_2nd"},
		{"v1.api", "v1_api"},
		{"", "_"},
		{"api_x", "api_x"},
	}
	for _, c := range cases {
		if got := domainIdent(c.in); got != c.want {
			t.Fatalf("domainIdent(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
