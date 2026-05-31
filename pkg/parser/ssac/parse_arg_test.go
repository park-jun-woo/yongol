//ff:func feature=ssac-parse type=test control=iteration dimension=1
//ff:what parseArg / parseArgs / filterSubscribe / buildSubscribeInfo 단위 검증
package ssac

import (
	"testing"
)

func TestParseArg(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want Arg
	}{
		{"quoted literal", `"admin"`, Arg{Literal: "admin", IsQuoted: true}},
		{"numeric literal", "42", Arg{Literal: "42"}},
		{"float literal", "3.14", Arg{Literal: "3.14"}},
		{"source field", "request.CourseID", Arg{Source: "request", Field: "CourseID"}},
		{"bare variable", "course", Arg{Source: "course"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := parseArg(c.in)
			if got != c.want {
				t.Errorf("parseArg(%q) = %+v, want %+v", c.in, got, c.want)
			}
		})
	}
}
