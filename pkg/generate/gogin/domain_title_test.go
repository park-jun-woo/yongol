//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what domainTitle — 빈문자열/하이픈/대문자/숫자선행 PascalCase 변환 테이블 테스트
package gogin

import "testing"

func TestDomainTitle(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"plain", "public", "Public"},
		{"admin", "admin", "Admin"},
		{"hyphen", "my-admin", "MyAdmin"},
		{"multiHyphen", "a-b-c", "ABC"},
		{"uppercaseInput", "Admin", "Admin"},
		{"underscore", "api_x", "ApiX"},
		{"leadingDigit", "2nd", "2nd"},
		{"digitsAndLetters", "shop9", "Shop9"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := domainTitle(c.in); got != c.want {
				t.Fatalf("domainTitle(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
