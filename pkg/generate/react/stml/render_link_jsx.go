//ff:func feature=stml-gen type=generator control=sequence
//ff:what LinkRef를 react-router <Link to=...> JSX로 렌더링한다 (자식은 bind/static 렌더 재사용)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderLinkJSX renders a LinkRef as a react-router <Link> element. The
// `to` value is the target page's resolved route pattern with the link's
// param sources substituted (LinkToAttr); children reuse the existing
// bind/static renderers (page-flow Phase007).
func renderLinkJSX(lr stmlparser.LinkRef, dataVar, itemVar string, indent int, noBodyOps map[string]bool, ctx bindCtx) string {
	ind := indentStr(indent)
	cls := clsAttr(lr.ClassName)
	to := LinkToAttr(lr)

	if len(lr.Children) == 0 {
		return fmt.Sprintf("%s<Link %s%s>%s</Link>", ind, to, cls, lr.Text)
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("%s<Link %s%s>", ind, to, cls))
	if lr.Text != "" {
		lines = append(lines, fmt.Sprintf("%s  %s", ind, lr.Text))
	}
	lines = append(lines, renderChildNodes(lr.Children, dataVar, itemVar, indent+2, noBodyOps, ctx)...)
	lines = append(lines, fmt.Sprintf("%s</Link>", ind))
	return strings.Join(lines, "\n")
}
