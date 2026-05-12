//ff:func feature=gen-react type=generator control=sequence
//ff:what renderComponentTSX — ComponentToken → TSX 소스 문자열 렌더링 (variants/sizes 유무에 따라 분기)

package react

import (
	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// renderComponentTSX produces a complete TSX source for a single component.
// Components with variants and/or sizes get a variant-aware forwardRef
// component. Simple components (base only) get a minimal forwardRef wrapper.
func renderComponentTSX(name string, tok design.ComponentToken) string {
	hasVariants := len(tok.Variants) > 0
	hasSizes := len(tok.Sizes) > 0

	if hasVariants || hasSizes {
		return renderVariantComponent(name, tok)
	}
	return renderSimpleComponent(name, tok)
}
