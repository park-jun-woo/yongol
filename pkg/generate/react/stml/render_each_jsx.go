//ff:func feature=stml-gen type=generator control=sequence
//ff:what EachBlock의 배열 순회 JSX를 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderEachJSX generates JSX for an EachBlock.
func renderEachJSX(e stmlparser.EachBlock, dataVar string, indent int) string {
	ind := indentStr(indent)
	tag := orDefault(e.Tag, "div")
	cls := clsAttr(e.ClassName)
	itemTag := orDefault(e.ItemTag, "div")
	itemCls := clsAttr(e.ItemClassName)

	// Determine map callback parameters and key expression
	mapParams := "(item)"
	keyExpr := "item." + e.KeyField
	if e.KeyField == "" {
		mapParams = "(item, index)"
		keyExpr = "index"
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("%s<%s%s>", ind, tag, cls))
	lines = append(lines, fmt.Sprintf("%s  {%s.%s?.map(%s => (", ind, dataVar, e.Field, mapParams))
	lines = append(lines, fmt.Sprintf("%s    <%s key={%s}%s>", ind, itemTag, keyExpr, itemCls))

	if len(e.Children) > 0 {
		lines = append(lines, renderChildNodes(e.Children, "item", "item", indent+6)...)
	} else {
		for _, b := range e.Binds {
			lines = append(lines, renderBindJSX(b, "item", indent+6))
		}
	}

	lines = append(lines, fmt.Sprintf("%s    </%s>", ind, itemTag))
	lines = append(lines, fmt.Sprintf("%s  ))}", ind))
	lines = append(lines, fmt.Sprintf("%s</%s>", ind, tag))

	return strings.Join(lines, "\n")
}
