//ff:func feature=validate type=test control=iteration dimension=1 topic=ssac-structural
//ff:what inputValueRefBase — input 값에서 선두 변수 추출 (정상/literal/empty/quoted) 검증

package ssac

import "testing"

func TestInputValueRefBase(t *testing.T) {
	cases := []struct {
		name string
		val  string
		want string
	}{
		{name: "simple var", val: "user", want: "user"},
		{name: "var with field", val: "user.ID", want: "user"},
		{name: "empty", val: "", want: ""},
		{name: "quoted string", val: `"hello"`, want: ""},
		{name: "numeric literal", val: "42", want: ""},
		{name: "true literal", val: "true", want: ""},
		{name: "false literal", val: "false", want: ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := inputValueRefBase(c.val)
			if got != c.want {
				t.Errorf("inputValueRefBase(%q) = %q, want %q", c.val, got, c.want)
			}
		})
	}
}
