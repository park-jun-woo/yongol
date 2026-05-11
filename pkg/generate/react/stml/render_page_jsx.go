//ff:func feature=stml-gen type=generator control=sequence
//ff:what 페이지의 JSX return 블록을 Children 유무에 따라 렌더링한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageJSX(page stmlparser.PageSpec, sb *strings.Builder) {
	sb.WriteString("  return (\n")

	if len(page.Children) > 0 {
		renderPageJSXWithChildren(page.Children, sb)
	} else {
		renderPageJSXFallback(page, sb)
	}

	sb.WriteString("  )\n")
}
