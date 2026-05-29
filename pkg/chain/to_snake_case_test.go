//ff:func feature=chain type=test control=iteration dimension=1
//ff:what toSnakeCase PascalCase -> snake_case 변환을 케이스별로 검증
package chain

import "testing"

func TestToSnakeCase(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{
		{"", ""},
		{"Course", "course"},
		{"GetCourse", "get_course"},
		{"hashPassword", "hash_password"},
		{"ID", "i_d"},
		{"already_snake", "already_snake"},
		{"A", "a"},
	}
	for _, c := range cases {
		if got := toSnakeCase(c.in); got != c.want {
			t.Errorf("toSnakeCase(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
