//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
)

func TestByNameColorAndTailwind_ZeroCov(t *testing.T) {
	var b strings.Builder
	writeColorToken(&b, "primary", "#fff", "#000")
	writeExtraDesignColors(&b, map[string]string{"brand": "#123", "accent2": "#456"})
	if b.Len() == 0 {
		t.Errorf("color writers empty")
	}

	theme := &manifest.FrontendTheme{Radius: "0.5rem"}
	dspec := &design.DesignSpec{
		Rounded: map[string]string{"sm": "2px", "lg": "8px"},
		Spacing: map[string]string{"xs": "4px", "xl": "32px"},
	}
	var b2 strings.Builder
	writeTailwindBorderRadius(&b2, theme, dspec)
	writeTailwindSpacing(&b2, dspec)
	if b2.Len() == 0 {
		t.Errorf("tailwind writers empty")
	}
}
