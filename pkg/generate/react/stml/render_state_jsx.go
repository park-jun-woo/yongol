//ff:func feature=stml-gen type=generator control=sequence
//ff:what StateBind의 조건부 렌더링 JSX를 Condition 패턴에 따라 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderStateJSX generates JSX for a StateBind.
func renderStateJSX(s stmlparser.StateBind, dataVar string, indent int, noBodyOps map[string]bool) string {
	ind := indentStr(indent)
	tag := orDefault(s.Tag, "div")
	cls := clsAttr(s.ClassName)

	cond := resolveStateCondition(s.Condition, dataVar)

	// Has children (e.g. action inside state)
	if len(s.Children) > 0 {
		var lines []string
		lines = append(lines, fmt.Sprintf("%s{%s && (", ind, cond))
		lines = append(lines, fmt.Sprintf("%s  <%s%s>", ind, tag, cls))
		lines = append(lines, renderChildNodes(s.Children, dataVar, "item", indent+4, noBodyOps)...)
		lines = append(lines, fmt.Sprintf("%s  </%s>", ind, tag))
		lines = append(lines, fmt.Sprintf("%s)}", ind))
		return strings.Join(lines, "\n")
	}

	// Simple text
	text := orDefault(s.Text, "")
	if text != "" {
		return fmt.Sprintf("%s{%s && <%s%s>%s</%s>}", ind, cond, tag, cls, text, tag)
	}

	return fmt.Sprintf("%s{%s && <%s%s />}", ind, cond, tag, cls)
}
