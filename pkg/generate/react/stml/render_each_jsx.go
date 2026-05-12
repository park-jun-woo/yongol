//ff:func feature=stml-gen type=generator control=iteration dimension=1
//ff:what EachBlock의 배열 순회 JSX를 Table 구조로 생성한다
package stml

import (
	"fmt"
	"strings"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderEachJSX generates JSX for an EachBlock as a Table.
func renderEachJSX(e stmlparser.EachBlock, dataVar string, indent int) string {
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

	var lines []string
	lines = append(lines, fmt.Sprintf("%s<Table%s>", ind, cls))

	// THead
	lines = append(lines, fmt.Sprintf("%s  <THead>", ind))
	lines = append(lines, fmt.Sprintf("%s    <TR>", ind))
	for _, f := range fields {
		lines = append(lines, fmt.Sprintf("%s      <TH>%s</TH>", ind, toLabel(f)))
	}
	lines = append(lines, fmt.Sprintf("%s    </TR>", ind))
	lines = append(lines, fmt.Sprintf("%s  </THead>", ind))

	// TBody
	lines = append(lines, fmt.Sprintf("%s  <TBody>", ind))
	lines = append(lines, fmt.Sprintf("%s    {%s.%s?.map(%s => (", ind, dataVar, optionalChainPath(e.Field), mapParams))
	lines = append(lines, fmt.Sprintf("%s      <TR key={%s}>", ind, keyExpr))
	for _, f := range fields {
		lines = append(lines, fmt.Sprintf("%s        <TD>{item.%s}</TD>", ind, optionalChainPath(f)))
	}
	lines = append(lines, fmt.Sprintf("%s      </TR>", ind))
	lines = append(lines, fmt.Sprintf("%s    ))}", ind))
	lines = append(lines, fmt.Sprintf("%s  </TBody>", ind))

	lines = append(lines, fmt.Sprintf("%s</Table>", ind))

	return strings.Join(lines, "\n")
}
