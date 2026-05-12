//ff:func feature=gen-react type=emitter control=sequence
//ff:what variant/size TypeScript 타입 선언을 Builder에 기록한다

package react

import "strings"

// writeVariantTypes writes Variant and Size type declarations.
func writeVariantTypes(b *strings.Builder, variantKeys, sizeKeys []string) {
	if len(variantKeys) > 0 {
		b.WriteString("type Variant = " + quotedUnion(variantKeys) + "\n")
	}
	if len(sizeKeys) > 0 {
		b.WriteString("type Size = " + quotedUnion(sizeKeys) + "\n")
	}
	b.WriteString("\n")
}
