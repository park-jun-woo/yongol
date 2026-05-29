//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=size-parse
//ff:what TestParseSize_Invalid — 잘못된 입력은 모두 에러 반환

package middleware

import "testing"

func TestParseSize_Invalid(t *testing.T) {
	cases := []string{"", "abc", "-1MiB", "1ZZ"}
	for _, c := range cases {
		if _, err := ParseSize(c); err == nil {
			t.Errorf("ParseSize(%q) expected error, got nil", c)
		}
	}
}
