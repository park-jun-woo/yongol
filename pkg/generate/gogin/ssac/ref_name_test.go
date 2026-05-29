//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what refName 단위 테스트 ($ref 경로에서 스키마 이름 추출)

package ssac

import "testing"

func TestRefName(t *testing.T) {
	cases := map[string]string{
		"#/components/schemas/Workflow": "Workflow",
		"#/components/schemas/User":     "User",
		"Plain":                         "Plain",
		"":                              "",
		"a/b/c":                         "c",
	}
	for in, want := range cases {
		if got := refName(in); got != want {
			t.Errorf("refName(%q) = %q, want %q", in, got, want)
		}
	}
}
