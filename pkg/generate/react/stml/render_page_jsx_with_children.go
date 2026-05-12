//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what Children이 있는 페이지의 자식 노드 JSX를 렌더링한다
package stml

import (
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageJSXWithChildren(children []stmlparser.ChildNode, sb *strings.Builder, noBodyOps map[string]bool) {
	inner := children
	if len(children) == 1 && children[0].Kind == "static" {
		inner = children[0].Static.Children
	}
	for _, line := range renderChildNodes(inner, "", "item", 6, noBodyOps) {
		sb.WriteString(line)
		sb.WriteString("\n")
	}
}
