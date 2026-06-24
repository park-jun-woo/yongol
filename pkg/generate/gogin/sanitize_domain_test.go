//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what sanitizeDomainName — 하이픈/대문자/숫자선행/특수문자 정규화 테이블 테스트
package gogin

import "testing"

func TestSanitizeDomainName(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want string
	}{
		{"plain", "public", "public"},
		{"hyphen", "my-admin", "my_admin"},
		{"multiHyphen", "a-b-c", "a_b_c"},
		{"uppercase", "Admin", "admin"},
		{"mixedCase", "MyAdmin", "myadmin"},
		{"leadingDigit", "2nd", "_2nd"},
		{"dot", "v1.api", "v1_api"},
		{"special", "a@b!", "a_b_"},
		{"underscoreKept", "api_x", "api_x"},
		{"digitsAndLetters", "shop9", "shop9"},
		{"empty", "", "_"},
		{"allSpecial", "!!!", "___"},
		{"leadingDigitAfterStrip", "9", "_9"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := sanitizeDomainName(c.in)
			if got != c.want {
				t.Fatalf("sanitizeDomainName(%q) = %q, want %q", c.in, got, c.want)
			}
		})
	}
}
