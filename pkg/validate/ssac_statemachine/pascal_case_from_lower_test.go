//ff:func feature=validate type=test control=selection topic=states
//ff:what TestPascalCaseFromLower — pascalCaseFromLower 첫 글자 대문자화 분기 검증

package ssac_statemachine

import "testing"

func TestPascalCaseFromLower(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"workflow", "Workflow"},
		{"a", "A"},
	}
	for _, tc := range cases {
		t.Run(tc.in, func(t *testing.T) {
			if got := pascalCaseFromLower(tc.in); got != tc.want {
				t.Errorf("pascalCaseFromLower(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
