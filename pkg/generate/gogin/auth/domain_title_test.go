//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what domainTitle — 빈 이름 → "" + 하이픈/선행숫자 → PascalCase 분기 검증

package auth

import "testing"

func TestDomainTitle(t *testing.T) {
	cases := map[string]string{
		"":         "",
		"public":   "Public",
		"my-admin": "MyAdmin",
		"2nd":      "2nd",
	}
	for in, want := range cases {
		if got := domainTitle(in); got != want {
			t.Errorf("domainTitle(%q) = %q, want %q", in, got, want)
		}
	}
}
