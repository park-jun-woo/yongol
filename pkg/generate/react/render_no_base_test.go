//ff:func feature=gen-react type=test control=sequence
//ff:what renderComponentTSX base 없는 컴포넌트 className 전달 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

func TestRenderComponentTSX_NoBase(t *testing.T) {
	tok := design.ComponentToken{}
	src := renderComponentTSX("Panel", tok)
	if !strings.Contains(src, "className={className}") {
		t.Errorf("no-base component should just pass className through\n\ngot:\n%s", src)
	}
}
