//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestClassifySingleTokenRounded — classifySingleToken rounded-prefix early-return 분기 검증

package stml_design

import "testing"

func TestClassifySingleTokenRounded(t *testing.T) {
	var out pageTokenRefs
	classifySingleToken("rounded-card", "rounded-card", "p.stml", &out)
	if len(out.Rounded) != 1 || out.Rounded[0].Name != "card" {
		t.Errorf("rounded = %+v", out.Rounded)
	}
}
