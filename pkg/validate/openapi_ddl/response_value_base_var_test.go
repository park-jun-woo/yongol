//ff:func feature=validate type=test control=iteration dimension=1 topic=openapi-ddl
//ff:what responseValueBaseVar — bare/dotted var 의 base var 추출 및 리터럴(따옴표/숫자/bool/null) → "" 검증

package openapi_ddl

import "testing"

func TestResponseValueBaseVar(t *testing.T) {
	cases := []struct{ in, want string }{
		{"rule", "rule"},
		{"rule.Name", "rule"},
		{"  rule.Name  ", "rule"},
		{"", ""},
		{`"literal"`, ""},
		{"'x'", ""},
		{"-1", ""},
		{"42", ""},
		{"true", ""},
		{"false", ""},
		{"null", ""},
		{"nil", ""},
	}
	for _, c := range cases {
		if got := responseValueBaseVar(c.in); got != c.want {
			t.Errorf("responseValueBaseVar(%q) = %q, want %q", c.in, got, c.want)
		}
	}
}
