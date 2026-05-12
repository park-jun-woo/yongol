//ff:func feature=gen-react type=emitter control=sequence
//ff:what variant/size Props 타입 선언을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"
)

// writeVariantProps writes the component Props type declaration.
func writeVariantProps(b *strings.Builder, name, htmlTag string, variantKeys, sizeKeys []string) {
	b.WriteString(fmt.Sprintf("export type %sProps = React.%s & {\n", name, htmlAttrsType(htmlTag)))
	if len(variantKeys) > 0 {
		b.WriteString("  variant?: Variant\n")
	}
	if len(sizeKeys) > 0 {
		b.WriteString("  size?: Size\n")
	}
	b.WriteString("}\n\n")
}
