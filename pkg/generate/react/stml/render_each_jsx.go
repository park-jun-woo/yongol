//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what EachBlock의 배열 순회 JSX를 Table 구조로 생성한다 (행 액션 셀 포함)
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderEachJSX generates JSX for an EachBlock as a Table. Row-level
// actions (page-flow Phase006) are rendered as one trailing cell per
// action inside the map callback, where `item` is in scope for the
// item.<Field> mutate arguments. A RowLink (data-link on the item
// template, page-flow Phase007) wraps every field cell's content in a
// <Link> so clicking the row content navigates to the target page;
// row-child links render as trailing cells like row actions.
func renderEachJSX(e stmlparser.EachBlock, dataVar string, indent int, noBodyOps map[string]bool) string {
	ind := indentStr(indent)
	cls := clsAttr(e.ClassName)

	// Determine map callback parameters and key expression
	mapParams := "(item)"
	keyExpr := "item." + e.KeyField
	if e.KeyField == "" {
		mapParams = "(item, index)"
		keyExpr = "index"
	}

	// Extract bind field names for table headers
	fields := extractBindFieldsFromChildren(e.Children)
	if len(fields) == 0 {
		for _, b := range e.Binds {
			fields = append(fields, b.Name)
		}
	}

	// Row actions in DOM order (direct children and static-nested)
	rowActions := collectAllActions(e.Children)
	// Row-child links in DOM order (direct children and static-nested)
	rowLinks := collectAllLinks(e.Children)

	// Whole-row link: wrap every field cell's content in a <Link>
	linkOpen, linkClose := "", ""
	if e.RowLink != nil {
		linkOpen = fmt.Sprintf("<Link %s>", linkToAttr(*e.RowLink))
		linkClose = "</Link>"
	}

	var lines []string
	lines = append(lines, fmt.Sprintf("%s<Table%s>", ind, cls))

	// THead
	lines = append(lines, fmt.Sprintf("%s  <THead>", ind))
	lines = append(lines, fmt.Sprintf("%s    <TR>", ind))
	for _, f := range fields {
		lines = append(lines, fmt.Sprintf("%s      <TH>%s</TH>", ind, toLabel(f)))
	}
	for range rowLinks {
		lines = append(lines, fmt.Sprintf("%s      <TH></TH>", ind))
	}
	for range rowActions {
		lines = append(lines, fmt.Sprintf("%s      <TH></TH>", ind))
	}
	lines = append(lines, fmt.Sprintf("%s    </TR>", ind))
	lines = append(lines, fmt.Sprintf("%s  </THead>", ind))

	// TBody
	lines = append(lines, fmt.Sprintf("%s  <TBody>", ind))
	lines = append(lines, fmt.Sprintf("%s    {%s.%s?.map(%s => (", ind, dataVar, optionalChainPath(e.Field), mapParams))
	lines = append(lines, fmt.Sprintf("%s      <TR key={%s}>", ind, keyExpr))
	for _, f := range fields {
		lines = append(lines, fmt.Sprintf("%s        <TD>%s{item.%s}%s</TD>", ind, linkOpen, optionalChainPath(f), linkClose))
	}
	for _, l := range rowLinks {
		lines = append(lines, fmt.Sprintf("%s        <TD>", ind))
		lines = append(lines, renderLinkJSX(l, dataVar, "item", indent+10, noBodyOps))
		lines = append(lines, fmt.Sprintf("%s        </TD>", ind))
	}
	for _, a := range rowActions {
		lines = append(lines, fmt.Sprintf("%s        <TD>", ind))
		lines = append(lines, renderActionJSX(a, indent+10, noBodyOps))
		lines = append(lines, fmt.Sprintf("%s        </TD>", ind))
	}
	lines = append(lines, fmt.Sprintf("%s      </TR>", ind))
	lines = append(lines, fmt.Sprintf("%s    ))}", ind))
	lines = append(lines, fmt.Sprintf("%s  </TBody>", ind))

	lines = append(lines, fmt.Sprintf("%s</Table>", ind))

	return strings.Join(lines, "\n")
}
