//ff:func feature=stml-gen type=generator control=sequence
//ff:what StaticElement의 구조를 보존하며 JSX를 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderStaticJSX renders a StaticElement preserving structure.
func renderStaticJSX(se stmlparser.StaticElement, dataVar, itemVar string, indent int, noBodyOps map[string]bool) string {
	ind := indentStr(indent)
	tag := se.Tag
	cls := clsAttr(se.ClassName)

	if len(se.Children) == 0 {
		if se.Text != "" {
			return fmt.Sprintf("%s<%s%s>%s</%s>", ind, tag, cls, se.Text, tag)
		}
		return fmt.Sprintf("%s<%s%s />", ind, tag, cls)
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("%s<%s%s>", ind, tag, cls))
	if se.Text != "" {
		lines = append(lines, fmt.Sprintf("%s  %s", ind, se.Text))
	}
	lines = append(lines, renderChildNodes(se.Children, dataVar, itemVar, indent+2, noBodyOps)...)
	lines = append(lines, fmt.Sprintf("%s</%s>", ind, tag))
	return strings.Join(lines, "\n")
}
