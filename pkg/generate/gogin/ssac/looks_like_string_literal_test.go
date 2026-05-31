//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what wrapJSONBLiteral / looksLikeStringLiteral 단위 테스트 (JSONB 컬럼 string 리터럴 → []byte)
package ssac

import (
	"testing"
)

func TestLooksLikeStringLiteral(t *testing.T) {
	cases := map[string]bool{
		`"x"`: true,
		`""`:  true,
		`"a`:  false,
		`a"`:  false,
		"x":   false,
		"":    false,
		`"`:   false,
	}
	for in, want := range cases {
		if got := looksLikeStringLiteral(in); got != want {
			t.Errorf("looksLikeStringLiteral(%q) = %v, want %v", in, got, want)
		}
	}
}
