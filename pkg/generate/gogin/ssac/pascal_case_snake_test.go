//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what pascalCaseSnake 단위 테스트

package ssac

import "testing"

func TestPascalCaseSnake(t *testing.T) {
	cases := map[string]string{
		"org_id":     "OrgId",
		"created_at": "CreatedAt",
		"__leading":  "Leading",
		"trailing__": "Trailing",
		"a__b":       "AB",
		"single":     "Single",
	}
	for in, want := range cases {
		if got := pascalCaseSnake(in); got != want {
			t.Errorf("pascalCaseSnake(%q) = %q, want %q", in, got, want)
		}
	}
}
