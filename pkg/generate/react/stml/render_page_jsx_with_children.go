//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what Children이 있는 페이지의 루트 엘리먼트와 자식 노드 JSX를 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func renderPageJSXWithChildren(children []stmlparser.ChildNode, sb *strings.Builder) {
	rootTag := "div"
	rootCls := ""
	inner := children
	if len(children) == 1 && children[0].Kind == "static" {
		root := children[0].Static
		rootTag = root.Tag
		rootCls = root.ClassName
		inner = root.Children
	}
	sb.WriteString(fmt.Sprintf("    <%s%s>\n", rootTag, clsAttr(rootCls)))
	for _, line := range renderChildNodes(inner, "", "item", 6) {
		sb.WriteString(line)
		sb.WriteString("\n")
	}
	sb.WriteString(fmt.Sprintf("    </%s>\n", rootTag))
}
