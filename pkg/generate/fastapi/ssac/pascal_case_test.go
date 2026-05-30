//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestPascalCase — snake_case → PascalCase 변환 (빈 세그먼트/단일/다중)

package ssac

import "testing"

func TestPascalCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"user", "User"},
		{"user_account", "UserAccount"},
		{"", ""},
		{"_leading", "Leading"},
		{"trailing_", "Trailing"},
		{"double__underscore", "DoubleUnderscore"},
		{"a", "A"},
	}
	for _, c := range cases {
		if got := pascalCase(c.in); got != c.want {
			t.Errorf("pascalCase(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
