//ff:func feature=gen-react type=generator control=sequence
//ff:what variant/size 지원 forwardRef TSX 컴포넌트 소스를 생성한다

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// renderVariantComponent produces a forwardRef component with variant/size support.
func renderVariantComponent(name string, tok design.ComponentToken) string {
	var b strings.Builder

	b.WriteString("import * as React from 'react'\n")
	b.WriteString("import { cn } from '@/lib/utils'\n\n")

	htmlTag := inferHTMLTag(name)
	htmlElement := inferHTMLElement(htmlTag)

	variantKeys := sortedKeys(tok.Variants)
	sizeKeys := sortedKeys(tok.Sizes)

	writeVariantTypes(&b, variantKeys, sizeKeys)
	writeVariantProps(&b, name, htmlTag, variantKeys, sizeKeys)
	writeVariantRecords(&b, variantKeys, sizeKeys, tok)

	defaultVariant := tok.DefaultVariant
	if defaultVariant == "" && len(variantKeys) > 0 {
		defaultVariant = variantKeys[0]
	}
	defaultSize := tok.DefaultSize
	if defaultSize == "" && len(sizeKeys) > 0 {
		defaultSize = sizeKeys[0]
	}

	b.WriteString(fmt.Sprintf("export const %s = React.forwardRef<%s, %sProps>(\n", name, htmlElement, name))

	destructParams := buildDestructParams(variantKeys, sizeKeys, defaultVariant, defaultSize)
	b.WriteString(fmt.Sprintf("  ({ %s }, ref) => (\n", strings.Join(destructParams, ", ")))

	cnArgs := buildCNArgs(tok.Base, variantKeys, sizeKeys)
	b.WriteString(fmt.Sprintf("    <%s\n", htmlTag))
	b.WriteString("      ref={ref}\n")
	b.WriteString(fmt.Sprintf("      className={cn(%s)}\n", strings.Join(cnArgs, ", ")))
	b.WriteString("      {...props}\n")
	b.WriteString("    />\n")
	b.WriteString("  ),\n")
	b.WriteString(")\n")
	b.WriteString(fmt.Sprintf("%s.displayName = '%s'\n", name, name))

	return b.String()
}
