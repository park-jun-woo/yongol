//ff:func feature=gen-react type=test control=sequence
//ff:what renderComponentTSX Input 태그 추론 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestRenderComponentTSX_InputTag(t *testing.T) {
	tok := design.ComponentToken{
		Base: "flex h-10 w-full rounded-md border",
	}
	src := renderComponentTSX("Input", tok)
	if !strings.Contains(src, "<input ref={ref}") {
		t.Errorf("Input component should use <input> tag\n\ngot:\n%s", src)
	}
	if !strings.Contains(src, "HTMLInputElement") {
		t.Errorf("Input component should reference HTMLInputElement\n\ngot:\n%s", src)
	}
}
