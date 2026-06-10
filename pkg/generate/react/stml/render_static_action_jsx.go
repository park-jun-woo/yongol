//ff:func feature=stml-gen type=generator control=sequence
//ff:what Action 폼 내부의 StaticElement JSX를 렌더링한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderStaticActionJSX renders a StaticElement inside an action form.
// A data-on-error marker element becomes a conditional error render bound
// to errVar instead of a plain static element.
func renderStaticActionJSX(se stmlparser.StaticElement, formName, errVar string, indent int) string {
	if se.OnError && errVar != "" {
		return renderOnErrorElement(se, errVar, indent)
	}

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
	lines = append(lines, renderActionChildNodes(se.Children, formName, errVar, indent+2)...)
	lines = append(lines, fmt.Sprintf("%s</%s>", ind, tag))
	return strings.Join(lines, "\n")
}
