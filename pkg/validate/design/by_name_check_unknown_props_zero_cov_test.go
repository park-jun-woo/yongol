//ff:func feature=validate type=test control=sequence
//ff:what TestByName_ZeroCov — design 토큰 참조/미지 prop 검사 헬퍼 직접 호출
package design

import (
	"testing"
)

func TestByNameCheckUnknownProps_ZeroCov(t *testing.T) {
	props := map[string]string{
		"variant":    "primary", // known
		"weirdoProp": "x",       // unknown → warn
	}
	diags := checkUnknownProps("DESIGN.md", "Button", props)
	if len(diags) != 1 {
		t.Errorf("expected 1 unknown-prop warning, got %d", len(diags))
	}
}
