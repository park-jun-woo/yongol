//ff:func feature=gen-react type=test control=sequence
//ff:what renderComponentTSX variants만 있고 sizes 없는 경우 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestRenderComponentTSX_VariantsOnlyNoSizes(t *testing.T) {
	tok := design.ComponentToken{
		Base: "inline-flex items-center rounded-full px-2.5 py-0.5 text-xs font-medium",
		Variants: map[string]string{
			"default":     "bg-primary text-primary-foreground",
			"destructive": "bg-destructive text-destructive-foreground",
		},
		DefaultVariant: "default",
	}
	src := renderComponentTSX("Badge", tok)

	if !strings.Contains(src, "type Variant") {
		t.Error("should have Variant type")
	}
	if strings.Contains(src, "type Size") {
		t.Error("should not have Size type when no sizes defined")
	}
	if !strings.Contains(src, "variant = 'default'") {
		t.Errorf("should default to 'default' variant\n\ngot:\n%s", src)
	}
}
