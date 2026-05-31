//ff:func feature=gen-react type=test control=sequence
//ff:what TestByName_ZeroCov — react 코드젠 헬퍼들을 이름으로 직접 호출해 커버리지 귀속
package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestByNameRenderComponents_ZeroCov(t *testing.T) {
	simple := renderSimpleComponent("Input", design.ComponentToken{Base: "px-2"})
	if !strings.Contains(simple, "Input") {
		t.Errorf("renderSimpleComponent missing name")
	}
	_ = renderSimpleComponent("Card", design.ComponentToken{})

	variant := renderVariantComponent("Button", design.ComponentToken{
		Base:           "px-4",
		Variants:       map[string]string{"primary": "bg-blue", "ghost": "bg-none"},
		Sizes:          map[string]string{"sm": "text-sm", "lg": "text-lg"},
		DefaultVariant: "primary",
		DefaultSize:    "sm",
	})
	if !strings.Contains(variant, "Button") {
		t.Errorf("renderVariantComponent missing name")
	}
}
