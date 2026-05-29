//ff:func feature=rule type=test control=iteration dimension=1
//ff:what TestBaseSpec_SpecName — BaseSpec.SpecName 이 Rule 필드를 그대로 반환하는지 테이블 기반 검증

package rule

import "testing"

func TestBaseSpec_SpecName(t *testing.T) {
	cases := []struct {
		in   BaseSpec
		want string
	}{
		{BaseSpec{Rule: "V11-1", Level: "ERROR"}, "V11-1"},
		{BaseSpec{Rule: "", Level: "ERROR"}, ""},
		{BaseSpec{Rule: "CC-usage", Level: "WARNING", Message: "m"}, "CC-usage"},
	}
	for _, c := range cases {
		if got := c.in.SpecName(); got != c.want {
			t.Errorf("SpecName(%+v) = %q; want %q", c.in, got, c.want)
		}
	}
}
