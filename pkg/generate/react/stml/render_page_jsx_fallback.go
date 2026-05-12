//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what Children이 없는 페이지의 Fetch/Action을 순회하며 JSX를 렌더링한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageJSXFallback(page stmlparser.PageSpec, sb *strings.Builder) {
	for _, f := range page.Fetches {
		sb.WriteString(renderFetchJSX(f, 6))
		sb.WriteString("\n")
	}
	for _, a := range page.Actions {
		sb.WriteString(renderActionJSX(a, 6))
		sb.WriteString("\n")
	}
}
