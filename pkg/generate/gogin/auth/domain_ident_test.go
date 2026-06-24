//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what domainIdent — 대문자/하이픈/선행숫자/빈 이름 정규화 분기 검증

package auth

import "testing"

func TestDomainIdent(t *testing.T) {
	cases := map[string]string{
		"public":   "public",
		"My-Admin": "my_admin",
		"2nd":      "_2nd",
		"":         "_",
	}
	for in, want := range cases {
		if got := domainIdent(in); got != want {
			t.Errorf("domainIdent(%q) = %q, want %q", in, got, want)
		}
	}
}
