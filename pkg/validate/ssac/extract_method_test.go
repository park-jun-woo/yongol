//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what extractMethod — Method 추출 (정상/없음/점만/빈 문자열) 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestExtractMethod(t *testing.T) {
	cases := []struct {
		name  string
		model string
		want  string
	}{
		{name: "Model.Method", model: "User.Create", want: "Create"},
		{name: "no dot", model: "User", want: ""},
		{name: "dot at end", model: "User.", want: ""},
		{name: "empty", model: "", want: ""},
		{name: "dot at start", model: ".Method", want: ""},
		{name: "multiple dots", model: "A.B.C", want: "B.C"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			seq := parsessac.Sequence{Model: c.model}
			got := extractMethod(seq)
			if got != c.want {
				t.Errorf("extractMethod(%q) = %q, want %q", c.model, got, c.want)
			}
		})
	}
}
