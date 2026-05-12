//ff:func feature=gen-react type=emitter control=sequence
//ff:what variant/size Record 상수 선언을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// writeVariantRecords writes variants and sizes Record declarations.
func writeVariantRecords(b *strings.Builder, variantKeys, sizeKeys []string, tok design.ComponentToken) {
	if len(variantKeys) > 0 {
		b.WriteString("const variants: Record<Variant, string> = {\n")
		for _, k := range variantKeys {
			b.WriteString(fmt.Sprintf("  %s: '%s',\n", k, tok.Variants[k]))
		}
		b.WriteString("}\n\n")
	}
	if len(sizeKeys) > 0 {
		b.WriteString("const sizes: Record<Size, string> = {\n")
		for _, k := range sizeKeys {
			b.WriteString(fmt.Sprintf("  %s: '%s',\n", k, tok.Sizes[k]))
		}
		b.WriteString("}\n\n")
	}
}
