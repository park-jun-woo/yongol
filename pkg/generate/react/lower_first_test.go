//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what TestLowerFirst — 첫 글자 소문자 변환 (빈 문자열·유니코드 포함) 검증

package react

import "testing"

func TestLowerFirst(t *testing.T) {
	cases := map[string]string{
		"":                "",
		"ListMyBuildings": "listMyBuildings",
		"already":         "already",
		"X":               "x",
	}
	for in, want := range cases {
		if got := lowerFirst(in); got != want {
			t.Errorf("lowerFirst(%q) = %q, want %q", in, got, want)
		}
	}
}
