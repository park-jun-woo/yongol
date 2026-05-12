//ff:func feature=gen-react type=generator control=sequence
//ff:what base-only forwardRef TSX 컴포넌트 소스를 생성한다

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
)

// renderSimpleComponent produces a base-only forwardRef component.
func renderSimpleComponent(name string, tok design.ComponentToken) string {
	var b strings.Builder

	b.WriteString("import * as React from 'react'\n")
	b.WriteString("import { cn } from '@/lib/utils'\n\n")

	htmlTag := inferHTMLTag(name)
	htmlElement := inferHTMLElement(htmlTag)

	b.WriteString(fmt.Sprintf("export const %s = React.forwardRef<%s, React.%s>(\n", name, htmlElement, htmlAttrsType(htmlTag)))
	if tok.Base != "" {
		b.WriteString(fmt.Sprintf("  ({ className, ...props }, ref) => (\n"))
		b.WriteString(fmt.Sprintf("    <%s ref={ref} className={cn('%s', className)} {...props} />\n", htmlTag, tok.Base))
	} else {
		b.WriteString(fmt.Sprintf("  ({ className, ...props }, ref) => (\n"))
		b.WriteString(fmt.Sprintf("    <%s ref={ref} className={className} {...props} />\n", htmlTag))
	}
	b.WriteString("  ),\n")
	b.WriteString(")\n")
	b.WriteString(fmt.Sprintf("%s.displayName = '%s'\n", name, name))

	return b.String()
}
