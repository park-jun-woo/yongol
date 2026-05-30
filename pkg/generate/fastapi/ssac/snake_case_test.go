//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestSnakeCase — PascalCase/camelCase → snake_case (연속 대문자 약어 처리)

package ssac

import "testing"

func TestSnakeCase(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"ID", "id"},
		{"OrgID", "org_id"},
		{"ResolveRootID", "resolve_root_id"},
		{"camelCase", "camel_case"},
		{"PascalCase", "pascal_case"},
		{"URLPath", "url_path"},
		{"already_snake", "already_snake"},
		{"A", "a"},
	}
	for _, c := range cases {
		if got := snakeCase(c.in); got != c.want {
			t.Errorf("snakeCase(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
